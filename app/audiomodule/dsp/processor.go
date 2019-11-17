package dsp

import (
	"errors"

	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"
)

//Processor interface
type Processor interface {
	Start(sampleRate int)
	// Stop()
	ProcessArgs([]float32) []float32
	ProcessSamples([][]float32, int) [][]float32
	Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter)
	SetParameter(index int, value float32)
}

func getProcessorInputIndex(processor Processor, name string) (int, error) {
	_, inputs, _, _ := processor.Definition()

	for i := 0; i < len(inputs); i++ {
		if inputs[i] == name {
			return i, nil
		}
	}
	return 0, errors.New("processor input not found")
}

func getProcessorOutputIndex(processor Processor, name string) (int, error) {
	_, _, outputs, _ := processor.Definition()

	for i := 0; i < len(outputs); i++ {
		if outputs[i] == name {
			return i, nil
		}
	}
	return 0, errors.New("processor input not found")
}

func getProcessorParameterIndex(processor Processor, name string) (int, error) {
	_, _, _, parameters := processor.Definition()

	for i := 0; i < len(parameters); i++ {
		if parameters[i].Name == name {
			return i, nil
		}
	}
	return 0, errors.New("processor parameter not found")
}

func setProcessorDefaults(processor Processor) {
	_, _, _, paramters := processor.Definition()
	for i := 0; i < len(paramters); i++ {
		processor.SetParameter(i, paramters[i].Default)
	}
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
