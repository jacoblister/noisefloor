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
	Name      string
	Processor Processor
}

// MaxConnectors get this maximum of input and output connectors
func (p *ProcessorDefinition) MaxConnectors() int {
	_, procInputs, procOutputs := p.Processor.Definition()
	result := len(procInputs)
	if len(procOutputs) > result {
		result = len(procOutputs)
	}
	return result
}
