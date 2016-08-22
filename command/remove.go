package command

import (
	"fmt"

	"github.com/DaveBlooman/spurious/output"
	"github.com/fsouza/go-dockerclient"
	"github.com/urfave/cli"
)

// CmdRemove removes containers
func CmdRemove(c *cli.Context) error {
	client, err := docker.NewClient(GetEndpoint())
	if err != nil {
		output.Error("Unable to connect to Docker Daemon")
		return err
	}

	containers, err := client.ListContainers(docker.ListContainersOptions{Filters: stoppedSpuriousContainers})
	if err != nil {
		output.Error("Enable to list containers")
		return err
	}
	for _, container := range containers {

		client.RemoveContainer(docker.RemoveContainerOptions{ID: (container.ID), Force: true})
		fmt.Printf(`Removed "%v"`+"\n", container.Names[0])
	}

	return nil
}
