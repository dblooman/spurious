package command

import (
	"fmt"
	"os"

	"github.com/DaveBlooman/spurious/output"
	"github.com/fsouza/go-dockerclient"
	"github.com/urfave/cli"
)

// CmdInit downloads images
func CmdInit(c *cli.Context) error {
	for _, image := range Images {
		err := getImage(image)
		if err != nil {
			return err
		}
	}
	return nil
}

func getImage(image string) error {
	client, _ := docker.NewClient(Endpoint)

	err := client.PullImage(docker.PullImageOptions{Repository: image, OutputStream: os.Stdout}, docker.AuthConfiguration{})

	if err != nil {
		output.Standard(fmt.Sprintf("Error pulling image: %s", err))
		return err
	}

	output.Standard("Container: " + image + " finished")
	return nil
}
