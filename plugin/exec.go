package plugin

// Dependencies
import (
	"os"
)

// Executes the auth command
func (p *Plugin) ExecCommandAuth() error {
	// Create a temporary file
	createTmpFile(googleKeyJsonPath, p.Google.Credentials, p.Debug)

	// Remove the file when everything is done
	defer os.Remove(googleKeyJsonPath)

	// Generate the command
	cmd := commandAuth(googleKeyJsonPath)

	// Trace the bug
	traceCommand(cmd)

	// Run the command
	return runCommand(cmd, p.Debug)
}

// Get container engine credentials
func (p *Plugin) ExecGetCredentials() error {
	//
	cmd := commandGetClusterCredentials(p.Kubernetes.Cluster, p.Google.Zone, p.Google.Project)

	// Trace the bug
	traceCommand(cmd)

	return runCommand(cmd, p.Debug)
}

// Set namespace
func (p *Plugin) ExecSetNamespace() error {
	// Set Kubernetes context
	if p.Kubernetes.Cluster != "" {
		//
		cmd := commandSetKubernetesContext(generateKubernetesClusterName(p.Google.Project, p.Google.Zone, p.Kubernetes.Cluster), p.Kubernetes.Namespace)

		// Trace the bug
		traceCommand(cmd)

		// Run the command
		return runCommand(cmd, p.Debug)
	}

	// Nothing here!
	return nil
}

// Run the deployment update
func (p *Plugin) ExecDeploymentUpdate() error {
	//
	cmd := commandKubernetesUpdateDeployment(p.Drone.Name, p.Google.Project, p.Drone.Tag)

	// Trace the bug
	traceCommand(cmd)

	// Run the command
	return runCommand(cmd, p.Debug)
}
