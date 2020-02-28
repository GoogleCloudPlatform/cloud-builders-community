package builder

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/compute/metadata"
	"github.com/masterzen/winrm"
	"github.com/packer-community/winrmcp/winrmcp"
)

const (
	copyTimeout = 5 * time.Minute
	runTimeout  = 5 * time.Minute
)

// Remote represents a remote Windows server.
type Remote struct {
	Hostname   *string
	Username   *string
	Password   *string
	BucketName *string
}

type BuilderServer struct {
	ImageUrl *string
	VPC      *string
	Subnet   *string
	Region   *string
	Zone     *string
}

// Wait for server to be available.
func (r *Remote) Wait() error {
	timeout := time.Now().Add(time.Minute * 5)
	for time.Now().Before(timeout) {
		err := r.Run("ver")
		if err == nil {
			return nil
		}
		time.Sleep(10 * time.Second)
	}
	return errors.New("Timed out waiting for server to be available")
}

// Copy workspace from Linux to Windows.
func (r *Remote) Copy(inputPath string) error {
	defer func() {
		// Flush stdout
		fmt.Println()
	}()

	hostport := fmt.Sprintf("%s:5986", *r.Hostname)
	c, err := winrmcp.New(hostport, &winrmcp.Config{
		Auth:                  winrmcp.Auth{User: *r.Username, Password: *r.Password},
		Https:                 true,
		Insecure:              true,
		TLSServerName:         "",
		CACertBytes:           nil,
		OperationTimeout:      copyTimeout,
		MaxOperationsPerShell: 15,
	})
	if err != nil {
		log.Printf("Error creating connection to remote for copy: %+v", err)
		return err
	}

	// First try to create a bucket and have the Windows VM download it via a
	// GS URL. If that fails, use the remote copy method.
	err = r.copyViaBucket(
		context.Background(),
		inputPath,
		`C:\workspace`,
	)
	if err == nil {
		// Successfully copied via GCE bucket
		log.Printf("Successfully copied data via GCE bucket")
		return nil
	}

	log.Printf("Failed to copy data via GCE bucket: %v", err)

	err = c.Copy(inputPath, `C:\workspace`)
	if err != nil {
		log.Printf("Error copying workspace to remote: %+v", err)
		return err
	}

	return nil
}

func (r *Remote) copyViaBucket(ctx context.Context, inputPath, outputPath string) error {
	var bucket string
	if r.BucketName == nil || *r.BucketName == "" {
		// Cloud Build creates a bucket called <PROJECT-ID>_cloudbuild. Put
		// the object there.
		client := metadata.NewClient(http.DefaultClient)
		projectID, err := client.ProjectID()
		if err != nil {
			return fmt.Errorf("metadata.ProjectID: %v", err)
		}
		bucket = fmt.Sprintf("%s_cloudbuild", projectID)
	} else {
		bucket = *r.BucketName
	}
	object := fmt.Sprintf("windows-builder-%d", time.Now().UnixNano())

	gsURL, err := writeZipToBucket(
		ctx,
		bucket,
		object,
		inputPath,
	)
	if err != nil {
		return err
	}

	pwrScript := fmt.Sprintf(`
$ErrorActionPreference = "Stop"
$ProgressPreference = 'SilentlyContinue'
gsutil cp %q c:\workspace.zip
Expand-Archive -Path c:\workspace.zip -DestinationPath c:\workspace -Force
`, gsURL)

	// Now tell the Windows VM to download it.
	return r.Run(winrm.Powershell(pwrScript))
}

// Run a command on the Windows remote.
func (r *Remote) Run(command string) error {
	cmdstring := fmt.Sprintf(`cd c:\workspace & %s`, command)
	endpoint := winrm.NewEndpoint(*r.Hostname, 5986, true, true, nil, nil, nil, runTimeout)
	w, err := winrm.NewClient(endpoint, *r.Username, *r.Password)
	if err != nil {
		return err
	}
	shell, err := w.CreateShell()
	if err != nil {
		return err
	}
	var cmd *winrm.Command
	cmd, err = shell.Execute(cmdstring)
	if err != nil {
		return err
	}

	go io.Copy(os.Stdout, cmd.Stdout)
	go io.Copy(os.Stderr, cmd.Stderr)

	cmd.Wait()
	shell.Close()
	return nil
}
