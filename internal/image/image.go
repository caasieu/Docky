package image

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func List(client *client.Client) error {
	images, err := client.ImageList(
		context.Background(),
		image.ListOptions{},
	)

	if err != nil {
		return err
	}

	for _, img := range images {
		fmt.Printf(
			"%s\t%d MB\t%v\n",
			img.ID[7:19],
			img.Size/(1024*1024),
			img.RepoTags,
		)
	}

	return nil
}

func Pull(client *client.Client, name string) error {
	reader, err := client.ImagePull(
		context.Background(),
		name,
		image.PullOptions{},
	)

	if err != nil {
		return err
	}

	defer reader.Close()

	_, err = io.Copy(os.Stdout, reader)

	return err
}

func Remove(client *client.Client, id string) error {
	_, err := client.ImageRemove(
		context.Background(),
		id,
		image.RemoveOptions{},
	)

	return err
}
