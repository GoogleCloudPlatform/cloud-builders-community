package main

import (
	"context"
	"flag"
	"log"
	"os"

	"./builder"
)

var (
	hostname = flag.String("hostname", "", "Hostname of remote Windows server")
	username = flag.String("username", "", "Username on remote Windows server")
	password = flag.String("password", os.Getenv("PASSWORD"), "Password on remote Windows server")
	command  = flag.String("command", "", "Command to run on remote Windows server")
	notCopyWorkspace = flag.Bool("not-copy-workspace", false, "If copy workspace or not")
	image    = flag.String("image", "windows-cloud/global/images/windows-server-2019-dc-for-containers-v20191210", "Windows image to start the server from")
	network = flag.String("network", "default", "The VPC name to use when creating the Windows server")
	subnetwork = flag.String("subnetwork", "default", "The Subnetwork name to use when creating the Windows server")
	region = flag.String("region", "us-central1", "The region name to use when creating the Windows server")
	zone = flag.String("zone", "us-central1-f", "The zone name to use when creating the Windows server")
)

func main() {
	log.Print("Starting Windows builder")
	flag.Parse()
	var r *builder.Remote
	var s *builder.Server
	var bs *builder.BuilderServer

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
		bs = &builder.BuilderServer{
			ImageUrl: image,
			VPC: network,
			Subnet: subnetwork,
			Region: region,
			Zone: zone,
		}
		s = builder.NewServer(ctx, bs)
		r = &s.Remote
	}
	log.Print("Waiting for server to become available")
	err := r.Wait()
	if err != nil {
		log.Fatalf("Error connecting to server: %+v", err)
	}

	// Copy workspace to remote machine
	if (!*notCopyWorkspace) {
		log.Print("Copying workspace")
		err = r.Copy()
		if err != nil {
			log.Fatalf("Error copying workspace: %+v", err)
		}
	}

	// Execute on remote
	log.Printf("Executing command %s", *command)
	err = r.Run(*command)
	if err != nil {
		log.Fatalf("Error executing command: %+v", err)
	}

	// Shut down server if started
	if s != nil {
		err = s.DeleteInstance(bs)
		if err != nil {
			log.Fatalf("Failed to shut down instance: %+v", err)
		}
	}
}
