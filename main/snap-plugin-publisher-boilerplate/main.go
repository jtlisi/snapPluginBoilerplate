package main

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"github.com/jtlisi/snapPluginBoilerplate"
)

const (
	pluginName    = "example-publisher"
	pluginVersion = 1
)

func main() {
	plugin.StartProcessor(snapPluginBoilerplate.ExampleProcessor{}, pluginName, pluginVersion)
}
