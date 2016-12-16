package plugin

// Dependencies
import (
	"testing"
	"fmt"
)

// Define the sample plugin
var p = Plugin{
	// Drone env
	DroneEnv: map[string]string{},
	// Google configurations
	Google: Google{
		Credentials: "",
		Zone: "us-east-d",
		Project: "drone-gke",
	},
	// Kubernetes Configurations
	Kubernetes: Kubernetes{
		Cluster: "drone",
		Namespace: "drone",
	},
}

//
func TestAuthCommand(t *testing.T) {
}
