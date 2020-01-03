package processorbasic

// Gain - linear of exponential gain
type Gain struct {
	Level       float32 `default:"1" min:"0" max:"2"`
	Exponential bool
}

// Start - init envelope generator
func (g *Gain) Start(sampleRate int) {}

// Process - produce next sample
func (g *Gain) Process(input float32, gain float32) (output float32) {
	output = input * gain * g.Level
	return
}
