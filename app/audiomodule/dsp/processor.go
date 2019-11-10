package dsp

import "github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"

//Processor interface
type Processor interface {
	Start(sampleRate int)
	// Stop()
	ProcessArgs([]float32) []float32
	ProcessSamples([][]float32, int) [][]float32
	Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter)
	SetParameter(index int, value float32)
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
	_, procInputs, procOutputs, _ := p.Processor.Definition()
	result := len(procInputs)
	if len(procOutputs) > result {
		result = len(procOutputs)
	}
	return result
}

// GetName gets the ProcessorDefinition name, defaulting to the Processor Name if not provided
func (p *ProcessorDefinition) GetName() string {
	if len(p.Name) > 0 {
		return p.Name
	}
	procDefName, _, _, _ := p.Processor.Definition()
	return procDefName
}
