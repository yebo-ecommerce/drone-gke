package plugin

// Dependenciesw:w
import (
	"os/exec"
	"os"
	"strings"
	"fmt"
)

// Constants
const (
	gcloudCmd = "gcloud"
	kubectlCmd = "kubectl"
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
	}
)

//
func (p *Plugin) Exec() error {
	// Auth command
	cmd := p.commandAuth("hello.json")
	traceCommand(cmd)

	// Get Credentials command
	cmd = p.commandGetClusterCredentials()
	traceCommand(cmd)

	// Set Kubernetes context
	if p.Kubernetes.Cluster != "" {
		cmd = p.commandSetKubernetesContext(p.generateKubernetesClusterName())
		traceCommand(cmd)
	}

	// Apply the changes from the file
	cmd = p.commandKubernetesApply("somefile.yml")
	traceCommand(cmd)

	// Everything OK
	return nil
}

// Command that auths into the gCloud
func (p *Plugin) commandAuth(keyPath string) *exec.Cmd {
	return exec.Command(gcloudCmd, "auth", "activate-service-account", "--key-file", keyPath)
}

// Command that will get the Google Container Cluster crendentials (Kubernetes)
func (p *Plugin) commandGetClusterCredentials() *exec.Cmd {
	return exec.Command(gcloudCmd, "container", "clusters", "get-credentials", p.Kubernetes.Cluster, "--zone", p.Google.Zone, "--project", p.Google.Project)
}

// Generate the cluster name ("gke_${project}_${zone}_${cluste-name}")
func (p *Plugin) generateKubernetesClusterName() string {
	return strings.Join([]string{"gke", p.Google.Project, p.Google.Zone, p.Kubernetes.Cluster}, "_")
}

// Set the Kubernetes context with namespace
func (p *Plugin) commandSetKubernetesContext(cluster string) *exec.Cmd {
	return exec.Command(kubectlCmd, "config", "set-context", cluster, "--namespace", p.Kubernetes.Namespace)
}

// Apply the file changes (or creation) to Kubernetes
func (p *Plugin) commandKubernetesApply(file string) *exec.Cmd {
	return exec.Command(kubectlCmd, "apply", "--filename", file)
}

// Traces the command to the os.Stdout
func traceCommand(cmd *exec.Cmd) {
	fmt.Fprintf(os.Stdout, "$ %s\n", strings.Join(cmd.Args, " "))
}
