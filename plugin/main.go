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

// Command that auths into the gCloud
func commandAuth(keyPath string) *exec.Cmd {
	return exec.Command(gcloudCmd, "auth", "activate-service-account", "--key-file", keyPath)
}

// Command that will get the Google Container Cluster crendentials (Kubernetes)
func commandGetClusterCredentials(cluster, zone, project string) *exec.Cmd {
	return exec.Command(gcloudCmd, "container", "clusters", "get-credentials", cluster, "--zone", zone, "--project", project)
}

// Generate the cluster name ("gke_${project}_${zone}_${cluste-name}")
func generateKubernetesClusterName(project, zone, cluster string) string {
	return strings.Join([]string{"gke", project, zone, cluster}, "_")
}

// Set the Kubernetes context with namespace
func commandSetKubernetesContext(cluster, namespace string) *exec.Cmd {
	return exec.Command(kubectlCmd, "config", "set-context", cluster, "--namespace", namespace)
}

// Apply the file changes (or creation) to Kubernetes
func commandKubernetesApply(file string) *exec.Cmd {
	return exec.Command(kubectlCmd, "apply", "--filename", file)
}

// Traces the command to the os.Stdout
func traceCommand(cmd *exec.Cmd) {
	fmt.Fprintf(os.Stdout, "$ %s\n", strings.Join(cmd.Args, " "))
}
