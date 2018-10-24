package processor

import (
	. "github.com/jacoblister/noisefloor/common"
)

// Gain - linear of exponential gain
type Gain struct {
	Exponential bool
}

// Process - produce next sample
func (o *Gain) Process(input AudioFloat, gain AudioFloat) (output AudioFloat) {
	output = input * gain
	return
}
