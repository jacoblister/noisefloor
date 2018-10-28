package processor

import (
	. "github.com/jacoblister/noisefloor/common"
)

// Gain - linear of exponential gain
type Gain struct {
	Exponential bool
}

// Start - init envelope generator
func (g *Gain) Start(sampleRate int) {}

// Process - produce next sample
func (g *Gain) Process(input AudioFloat, gain AudioFloat) (output AudioFloat) {
	output = input * gain
	return
}
