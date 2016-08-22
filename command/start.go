package command

import (
	"fmt"
	"net/http"
	"time"

	"github.com/DaveBlooman/spurious/output"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/urfave/cli"
)

// CmdStart Start containers
func CmdStart(c *cli.Context) error {
	client, err := docker.NewClient(GetEndpoint())
	if err != nil {
		output.Error("Unable to connect to Docker Daemon")
		return err
	}
	containers, err := client.ListContainers(docker.ListContainersOptions{Filters: stoppedSpuriousContainers})
	if err != nil {
		output.Error("Unable to list containers")
		return err
	}

	// Check for stopped containers and restart them
	if len(containers) != 0 {
		for _, containerImage := range Images {
			for _, container := range containers {
				if containerImage == container.Image {
					err = client.StartContainer(container.ID, nil)
					if err != nil {
						output.Error("Unable to start container")
						return err
					}
				}
				if container.Image == "spurious/sqs" {
					sqs, sqserr := client.InspectContainer(container.ID)
					if sqserr != nil {
						output.Error("Unable to setup SQS")
						return sqserr
					}

					sqsPortNumber := sqs.NetworkSettings.Ports["4568/tcp"][0].HostPort
					setupSQS(sqsPortNumber)
				}
			}
			output.Standard("Resuming containers")
		}
		return nil
	}

	output.Standard("Starting containers")

	container, err := client.CreateContainer(docker.CreateContainerOptions{
		Name: "spurious-sqs",
		HostConfig: &docker.HostConfig{
			PublishAllPorts: true,
		},
		Config: &docker.Config{
			Image:    "spurious/sqs",
			Labels:   map[string]string{"uk.co.spurious": "true"},
			Hostname: "sqs.spurious.localhost",
		},
	})
	if err != nil {
		output.Error("Unable to create container")
		return err
	}
	err = client.StartContainer(container.ID, nil)
	if err != nil {
		output.Error("Unable to start container")
		return err
	}

	sqs, err := client.InspectContainer(container.ID)
	if err != nil {
		output.Error("Unable to setup SQS")
		return err
	}
	sqsPortNumber := sqs.NetworkSettings.Ports["4568/tcp"][0].HostPort

	setupSQS(sqsPortNumber)

	container, err = client.CreateContainer(docker.CreateContainerOptions{
		Name: "spurious-s3",
		HostConfig: &docker.HostConfig{
			PublishAllPorts: true,
		},
		Config: &docker.Config{
			Hostname: "s3.spurious.localhost",
			Image:    "spurious/s3",
			Labels:   map[string]string{"uk.co.spurious": "true"},
			Cmd:      []string{"-r", "/var/data/fake-s3", "-p", "4569", "-H", "s3.spurious.localhost"},
		},
	})
	if err != nil {
		output.Error("Unable to create container")
		return err
	}
	err = client.StartContainer(container.ID, nil)
	if err != nil {
		output.Error("Unable to start container")
		return err
	}

	container, err = client.CreateContainer(docker.CreateContainerOptions{
		Name: "spurious-dynamo",
		HostConfig: &docker.HostConfig{
			PublishAllPorts: true,
		},
		Config: &docker.Config{
			Hostname: "dynamodb.spurious.localhost",
			Image:    "spurious/dynamodb",
			Labels:   map[string]string{"uk.co.spurious": "true"},
		},
	})
	if err != nil {
		output.Error("Unable to create container")
		return err
	}
	err = client.StartContainer(container.ID, nil)
	if err != nil {
		output.Error("Unable to start container")
		return err
	}

	container, err = client.CreateContainer(docker.CreateContainerOptions{
		Name: "spurious-memcached",
		HostConfig: &docker.HostConfig{
			PublishAllPorts: true,
		},
		Config: &docker.Config{
			Hostname: "localhost",
			Image:    "smaj/memcached",
			Labels:   map[string]string{"uk.co.spurious": "true"},
		},
	})
	if err != nil {
		output.Error("Unable to create container")
		return err
	}
	err = client.StartContainer(container.ID, nil)
	if err != nil {
		output.Error("Unable to start container")
		return err
	}

	container, err = client.CreateContainer(docker.CreateContainerOptions{
		Name: "spurious-elasticache",
		HostConfig: &docker.HostConfig{
			Links:           []string{"spurious-memcached:memcached01"},
			PublishAllPorts: true,
		},
		Config: &docker.Config{
			Hostname: "localhost",
			Image:    "spurious/elasticache",
			Labels:   map[string]string{"uk.co.spurious": "true"},
			Env:      []string{"FAKEELASTICACHEDEFAULTHOST=127.0.0.1"},
		},
	})
	if err != nil {
		output.Error("Unable to create container")
		return err
	}
	err = client.StartContainer(container.ID, nil)
	if err != nil {
		output.Error("Unable to start container")
		return err
	}

	container, err = client.CreateContainer(docker.CreateContainerOptions{
		Name: "spurious-elasticache-docker",
		HostConfig: &docker.HostConfig{
			Links:           []string{"spurious-memcached:memcached01"},
			PublishAllPorts: true,
		},
		Config: &docker.Config{
			Hostname: "localhost",
			Image:    "spurious/elasticache",
			Labels:   map[string]string{"uk.co.spurious": "true"},
		},
	})
	if err != nil {
		output.Error("Unable to create container")
		return err
	}
	err = client.StartContainer(container.ID, nil)
	if err != nil {
		output.Error("Unable to start container")
		return err
	}

	container, err = client.CreateContainer(docker.CreateContainerOptions{
		Name: "spurious-browser",
		HostConfig: &docker.HostConfig{
			PublishAllPorts: true,
			Links: []string{
				"spurious-sqs:sqs.spurious.localhost",
				"spurious-dynamo:dynamodb.spurious.localhost",
				"spurious-s3:s3.spurious.localhost"},
		},
		Config: &docker.Config{
			Hostname: "browser.spurious.localhost",
			Image:    "spurious/browser",
			Labels:   map[string]string{"uk.co.spurious": "true"},
		},
	})
	err = client.StartContainer(container.ID, nil)
	if err != nil {
		output.Error("Unable to start container")
		return err
	}

	output.Standard("All containers start")

	return nil
}

func setupSQS(sqsPortNumber string) {

	// Setup SQS
	time.Sleep(time.Second * 2)
	response, err := http.Get(fmt.Sprintf("http://localhost:%s/host-details?host=sqs.spurious.localhost&port=%s", sqsPortNumber, sqsPortNumber))

	if err != nil {
		output.Error("Unable to setup SQS")
	} else {
		defer response.Body.Close()
	}

}
