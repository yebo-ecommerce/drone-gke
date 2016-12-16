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

	if p.Debug {
		traceCommand(cmd)
	}

	// Run the command
	return runCommand(cmd, p.Debug)
}
