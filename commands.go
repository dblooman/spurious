package main

import (
	"fmt"
	"os"

	"github.com/DaveBlooman/spurious/command"
	"github.com/urfave/cli"
)

// GlobalFlags for CLI
var GlobalFlags = []cli.Flag{}

// Commands for CLI
var Commands = []cli.Command{

	{
		Name:   "ports",
		Usage:  "View host ports for containers",
		Action: command.CmdPorts,
		Flags:  []cli.Flag{},
	},

	{
		Name:   "start",
		Usage:  "Start Spurious containers",
		Action: command.CmdStart,
		Flags:  []cli.Flag{},
	},

	{
		Name:   "stop",
		Usage:  "Stop Spurious containers",
		Action: command.CmdStop,
		Flags:  []cli.Flag{},
	},

	{
		Name:   "update",
		Usage:  "Update Spurious Docker images",
		Action: command.CmdUpdate,
		Flags:  []cli.Flag{},
	},

	{
		Name:   "remove",
		Usage:  "Remove exited Spurious images",
		Action: command.CmdRemove,
		Flags:  []cli.Flag{},
	},

	{
		Name:   "init",
		Usage:  "Pull Images from Docker Hub",
		Action: command.CmdInit,
		Flags:  []cli.Flag{},
	},
}

// CommandNotFound for non existent CLI command
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
