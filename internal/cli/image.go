package cli

import (
	dockerimage "github.com/caasieu/dockyard/internal/image"
	"github.com/spf13/cobra"
	"fmt"
)

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Manage images",
}

var imageListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List images",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getDockerClient()

		if err != nil {
			errorMessage(err)
			return
		}

		defer client.Close()

		if err := dockerimage.List(client); err != nil {
			errorMessage(err)
		}
	},
}

var imagePullCmd = &cobra.Command{
	Use:   "pull [image]",
	Short: "Pull an image",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getDockerClient()

		if err != nil {
			errorMessage(err)
			return
		}

		defer client.Close()

		if err := dockerimage.Pull(client, args[0]); err != nil {
			errorMessage(err)
		}
	},
}

var imageRemoveCmd = &cobra.Command{
	Use:   "rm [image]",
	Short: "Remove an image",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := getDockerClient()

		if err != nil {
			errorMessage(err)
			return
		}

		defer client.Close()

		if err := dockerimage.Remove(client, args[0]); err != nil {
			errorMessage(err)
			return
		}

		fmt.Printf("Image %s removed\n", args[0])
	},
}

func init() {
	imageCmd.AddCommand(imageListCmd)
	imageCmd.AddCommand(imagePullCmd)
	imageCmd.AddCommand(imageRemoveCmd)
}
