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
	if p.Debug {
		traceCommand(cmd)
	}

	// Run the command
	return runCommand(cmd, p.Debug)
}

//
func (p *Plugin) ExecDeploymentUpdate() {
	//
	cmd := commandKubernetesUpdateDeployment(p.Drone.Name, p.Google.Project, p.Drone.Tag)

	// Run the command
	return runCommand(cmd, p.Debug)
}
