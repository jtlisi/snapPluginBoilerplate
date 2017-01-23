package snapPluginBoilerplate

import "github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

// ExamplePublisher is a testing publisher.
type ExamplePublisher struct {
}

/*
	GetConfigPolicy() returns the configPolicy for your plugin.
	A config policy is how users can provide configuration info to
	plugin. Here you define what sorts of config info your plugin
	needs and/or requires.
*/
func (f ExamplePublisher) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	return *policy, nil
}

// Publish test publish function
func (f ExamplePublisher) Publish(mts []plugin.Metric, cfg plugin.Config) error {
	return nil
}
