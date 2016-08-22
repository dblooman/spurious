package command

import (
	"fmt"

	"github.com/DaveBlooman/spurious/output"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/urfave/cli"
)

// CmdStop stops containers
func CmdStop(c *cli.Context) error {
	client, err := docker.NewClient(GetEndpoint())
	if err != nil {
		output.Error("Enable to connect to Docker Daemon")
		return err
	}
	containers, err := client.ListContainers(docker.ListContainersOptions{Filters: SpuriousContainers})
	if err != nil {
		output.Error("Unable to list containers")
		return err
	}

	if len(containers) == 0 {
		output.Standard("No Spurious containers running")
		return nil
	}

	output.Standard("Stopping Containers")

	for _, container := range containers {
		err := client.StopContainer(container.ID, 1)
		if err != nil {
			output.Error("Unable to stop container " + container.ID)
			return err
		}
	}

	output.Standard(fmt.Sprintf("Stopped %v containers", len(containers)))

	return nil
}
