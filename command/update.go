package command

import (
	"fmt"
	"os"

	"github.com/DaveBlooman/spurious/output"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/urfave/cli"
)

// CmdUpdate for updating images
func CmdUpdate(c *cli.Context) error {
	for _, image := range Images {
		err := updateImage(image)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateImage(image string) error {
	client, _ := docker.NewClient(GetEndpoint())

	err := client.PullImage(docker.PullImageOptions{Repository: image, OutputStream: os.Stdout}, docker.AuthConfiguration{})

	if err != nil {
		output.Standard(fmt.Sprintf("Error pulling image: %s", err))
		return err
	}

	output.Standard("Image: " + image + " update operation finished")
	return nil
}
