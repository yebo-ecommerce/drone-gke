package main

// Dependencies
import (
	"os"
	"fmt"
	"strings"

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
			EnvVar: "GOOGLE_CREDENTIALS,PLUGIN_CRENDENTIALS",
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
			Name:   "drone.name",
			Usage:  "Repository Name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "drone.tag",
			Usage:  "Commit Tag",
			EnvVar: "DRONE_TAG",
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
		},
	}

	// Execute it
	return p.Exec()
}
