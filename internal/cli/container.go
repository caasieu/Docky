package cli

import (
	"fmt"

	dockercontainer "github.com/caasieu/dockyard/internal/container"
	"github.com/spf13/cobra"
)

var containerCmd = &cobra.Command{
	Use:   "container",
	Short: "Manage containers",
}

var containerListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List containers",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getDockerClient()

		if err != nil {
			errorMessage(err)
			return
		}

		defer client.Close()

		if err := dockercontainer.List(client); err != nil {
			errorMessage(err)
		}
	},
}

var containerStartCmd = &cobra.Command{
	Use:   "start [container]",
	Short: "Start a container",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getDockerClient()

		if err != nil {
			errorMessage(err)
			return
		}

		defer client.Close()

		if err := dockercontainer.Start(client, args[0]); err != nil {
			errorMessage(err)
			return
		}

		fmt.Printf("Container %s started\n", args[0])
	},
}

var containerStopCmd = &cobra.Command{
	Use:   "stop [container]",
	Short: "Stop a container",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getDockerClient()

		if err != nil {
			errorMessage(err)
			return
		}

		defer client.Close()

		if err := dockercontainer.Stop(client, args[0]); err != nil {
			errorMessage(err)
			return
		}

		fmt.Printf("Container %s stopped\n", args[0])
	},
}

var containerRemoveCmd = &cobra.Command{
	Use:   "rm [container]",
	Short: "Remove a container",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getDockerClient()

		if err != nil {
			errorMessage(err)
			return
		}

		defer client.Close()

		if err := dockercontainer.Remove(client, args[0]); err != nil {
			errorMessage(err)
			return
		}

		fmt.Printf("Container %s removed\n", args[0])
	},
}

func init() {
	containerCmd.AddCommand(containerListCmd)
	containerCmd.AddCommand(containerStartCmd)
	containerCmd.AddCommand(containerStopCmd)
	containerCmd.AddCommand(containerRemoveCmd)
}
