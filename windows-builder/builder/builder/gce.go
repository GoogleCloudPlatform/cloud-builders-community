package builder

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
	
	"github.com/pborman/uuid"

	"cloud.google.com/go/compute/metadata"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
)

const (
	zone               = "us-central1-f"
	instanceNamePrefix = "windows-builder"
	prefix             = "https://www.googleapis.com/compute/v1/projects/"
	imageURL           = prefix + "windows-cloud/global/images/windows-server-1803-dc-core-for-containers-v20180802"
	winrmport          = 5986
	startupCmd         = `winrm set winrm/config/Service/Auth @{Basic="true"}`
)

// Server encapsulates a GCE Instance.
type Server struct {
	context   *context.Context
	projectID string
	service   *compute.Service
	instance  *compute.Instance
	Remote
}

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
	cmd := exec.Command("gcloud", "config", "get-value", "project")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Printf("Failed to shell out to gcloud: %+v", err)
		return "", err
	}
	projectID := strings.TrimSuffix(out.String(), "\n")
	return projectID, nil
}

// NewServer creates a new Windows server on GCE.
func NewServer(ctx context.Context) *Server {
	// Get the current project ID.
	projectID, err := getProject()
	if err != nil {
		log.Fatalf("Cannot create new server without project ID: %+v", err)
		return nil
	}
	s := &Server{projectID: projectID}

	log.Printf("Starting GCE service in project %s", projectID)
	err = s.newGCEService(ctx)
	if err != nil {
		log.Fatalf("Failed to start GCE service: %v", err)
		return nil
	}
	err = s.newInstance()
	if err != nil {
		log.Fatalf("Failed to start Windows VM: %v", err)
		return nil
	}

	// Reset password
	username := "windows-builder"
	password, err := s.resetWindowsPassword(username)
	if err != nil {
		log.Fatalf("Failed to reset Windows password: %+v", err)
	}

	// Set firewall rule.
	err = s.setFirewallRule()
	if err != nil {
		log.Fatalf("Failed to set ingress firewall rule: %v", err)
	}
	log.Printf("Set ingress firewall rule successfully")

	// Get IP address.
	ip, err := s.getExternalIP()
	if err != nil {
		log.Fatalf("Failed to get external IP address: %v", err)
		return nil
	}

	// Set and return Remote.
	s.Remote = Remote{
		Hostname: &ip,
		Username: &username,
		Password: &password,
	}
	return s
}

// newGCEService creates a new Compute service.
func (s *Server) newGCEService(ctx context.Context) error {
	client, err := google.DefaultClient(ctx, compute.ComputeScope)
	if err != nil {
		log.Printf("Failed to create Google Default Client: %v", err)
		return err
	}
	service, err := compute.New(client)
	if err != nil {
		log.Printf("Failed to create Compute Service: %v", err)
		return err
	}

	s.service = service
	return nil
}

// newInstance starts a Windows VM on GCE and returns host, username, password.
func (s *Server) newInstance() error {
	scmd := startupCmd // TODO: find better way to take address of const
	name := "windows-builder-" + uuid.New()
	instance := &compute.Instance{
		Name:        name,
		MachineType: prefix + s.projectID + "/zones/" + zone + "/machineTypes/n1-standard-1",
		Disks: []*compute.AttachedDisk{
			{
				AutoDelete: true,
				Boot:       true,
				Type:       "PERSISTENT",
				InitializeParams: &compute.AttachedDiskInitializeParams{
					DiskName:    fmt.Sprintf("%s-pd", name),
					SourceImage: imageURL,
				},
			},
		},
		Metadata: &compute.Metadata{
			Items: []*compute.MetadataItems{
				&compute.MetadataItems{
					Key:   "windows-startup-script-cmd",
					Value: &scmd,
				},
			},
		},
		NetworkInterfaces: []*compute.NetworkInterface{
			&compute.NetworkInterface{
				AccessConfigs: []*compute.AccessConfig{
					&compute.AccessConfig{
						Type: "ONE_TO_ONE_NAT",
						Name: "External NAT",
					},
				},
				Network: prefix + s.projectID + "/global/networks/default",
			},
		},
		ServiceAccounts: []*compute.ServiceAccount{
			{
				Email: "default",
				Scopes: []string{
					compute.DevstorageFullControlScope,
					compute.ComputeScope,
				},
			},
		},
	}

	op, err := s.service.Instances.Insert(s.projectID, zone, instance).Do()
	if err != nil {
		log.Printf("GCE Instances insert call failed: %v", err)
		return err
	}
	err = s.waitForComputeOperation(op)
	if err != nil {
		log.Printf("Wait for instance start failed: %v", err)
		return err
	}

	etag := op.Header.Get("Etag")
	inst, err := s.service.Instances.Get(s.projectID, zone, name).IfNoneMatch(etag).Do()
	if err != nil {
		log.Printf("Could not get GCE Instance details after creation: %v", err)
		return err
	}
	log.Printf("Successfully created instance: %s", inst.Name)
	s.instance = inst
	return nil
}

// refreshInstance refreshes latest info from GCE into struct.
func (s *Server) refreshInstance() error {
	inst, err := s.service.Instances.Get(s.projectID, zone, s.instance.Name).Do()
	if err != nil {
		log.Printf("Could not refresh instance: %v", err)
		return err
	}
	s.instance = inst
	return nil
}

// DeleteInstance stops a Windows VM on GCE.
func (s *Server) DeleteInstance() error {
	_, err := s.service.Instances.Delete(s.projectID, zone, s.instance.Name).Do()
	if err != nil {
		log.Printf("Could not delete instance: %v", err)
		return err
	}
	return nil
}

