package container

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func List(client *client.Client) error {
	containers, err := client.ContainerList(
		context.Background(),
		container.ListOptions{
			All: true,
		},
	)

	if err != nil {
		return err
	}

	for _, c := range containers {
		fmt.Printf(
			"%s\t%s\t%s\t%s\n",
			c.ID[:12],
			c.Image,
			c.State,
			c.Status,
		)
	}

	return nil
}

func Start(client *client.Client, id string) error {
	return client.ContainerStart(
		context.Background(),
		id,
		container.StartOptions{},
	)
}

func Stop(client *client.Client, id string) error {
	return client.ContainerStop(
		context.Background(),
		id,
		container.StopOptions{},
	)
}

func Remove(client *client.Client, id string) error {
	return client.ContainerRemove(
		context.Background(),
		id,
		types.ContainerRemoveOptions{},
	)
}
