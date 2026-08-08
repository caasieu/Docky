package docker

import (
	"context"

	"github.com/docker/docker/client"
)

func NewClient() (*client.Client, error) {
	return client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
}

func Ping(client *client.Client) error {
	_, err := client.Ping(context.Background())
	return err
}
