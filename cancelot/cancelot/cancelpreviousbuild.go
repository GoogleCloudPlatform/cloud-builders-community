package cancelot

import (
	"context"
	"fmt"
	"google.golang.org/api/cloudbuild/v1"
	"log"
)

// Checks for previous running builds on the same branch, in order to cancel them
func CancelPreviousBuild(ctx context.Context, currentBuildId string, branchName string) {
	svc := gcbClient(ctx)
	project, err := getProject()
	if err != nil {
		log.Fatalf("Failed to get project: %v", err)
	}

	log.Printf("Going to fetch current build details for: %s", currentBuildId)

	currentBuildResponse, currentBuildError := svc.Projects.Builds.Get(project, currentBuildId).Do()
	if currentBuildError != nil {
		log.Fatalf("Failed to get build details from Cloud Build.  Will retry in one minute.")
	}

	log.Printf("Going to check ongoing jobs for branch: %s", branchName)

	onGoingJobFilter := fmt.Sprintf(`
		build_id != "%s" AND 
		source.repo_source.branch_name = "%s" AND 
		status = "WORKING" AND 
		start_time<"%s"`,
		currentBuildId,
		branchName,
		currentBuildResponse.StartTime)

	log.Printf("Builds filter created: %s", onGoingJobFilter)

	onGoingBuildsResponse, onGoingBuildsError := svc.Projects.Builds.List(project).Filter(onGoingJobFilter).Do()

	if onGoingBuildsError != nil {
		log.Fatalf("Failed to get builds from Cloud Build.  Will retry in one minute.")
	}

	onGoingBuilds := onGoingBuildsResponse.Builds
	numOfOnGoingBuilds := len(onGoingBuilds)

	log.Printf("Ongoing builds for %s has size of: %d", branchName, numOfOnGoingBuilds)

	if numOfOnGoingBuilds == 0 {
		return
	}

	for _, build := range onGoingBuilds {
		log.Printf("Going to cancel build with id: %s", build.Id)

		cancelBuildCall := svc.Projects.Builds.Cancel(project, build.Id, &cloudbuild.CancelBuildRequest{})
		buildCancelResponse, buildCancelError := cancelBuildCall.Do()

		if buildCancelError != nil {
			log.Fatalf("Failed to cancel build with id:%s - %v", build.Id, buildCancelError)
		}

		log.Printf("Cancelled build with id:%s", buildCancelResponse.Id)
	}
}
