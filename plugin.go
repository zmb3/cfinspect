package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"code.cloudfoundry.org/cli/plugin"
)

const (
	name           = "inspect"
	hostCommand    = "host"
	dropletCommand = "droplet"
)

type inspectPlugin struct {
	out io.Writer
}

func newInspectPlugin() *inspectPlugin {
	p := &inspectPlugin{out: os.Stdout}
	return p
}

func (p *inspectPlugin) Run(conn plugin.CliConnection, args []string) {
	if args[0] != name {
		return
	}

	args = args[1:]
	var err error
	switch args[0] {
	case hostCommand:
		err = p.host(conn, args[1:])
	case dropletCommand:
		err = p.droplet(conn, args[1:])
	default:
		err = fmt.Errorf("unknown subcommand %s", args[0])
	}

	// TODO: more testable error handling
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func (p *inspectPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name:          name,
		Version:       plugin.VersionType{Major: 0, Minor: 1, Build: 0},
		MinCliVersion: plugin.VersionType{Major: 6, Minor: 7, Build: 0},
		Commands: []plugin.Command{
			{
				Name:         "inspect",
				HelpText:     "Inspect various CF metadata",
				UsageDetails: plugin.Usage{Usage: "inspect\n\tcf inspect <command> [args...]"},
			},
		},
	}
}
