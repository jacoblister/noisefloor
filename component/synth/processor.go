package synth

//Processor interface
type Processor interface {
	Start(sampleRate int)
	// Stop()
	// Process(vars ...[]*AudioFloat)
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
