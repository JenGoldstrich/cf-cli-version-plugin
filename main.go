package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"fmt"
)

type DetailedVersionPlugin struct{}

func (p *DetailedVersionPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	apiVersion, _ := cliConnection.ApiVersion()
	fmt.Println("cf API Version: " + apiVersion)
	cliVersion, _ := cliConnection.CliCommandWithoutTerminalOutput("version")
	fmt.Println("cf CLI Version: " + cliVersion[0])
	return
}

func (p *DetailedVersionPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "DetailedVersionPlugin",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "detailed-version",
				HelpText: "Tells you who is cool!",
				UsageDetails: plugin.Usage{
					Usage: "cf detailed-version",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(DetailedVersionPlugin))
}
