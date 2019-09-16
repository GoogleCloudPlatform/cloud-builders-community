// Post build status results to Slack.

package main

import (
	"context"
	"flag"
	"log"

	"./cancelot"
)

var (
	currentBuildId = flag.String("current_build_id", "", "The current build id, in order to be excluded")
	branchName     = flag.String("branch_name", "", "BranchName to cancel previous ongoing jobs on")
)

func main() {
	log.Print("Starting cancelot")
	flag.Parse()
	ctx := context.Background()

	if *currentBuildId == "" {
		log.Fatalf("CurrentBuildId must be provided.")
	}

	if *branchName == "" {
		log.Fatalf("BranchName must be provided.")
	}

	cancelot.CancelPreviousBuild(ctx, *currentBuildId, *branchName)
	return
}
