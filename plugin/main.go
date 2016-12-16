package plugin

// Dependenciesw:w
import (
	"os/exec"
	"io/ioutil"
	"os"
	"strings"
	"fmt"
)

// Constants
const (
	// Binaries
	gcloudCmd = "gcloud"
	kubectlCmd = "kubectl"

	// Paths
	googleKeyJsonPath = "/tmp/google-key.json"
)

// Plugin type
type (
	// Google configurations
	Google struct {
		Credentials string
		Zone        string
		Project     string
	}

	// Kubernetes stuff
	Kubernetes struct {
		Cluster   string
		Namespace string
	}

	// Plugin execution params
	Plugin struct {
		Google     Google
		Kubernetes Kubernetes
		DroneEnv   map[string]string
		Debug      bool
	}
)

//
func (p *Plugin) Exec() error {
	//
	if p.ExecCommandAuth() != nil {
		fmt.Println("error while executing")
	}

	// Auth command
	cmd := commandAuth("hello.json")
	traceCommand(cmd)

	// Get Credentials command
	cmd = commandGetClusterCredentials(p.Kubernetes.Cluster, p.Google.Zone, p.Google.Project)
	traceCommand(cmd)

	// Set Kubernetes context
	if p.Kubernetes.Cluster != "" {
		cmd = commandSetKubernetesContext(generateKubernetesClusterName(p.Google.Project, p.Google.Zone, p.Kubernetes.Cluster), p.Kubernetes.Namespace)
		traceCommand(cmd)
	}

	// Apply the changes from the file
	cmd = commandKubernetesApply("somefile.yml")
	traceCommand(cmd)

	// Everything OK
	return nil
}

//
func runCommand(cmd *exec.Cmd, debug bool) error {
	// Check if it's in debug mode
	if debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		cmd.Stdout = ioutil.Discard
		cmd.Stderr = ioutil.Discard
	}

	// Run the command
	return cmd.Run()
}

// Traces the command to the os.Stdout
func traceCommand(cmd *exec.Cmd) {
	fmt.Fprintf(os.Stdout, "$ %s\n", strings.Join(cmd.Args, " "))
}

// Create a temporary file
func createTmpFile(path, content string, debug bool) {
	// Create the file
	if ioutil.WriteFile(path, []byte(content), 0600) != nil {
		panic("error while creating tmp file")
	}

	// Debug?
	if debug {
		fmt.Println("Creating the temporary file:", path)
	}
}
