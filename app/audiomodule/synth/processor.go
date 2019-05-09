package synth

//Processor interface
type Processor interface {
	Start(sampleRate int)
	// Stop()
	// Process(vars ...[]*AudioFloat)
}

// ProcessorDefinition is a configured processor with screen coordinates
type ProcessorDefinition struct {
	X         int
	Y         int
	Processor Processor
}
