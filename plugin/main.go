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
	kubeFilePath = "/tmp/genereated-kube"
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
		Cluster    string
		Namespace  string
		Deployment string
		Container  string
		Image      string
		File       string
	}

	// Drone Values
	Drone struct {
		Tag         string
		Name        string
		BuildNumber string
	}

	// Plugin execution params
	Plugin struct {
		Google     Google
		Kubernetes Kubernetes
		Drone      Drone
		Debug      bool
	}

	// Information used in the templates
	TemplateInfos struct {
		Name        string
		Tag         string
		Namespace   string
		Cluster     string
		Container   string
		Image       string
		Deployment  string
		BuildNumber string
	}
)

//
func (p *Plugin) Exec() error {
	// Authenticate!
	if p.ExecCommandAuth() != nil {
		return fmt.Errorf("[ERROR] Could not authenticate")
	}

	// Get Credentials command
	if p.ExecGetCredentials() != nil {
		return fmt.Errorf("[ERROR] Could not get GKE credentials")
	}

	// Set namespace
	if p.ExecSetNamespace() != nil {
		return fmt.Errorf("[ERROR] Could not set the kubernetes namespace")
	}


	// Check if there is any kubernetes file
	if p.Kubernetes.File == "" {
		//
		if p.ExecDeploymentUpdate() != nil {
			return fmt.Errorf("[ERROR] Could not update the deployment")
		}
	} else {
		//
		if p.ExecApplyKubernetes() != nil {
			return fmt.Errorf("[ERROR] Could not apply changes to the files")
		}
	}

	// Everything OK
	return nil
}

//
func runCommand(cmd *exec.Cmd, debug bool) error {
	// Check if it's in debug mode
	if debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	// Run the command
	return cmd.Run()
}

// Traces the command to the os.Stdout
func traceCommand(cmd *exec.Cmd) {
	fmt.Fprintf(os.Stdout, "+ %s\n", strings.Join(cmd.Args, " "))
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
