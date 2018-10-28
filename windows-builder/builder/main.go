package main

import (
	"context"
	"flag"
	"log"

	"./builder"
)

var (
	hostname = flag.String("hostname", "", "Hostname of remote Windows server")
	username = flag.String("username", "", "Username on remote Windows server")
	password = flag.String("password", "", "Password on remote Windows server")
	command  = flag.String("command", "", "Command to run on remote Windows server")
)

func main() {
	log.Print("Starting Windows builder")
	flag.Parse()
	var r *builder.Remote
	var s *builder.Server

	// Connect to server
	if (*hostname != "") && (*username != "") && (*password != "") {
		r = &builder.Remote{
			Hostname: hostname,
			Username: username,
			Password: password,
		}
		log.Printf("Connecting to existing host %s", *r.Hostname)
	} else {
		ctx := context.Background()
		s = builder.NewServer(ctx)
		r = &s.Remote
	}
	log.Print("Waiting for server to become available")
	err := r.Wait()
	if err != nil {
		log.Fatalf("Error connecting to server: %+v", err)
	}

	// Copy workspace to remote machine
	log.Print("Copying workspace")
	err = r.Copy()
	if err != nil {
		log.Fatalf("Error copying workspace: %+v", err)
	}

	// Execute on remote
	log.Printf("Executing command %s", *command)
	err = r.Run(*command)
	if err != nil {
		log.Fatalf("Error executing command: %+v", err)
	}

	// Shut down server if started
	if s != nil {
		err = s.DeleteInstance()
		if err != nil {
			log.Fatalf("Failed to shut down instance: %+v", err)
		}
	}
}
