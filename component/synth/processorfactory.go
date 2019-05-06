package synth

import "github.com/jacoblister/noisefloor/component/synth/processor"

// ListProcessors returns a list of available processors
func ListProcessors() []string {
	return []string{"Envelope", "Gain", "Oscillator"}
}

//MakeProcessor generates a new processor by the given processor name
func MakeProcessor(name string) Processor {
	switch name {
	case "Envelope":
		return &processor.Envelope{}
	case "Gain":
		return &processor.Gain{}
	case "Oscillator":
		return &processor.Oscillator{}
	}

	return nil
}

func getProcessorInputs(p *Processor) []string {
	return []string{}
}

func getProcessorOutputs(p *Processor) []string {
	return []string{}
}

func getProcessorParameters(p *Processor) []ProcessorParameter {
	return []ProcessorParameter{}
}

func setProcessorParameter(p *Processor, name string, value float32) {
}
