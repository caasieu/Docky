package cli

import (
	"fmt"

	"github.com/caasieu/dockyard/internal/docker"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dockyard",
	Short: "A simple Docker container and image manager",
}

func Execute() error {
	return rootCmd.Execute()
}

func getDockerClient() (*client.Client, error) {
	return docker.NewClient()
}

func errorMessage(err error) {
	fmt.Printf("Error: %v\n", err)
}

func init() {
	rootCmd.AddCommand(containerCmd)
	rootCmd.AddCommand(imageCmd)
}
