package cancelot

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	"github.com/avast/retry-go"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/iterator"
)

// CancelPreviousBuild checks for previous running builds on the same branch, in order to cancel them
func CancelPreviousBuild(ctx context.Context, currentBuildID string, branchName string, sameTriggerOnly bool) {
	err := retry.Do(
		func() error {
			err := cancelPreviousBuild(ctx, currentBuildID, branchName, sameTriggerOnly)
			if err != nil {
				log.Printf("Will retry -- %v", err)
			}
			return err
		},
	)
	if err != nil {
		log.Fatal("Too many failed retries")
	}
}

func cancelPreviousBuild(ctx context.Context, currentBuildID string, branchName string, sameTriggerOnly bool) error {
	client := gcbClient(ctx)
	project, err := getProject()
	if err != nil {
		return fmt.Errorf("getProject: %w", err)
	}

	log.Printf("Going to fetch current build details for: %s", currentBuildID)

	currentBuild, err := client.GetBuild(ctx, &cloudbuildpb.GetBuildRequest{ProjectId: project, Id: currentBuildID})
	if err != nil {
		return fmt.Errorf("client.GetBuild: %w", err)
	}

	filter := fmt.Sprintf(
		`build_id != "%s" AND (status = "WORKING" OR status = "QUEUED") AND create_time < "%s"`,
		currentBuildID,
		currentBuild.StartTime.AsTime().Format(time.RFC3339),
	)

	if sameTriggerOnly {
		filter = fmt.Sprintf(`%s AND trigger_id = "%s"`, filter, currentBuild.BuildTriggerId)
	}

	filterRepoLocally := false
	repoName := currentBuild.Source.GetRepoSource().GetRepoName()

	if repoName == "" {
		// Connected repos don't have repo_name/branch_name filled in, so we need to resort to additional local filtering.
		filterRepoLocally = true
		repoName = currentBuild.Substitutions["REPO_NAME"]
	} else {
		filter = fmt.Sprintf(`%s AND source.repo_source.repo_name = "%s" AND source.repo_source.branch_name = "%s"`, filter, repoName, branchName)
	}

	log.Printf("Using builds filter: %s", filter)
	if filterRepoLocally {
		log.Println("Using local repo and branch filtering as this trigger is configured with a connected source")
	}

	var cancels errgroup.Group
	iter := client.ListBuilds(ctx, &cloudbuildpb.ListBuildsRequest{ProjectId: project, Filter: filter})
	for {
		build, err := iter.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				break
			}

			return fmt.Errorf("client.ListBuilds iter.Next: %w", err)
		}

		if filterRepoLocally && (build.Substitutions["REPO_NAME"] != repoName || build.Substitutions["BRANCH_NAME"] != branchName) {
			continue
		}

		cancels.Go(func() error {
			log.Printf("Canceling build %s (started at %s)", build.Id, build.CreateTime.AsTime().Format(time.RFC3339))

			_, err := client.CancelBuild(ctx, &cloudbuildpb.CancelBuildRequest{ProjectId: build.ProjectId, Id: build.Id})
			return err
		})
	}

	return cancels.Wait()
}
