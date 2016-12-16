package plugin

// Dependencies
import (
	"os/exec"
	"strings"
)

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

// Update the deployment image
func commandKubernetesUpdateDeployment(deployment, container, image, project, tag string) *exec.Cmd {
	// Alow this two options
	// `kubectl set image deployment/demo demo=gcr.io/yebo-project/demo:latest`
	// `kubectl set image deployment/develop develop=gcr.io/yebo-project/demo:latest`
	return exec.Command(kubectlCmd, "set", "image", "deployment/" + deployment, container + "=gcr.io/" + project + "/" + image + ":" + tag)
}
