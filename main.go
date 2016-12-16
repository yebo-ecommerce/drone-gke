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
		// Drone env
		DroneEnv: parseDroneEnvs(),
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

//
func parseDroneEnvs() map[string]string {
	// Create a new one
	res := map[string]string{}

	//
	for _, e := range os.Environ() {
		//
		pair := strings.Split(e, "=")

		// Check if it is a DRONE env var
		if strings.Contains(pair[0], "DRONE_") {
			// Set it to the result
			res[pair[0]] = pair[1]
		}
	}

	//
	return res
}
