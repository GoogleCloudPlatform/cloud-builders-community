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

	if *mode == "trigger" {
		// Trigger another build to run the monitor.
		log.Printf("Starting trigger mode for build %s", *build)
		slackbot.Trigger(ctx, *build, *webhook)
	} else {
		// Monitor the other build until completion.
		log.Printf("Starting monitor mode for build %s", *build)
		slackbot.Monitor(ctx, *build, *webhook)
	}

}
