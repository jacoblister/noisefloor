package dsp

//Processor interface
type Processor interface {
	Start(sampleRate int)
	// Stop()
	ProcessArray([]float32) []float32
	Definition() (name string, inputs []string, outputs []string)
}

// ProcessorDefinition is a configured processor with screen coordinates
type ProcessorDefinition struct {
	X         int
	Y         int
	Processor Processor
}
