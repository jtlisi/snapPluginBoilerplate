package snapPluginBoilerplate

import "github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

// ExampleProcessor test processor
type ExampleProcessor struct {
}

// Process test process function
func (r ExampleProcessor) Process(mts []plugin.Metric, cfg plugin.Config) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}
	return metrics, nil
}

/*
	GetConfigPolicy() returns the configPolicy for your plugin.
	A config policy is how users can provide configuration info to
	plugin. Here you define what sorts of config info your plugin
	needs and/or requires.
*/
func (r ExampleProcessor) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	return *policy, nil
}
