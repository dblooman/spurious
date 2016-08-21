package command

import (
	"fmt"
	"strings"

	"github.com/DaveBlooman/spurious/output"
	"github.com/apcera/termtables"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/urfave/cli"
)

// CmdPorts for spurious containers
func CmdPorts(c *cli.Context) error {
	client, err := docker.NewClient(Endpoint)
	if err != nil {
		output.Error("Enable to connect to Docker Daemon")
		return err
	}
	containers, err := client.ListContainers(docker.ListContainersOptions{Filters: SpuriousContainers})
	if err != nil {
		output.Error("Unable to list containers")
		return err
	}

	var containerList []*docker.Container
	for _, container := range containers {
		spuriousContainer, err := client.InspectContainer(container.ID)
		if err != nil {
			output.Error("Unable to list containers")
			return err
		}
		containerList = append(containerList, spuriousContainer)
	}

	table := termtables.CreateTable()
	table.Style.SkipBorder = true
	table.Style.BorderX = ""
	table.Style.BorderY = ""
	table.Style.BorderI = ""
	table.Style.PaddingLeft = 0

	table.AddHeaders(output.TableHeader("Service"), output.TableHeader("Host"), output.TableHeader("Port"), output.TableHeader("Browser Link"))

	tableData := createTable(table, containerList)
	fmt.Println(tableData.Render())

	return nil
}

func createTable(table *termtables.Table, containerList []*docker.Container) *termtables.Table {
	for _, img := range containerList {

		var port string
		var browserLink string

		for _, v := range img.NetworkSettings.Ports {
			port = output.TableBody(v[0].HostPort)
		}

		hostname := output.TableBody(img.Config.Hostname)
		containerName := output.TableBody(strings.TrimPrefix(img.Name, "/"))

		if img.Config.Hostname == "localhost" {
			browserLink = "-"
		} else {
			browserLink = "http://" + hostname + ":" + port
		}

		browserLink = output.TableBody(browserLink)
		table.AddRow(containerName, hostname, port, browserLink)
	}
	return table
}
