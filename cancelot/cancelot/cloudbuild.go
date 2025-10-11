package cancelot

import (
	"bytes"
	"context"
	"log"
	"os/exec"
	"strings"

	cloudbuild "cloud.google.com/go/cloudbuild/apiv1/v2"
	"cloud.google.com/go/compute/metadata"
)

// getProject gets the project ID.
func getProject() (string, error) {
	// Test if we're running on GCE.
	if metadata.OnGCE() {
		// Use the GCE Metadata service.
		projectID, err := metadata.ProjectID()
		if err != nil {
			log.Printf("Failed to get project ID from instance metadata")
			return "", err
		}

		return projectID, nil
	}

	// Shell out to gcloud.
	cmd := exec.Command("gcloud", "config", "get", "project")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Printf("Failed to shell out to gcloud: %+v", err)
		return "", err
	}
	projectID := strings.TrimSuffix(out.String(), "\n")
	projectID = strings.TrimSuffix(projectID, "\r")

	return projectID, nil
}

func gcbClient(ctx context.Context) *cloudbuild.Client {
	client, err := cloudbuild.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create cloudbuild client: %v", err)
	}

	return client
}
