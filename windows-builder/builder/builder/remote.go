package builder

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/masterzen/winrm"
	"github.com/packer-community/winrmcp/winrmcp"
)

const (
	copyTimeout = 5 * time.Minute
	runTimeout  = 5 * time.Minute
)

// Remote represents a remote Windows server.
type Remote struct {
	Hostname *string
	Username *string
	Password *string
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
func (r *Remote) Copy() error {
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
	err = c.Copy("/workspace", `C:\workspace`)
	if err != nil {
		log.Printf("Error copying workspace to remote: %+v", err)
		return err
	}

	// Flush stdout
	fmt.Println()
	return nil
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
