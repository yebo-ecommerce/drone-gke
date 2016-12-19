package main

// Dependencies
import (
	"os"
	"fmt"

	plugin "github.com/yebo-ecommerce/drone-gke/plugin"
	"github.com/urfave/cli"
)

// Entrypoint
func main() {
	// Create a new APP
	app := cli.NewApp()

	// App informations
	app.Name = "Drone Google Cloud Container Engine plugin"
	app.Version = "0.0.0-beta"

	// Define its action
	app.Action = run

	// Flags
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "If the program is running in the Debug mode",
			EnvVar: "PLUGIN_DEBUG",
		},
		cli.StringFlag{
			Name:   "gcloud.credentials",
			Usage:  "Google Cloud JSON token file",
			EnvVar: "GOOGLE_CREDENTIALS,PLUGIN_CREDENTIALS",
		},
		cli.StringFlag{
			Name:   "gcloud.zone",
			Usage:  "Google Cloud Zone",
			EnvVar: "PLUGIN_ZONE",
		},
		cli.StringFlag{
			Name:   "gcloud.cluster",
			Usage:  "Google Cloud Cluster",
			EnvVar: "GOOGLE_CLUSTER,PLUGIN_CLUSTER",
		},
		cli.StringFlag{
			Name:   "gcloud.project",
			Usage:  "Google Cloud Project",
			EnvVar: "GOOGLE_PROJECT,PLUGIN_PROJECT",
		},
		cli.StringFlag{
			Name:   "kube.cluster",
			Usage:  "Kubernetes Cluster name",
			EnvVar: "KUBERNETES_CLUSTER,PLUGIN_CLUSTER",
		},
		cli.StringFlag{
			Name:   "kube.namespace",
			Usage:  "Kubernetes Namespace",
			EnvVar: "PLUGIN_NAMESPACE",
		},
		cli.StringFlag{
			Name:   "kube.deployment",
			Usage:  "Kubernetes Deployment",
			EnvVar: "PLUGIN_DEPLOYMENT",
		},
		cli.StringFlag{
			Name:   "kube.container",
			Usage:  "Kubernetes Container",
			EnvVar: "PLUGIN_CONTAINER",
		},
		cli.StringFlag{
			Name:   "kube.image",
			Usage:  "Kubernetes Image",
			EnvVar: "PLUGIN_CONTAINER_IMAGE",
		},
		cli.StringFlag{
			Name:   "kube.file",
			Usage:  "Kubernetes File content",
			EnvVar: "KUBERNETES_FILE,PLUGIN_KUBEFILE",
		},
		cli.StringFlag{
			Name:   "drone.name",
			Usage:  "Repository Name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "drone.tag",
			Usage:  "Commit Tag",
			EnvVar: "DRONE_TAG",
			Value:  "latest",
		},
	}

	// Try to run the command
	if err := app.Run(os.Args); err != nil {
		fmt.Errorf("[ERROR] Invalid arguments.")
	}
}

// Run the plugin
func run(c *cli.Context) error {
	// Create the plugin
	p := plugin.Plugin{
		// Debug
		Debug: c.Bool("debug"),
		// Drone stuff
		Drone: plugin.Drone{
			Tag: c.String("drone.tag"),
			Name: c.String("drone.name"),
		},
		// Google configurations
		Google: plugin.Google{
			Credentials: c.String("gcloud.credentials"),
			Zone: c.String("gcloud.zone"),
			Project: c.String("gcloud.project"),
		},
		// Kubernetes Configurations
		Kubernetes: plugin.Kubernetes{
			Cluster: c.String("kube.cluster"),
			Namespace: c.String("kube.namespace"),
			Container: c.String("kube.container"),
			Deployment: c.String("kube.deployment"),
			Image: c.String("kube.image"),
			File: c.String("kube.file"),
		},
	}

	// Env to use old kubernetes behaviour
	// https://github.com/kubernetes/kubernetes/issues/30617
	os.Setenv("CLOUDSDK_CONTAINER_USE_CLIENT_CERTIFICATE", "True")

	// Validate
	if p.Kubernetes.Container == "" {
		p.Kubernetes.Container = p.Drone.Name
	}

	if p.Kubernetes.Image == "" {
		p.Kubernetes.Image = p.Drone.Name
	}

	if p.Kubernetes.Deployment == "" {
		p.Kubernetes.Deployment = p.Drone.Name
	}

	// Execute it
	return p.Exec()
}
