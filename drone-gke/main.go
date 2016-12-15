package main

// Dependencies
import (
	"os"
	"fmt"
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
			Name:   "kube.namespace",
			Usage:  "Kubernetes Namespace",
			EnvVar: "PLUGIN_NAMESPACE",
		},
	}

	// Try to run the command
	if err := app.Run(os.Args); err != nil {
		fmt.Errorf("[ERROR] Invalid arguments.")
	}
}

// Run the plugin
func run(c *cli.Context) error {
	//
	fmt.Println("Great stuff is coming!!")

	// Success!
	return nil
}
