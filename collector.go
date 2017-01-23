package snapPluginBoilerplate

import "github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

// ExampleCollector implementation used for testing
type ExampleCollector struct {
}

/*  CollectMetrics collects metrics for testing.
CollectMetrics() will be called by Snap when a task that collects one of the metrics returned from this plugins
GetMetricTypes() is started. The input will include a slice of all the metric types being collected.
The output is the collected metrics as plugin.Metric and an error.
*/
func (ExampleCollector) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}
	return metrics, nil
}

/*
	GetMetricTypes returns metric types for testing.
	GetMetricTypes() will be called when your plugin is loaded in order to populate the metric catalog(where snaps stores all
	available metrics).
	Config info is passed in. This config information would come from global config snap settings.
	The metrics returned will be advertised to users who list all the metrics and will become targetable by tasks.
*/
func (ExampleCollector) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}
	return metrics, nil
}

/*
	GetConfigPolicy() returns the configPolicy for your plugin.
	A config policy is how users can provide configuration info to
	plugin. Here you define what sorts of config info your plugin
	needs and/or requires.
*/
func (ExampleCollector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	return *policy, nil
}
