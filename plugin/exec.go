package plugin

// Dependencies
import (
	"os"
	"fmt"
	"io/ioutil"
	"text/template"
)

// Executes the auth command
func (p *Plugin) ExecCommandAuth() error {
	// Create a temporary file
	createTmpFile(googleKeyJsonPath, p.Google.Credentials, p.Debug)

	// Remove the file when everything is done
	defer os.Remove(googleKeyJsonPath)

	// Generate the command
	cmd := commandAuth(googleKeyJsonPath)

	// Info!
	fmt.Println("[INFO] Getting GCloud credentials")

	// Trace the command
	if p.Debug {
		traceCommand(cmd)
	}

	// Run the command
	return runCommand(cmd, p.Debug)
}

// Get container engine credentials
func (p *Plugin) ExecGetCredentials() error {
	//
	cmd := commandGetClusterCredentials(p.Kubernetes.Cluster, p.Google.Zone, p.Google.Project)

	// Info!
	fmt.Println("[INFO] Getting Kubernetes cluster credentials")

	// Trace the command
	if p.Debug {
		traceCommand(cmd)
	}

	return runCommand(cmd, p.Debug)
}

// Set namespace
func (p *Plugin) ExecSetNamespace() error {
	// Set Kubernetes context
	if p.Kubernetes.Namespace != "" {
		//
		cmd := commandSetKubernetesContext(generateKubernetesClusterName(p.Google.Project, p.Google.Zone, p.Kubernetes.Cluster), p.Kubernetes.Namespace)

		// Info!
		fmt.Println("[INFO] Set Kubernetes namespace")

		// Trace command
		if p.Debug {
			traceCommand(cmd)
		}

		// Run the command
		return runCommand(cmd, p.Debug)
	}

	// Nothing here!
	return nil
}

// Apply changes to the configuration
func (p *Plugin) ExecApplyKubernetes() error {
	// Generate the kubernetes filename
	kFile := kubeFilePath + "-" + p.Drone.BuildNumber + ".yml"

	// Check if the passed is a file:
	if _, err := os.Stat("./" + p.Kubernetes.File); err == nil {
		// Read the file contents
		content, err := ioutil.ReadFile("./" + p.Kubernetes.File);

		// Check any error
		if err != nil {
			return fmt.Errorf("Error while reading the file" + p.Kubernetes.File)
		}

		// Set the file content
		p.Kubernetes.File = string(content)
	}

	// Parse using the template
	tmpl, err := template.New(kFile).Option("missingkey=error").Parse(p.Kubernetes.File)
	if err != nil {
		return fmt.Errorf("Error parsing the template %s\n", err)
	}

	// Create a file
	f, err := os.Create(kFile)
	if err != nil {
		return fmt.Errorf("Error while creating the file %s\n", kFile)
	}

	// Create the informations used by the templates
	infos := TemplateInfos{
		Name:        p.Drone.Name,
		Tag:         p.Drone.Tag,
		BuildNumber: p.Drone.BuildNumber,
		Namespace:   p.Kubernetes.Namespace,
		Cluster:     p.Kubernetes.Cluster,
		Container:   p.Kubernetes.Container,
		Image:       p.Kubernetes.Image,
		Deployment:  p.Kubernetes.Deployment,
	}

	// Run the template
	err = tmpl.Execute(f, infos)
	if err != nil {
		return fmt.Errorf("Error while generating the file: %s\n", err)
	}

	// Close and delete the file
	defer f.Close()
	defer os.Remove(kFile)

	// Generate the command
	cmd := commandKubernetesApply(kFile)

	// Info
	fmt.Println("[INFO] Updating files")

	// Print the generated file
	if p.Debug {
		res, _ := ioutil.ReadFile(kFile)
		fmt.Printf("%+v\n", string(res))
	}

	// Trace command
	if p.Debug {
		traceCommand(cmd)
	}

	// Success!
	return runCommand(cmd, p.Debug)
}

// Run the deployment update
func (p *Plugin) ExecDeploymentUpdate() error {
	//
	cmd := commandKubernetesUpdateDeployment(p.Kubernetes.Deployment, p.Kubernetes.Container, p.Kubernetes.Image, p.Google.Project, p.Drone.Tag)

	// Info
	fmt.Println("[INFO] Updating the deployments")

	// Trace command
	if p.Debug {
		traceCommand(cmd)
	}

	// Run the command
	return runCommand(cmd, p.Debug)
}
