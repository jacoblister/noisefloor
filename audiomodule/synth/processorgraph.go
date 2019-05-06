package synth

// ProcessorGraph is a graph of processors and connectors, plus exported parameter map
type ProcessorGraph struct {
	name          string
	processorList []Processor
	connectorList []Connector
}
