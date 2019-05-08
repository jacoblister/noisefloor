package synth

//Processor interface
type Processor interface {
	Start(sampleRate int)
	// Stop()
	// Process(vars ...[]*AudioFloat)
}

// ProcessorDefinition is a configured processor with coordinates
type processorDefinition struct {
	x         int
	y         int
	processor Processor
}
