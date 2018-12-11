// Post build status results to Slack.

package main

import (
	"context"
	"flag"
	"log"

	"./slackbot"
)

var (
	build   = flag.String("build", "", "Build ID being monitored")
	webhook = flag.String("webhook", "", "Slack webhook URL")
	mode    = flag.String("mode", "trigger", "Mode the builder runs in")
)

func main() {
	log.Print("Starting slackbot")
	flag.Parse()
	ctx := context.Background()

<<<<<<< HEAD
	if *webhook == "" {
		log.Fatalf("Slack webhook must be provided.")
	}

	if *build == "" {
		log.Fatalf("Build ID must be provided.")
	}

=======
>>>>>>> c71c184bd14f37cb0dc2cc7f6e948bafacf15483
	if *mode == "trigger" {
		// Trigger another build to run the monitor.
		log.Printf("Starting trigger mode for build %s", *build)
		slackbot.Trigger(ctx, *build, *webhook)
<<<<<<< HEAD
		return
	}
	if *mode == "monitor" {
		// Monitor the other build until completion.
		log.Printf("Starting monitor mode for build %s", *build)
		slackbot.Monitor(ctx, *build, *webhook)
		return
	}
	log.Fatalf("Mode must be provided.")
=======
	} else {
		// Monitor the other build until completion.
		log.Printf("Starting monitor mode for build %s", *build)
		slackbot.Monitor(ctx, *build, *webhook)
	}

>>>>>>> c71c184bd14f37cb0dc2cc7f6e948bafacf15483
}
