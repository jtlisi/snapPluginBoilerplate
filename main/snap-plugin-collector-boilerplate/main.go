package main

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"github.com/jtlisi/snapPluginBoilerplate"
)

const (
	pluginName    = "example-collector"
	pluginVersion = 1
)

func main() {
	plugin.StartProcessor(snapPluginBoilerplate.ExampleCollector{}, pluginName, pluginVersion)
}
