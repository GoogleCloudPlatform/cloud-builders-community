package slackbot

import (
	"context"
	"log"
	"time"
)

const maxErrors = 3

// Monitor polls Cloud Build until the build reaches completed status, then triggers the Slack event.
func Monitor(ctx context.Context, build string, webhook string) {
	svc := gcbClient(ctx)
	errors := 0
	project, err := getProject()
	if err != nil {
		log.Fatalf("Failed to get project: %v", err)
	}

	t := time.Tick(20 * time.Second)
	for {
		log.Printf("Polling build %s", build)
		lc := svc.Projects.Builds.Get(project, build)
		b, err := lc.Do()
		if err != nil {
			if errors <= maxErrors {
				log.Printf("Failed to get build details from Cloud Build.  Will retry in one minute.")
				errors++
				continue
			} else {
				log.Fatalf("Reached maximum number of errors (%d).  Exiting", maxErrors)
			}
		}
		switch b.Status {
		case "SUCCESS", "FAILURE", "INTERNAL_ERROR", "TIMEOUT", "CANCELLED":
			log.Printf("Terminal status reached.  Notifying")
			Notify(b, webhook)
			return
		}
		<-t
	}
}