// getExternalIP gets the external IP for an instance.
func (s *Server) getExternalIP() (string, error) {
	err := s.refreshInstance()
	if err != nil {
		log.Printf("Error refreshing instance: %+v", err)
	}
	for _, ni := range s.instance.NetworkInterfaces {
		for _, ac := range ni.AccessConfigs {
			if ac.Name == "External NAT" {
				return ac.NatIP, nil
			}
		}
	}
	return "", errors.New("Could not get external NAT IP from list")
}

// setFirewallRule allows ingress on WinRM port.
func (s *Server) setFirewallRule() error {
	list, err := s.service.Firewalls.List(s.projectID).Do()
	if err != nil {
		log.Printf("Could not list GCE firewalls: %+v", err)
		return err
	}
	for _, f := range list.Items {
		if f.Name == "allow-winrm-ingress" {
			log.Print("Firewall rule already exists")
			return nil
		}
	}

	firewallRule := &compute.Firewall{
		Allowed: []*compute.FirewallAllowed{
			&compute.FirewallAllowed{
				IPProtocol: "tcp",
				Ports:      []string{"5986"},
			},
		},
		Direction:    "INGRESS",
		Name:         "allow-winrm-ingress",
		SourceRanges: []string{"0.0.0.0/0"},
	}
	_, err = s.service.Firewalls.Insert(s.projectID, firewallRule).Do()
	if err != nil {
		log.Printf("Error setting firewall rule: %v", err)
		return err
	}
	return nil
}

//WindowsPasswordConfig stores metadata to be sent to GCE.
type WindowsPasswordConfig struct {
	key      *rsa.PrivateKey
	password string
	UserName string    `json:"userName"`
	Modulus  string    `json:"modulus"`
	Exponent string    `json:"exponent"`
	Email    string    `json:"email"`
	ExpireOn time.Time `json:"expireOn"`
}

//WindowsPasswordResponse stores data received from GCE.
type WindowsPasswordResponse struct {
	UserName          string `json:"userName"`
	PasswordFound     bool   `json:"passwordFound"`
	EncryptedPassword string `json:"encryptedPassword"`
	Modulus           string `json:"modulus"`
	Exponent          string `json:"exponent"`
	ErrorMessage      string `json:"errorMessage"`
}

// resetWindowsPassword securely resets the admin Windows password.
// See https://cloud.google.com/compute/docs/instances/windows/automate-pw-generation
func (s *Server) resetWindowsPassword(username string) (string, error) {
	//Create random key and encode
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Printf("Failed to generate random RSA key: %v", err)
		return "", err
	}
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(key.E))
	wpc := WindowsPasswordConfig{
		key:      key,
		UserName: username,
		Modulus:  base64.StdEncoding.EncodeToString(key.N.Bytes()),
		Exponent: base64.StdEncoding.EncodeToString(buf[1:]),
		Email:    "nobody@nowhere.com",
		ExpireOn: time.Now().Add(time.Minute * 5),
	}
	data, err := json.Marshal(wpc)
	dstring := string(data)
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", err)
		return "", err
	}

	//Write key to instance metadata and wait for op to complete
	log.Print("Writing Windows instance metadata for password reset")
	s.instance.Metadata.Items = append(s.instance.Metadata.Items, &compute.MetadataItems{
		Key:   "windows-keys",
		Value: &dstring,
	})
	op, err := s.service.Instances.SetMetadata(s.projectID, zone, s.instance.Name, &compute.Metadata{
		Fingerprint: s.instance.Metadata.Fingerprint,
		Items:       s.instance.Metadata.Items,
	}).Do()
	if err != nil {
		log.Printf("Failed to set instance metadata: %v", err)
		return "", err
	}
	err = s.waitForComputeOperation(op)
	if err != nil {
		log.Printf("Compute operation timed out")
		return "", err
	}

	//Read and decode password
	log.Print("Waiting for Windows password response")
	timeout := time.Now().Add(time.Minute * 5)
	hash := sha1.New()
	for time.Now().Before(timeout) {
		output, err := s.service.Instances.GetSerialPortOutput(s.projectID, zone, s.instance.Name).Port(4).Do()
		if err != nil {
			log.Printf("Unable to get serial port output: %v", err)
			return "", err
		}
		responses := strings.Split(output.Contents, "\n")
		for _, response := range responses {
			var wpr WindowsPasswordResponse
			if err := json.Unmarshal([]byte(response), &wpr); err != nil {
				continue
			}
			if wpr.Modulus == wpc.Modulus {
				decodedPassword, err := base64.StdEncoding.DecodeString(wpr.EncryptedPassword)
				if err != nil {
					log.Printf("Cannot Base64 decode password: %v", err)
					return "", err
				}
				password, err := rsa.DecryptOAEP(hash, rand.Reader, wpc.key, decodedPassword, nil)
				if err != nil {
					log.Printf("Cannot decrypt password response: %v", err)
					return "", err
				}
				return string(password), nil
			}
		}
		time.Sleep(2 * time.Second)
	}
	err = errors.New("Could not retrieve password before timeout")
	return "", err
}

// waitForComputeOperation waits for a compute operation
func (s *Server) waitForComputeOperation(op *compute.Operation) error {
	log.Printf("Waiting for %+v to complete", op.Name)
	timeout := time.Now().Add(120 * time.Second)
	for time.Now().Before(timeout) {
		newop, err := s.service.ZoneOperations.Get(s.projectID, zone, op.Name).Do()
		if err != nil {
			log.Printf("Failed to update operation status: %v", err)
			return err
		}
		if newop.Status == "DONE" {
			return nil
		}
		time.Sleep(1 * time.Second)
	}
	err := fmt.Errorf("Compute operation %s timed out", op.Name)
	return err
}
