package processor

// Gain - linear of exponential gain
type Gain struct {
	Exponential bool
}

// Start - init envelope generator
func (g *Gain) Start(sampleRate int) {}

// Process - produce next sample
func (g *Gain) Process(input float32, gain float32) (output float32) {
	output = input * gain
	return
}
