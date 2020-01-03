package processor

import (
	"errors"
)

//Parameter is a single paramter with a range and default
type Parameter struct {
	Name    string
	Default float32
	Min     float32
	Max     float32
	Value   float32
}

//Processor interface
type Processor interface {
	Start(sampleRate int)
	// Stop()
	ProcessArgs([]float32) []float32
	ProcessSamples([][]float32, int) [][]float32
	Definition() (name string, inputs []string, outputs []string, parameters []Parameter)
	SetParameter(index int, value float32)
}

//GetProcessorInputIndex returns the index for a given import port
func GetProcessorInputIndex(processor Processor, port string) (int, error) {
	_, inputs, _, _ := processor.Definition()

	for i := 0; i < len(inputs); i++ {
		if inputs[i] == port {
			return i, nil
		}
	}
	return 0, errors.New("processor input not found")
}

//GetProcessorOutputIndex returns the index for a given output port
func GetProcessorOutputIndex(processor Processor, port string) (int, error) {
	_, _, outputs, _ := processor.Definition()

	for i := 0; i < len(outputs); i++ {
		if outputs[i] == port {
			return i, nil
		}
	}
	return 0, errors.New("processor input not found")
}

//GetProcessorParameterIndex returns the index for a given parameter
func GetProcessorParameterIndex(processor Processor, parameter string) (int, error) {
	_, _, _, parameters := processor.Definition()

	for i := 0; i < len(parameters); i++ {
		if parameters[i].Name == parameter {
			return i, nil
		}
	}
	return 0, errors.New("processor parameter not found")
}

//SetProcessorDefaults sets processor parametrers to defaults
func SetProcessorDefaults(processor Processor) {
	_, _, _, paramters := processor.Definition()
	for i := 0; i < len(paramters); i++ {
		processor.SetParameter(i, paramters[i].Default)
	}
}

// Definition is a configured processor with screen coordinates
type Definition struct {
	X         int
	Y         int
	Name      string
	Processor Processor
}

// MaxConnectors get this maximum of input and output connectors
func (d *Definition) MaxConnectors() int {
	_, procInputs, procOutputs, _ := d.Processor.Definition()
	result := len(procInputs)
	if len(procOutputs) > result {
		result = len(procOutputs)
	}
	return result
}

// GetName gets the ProcessorDefinition name, defaulting to the Processor Name if not provided
func (d *Definition) GetName() string {
	if len(d.Name) > 0 {
		return d.Name
	}
	procDefName, _, _, _ := d.Processor.Definition()
	return procDefName
}
