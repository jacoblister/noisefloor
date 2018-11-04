package engine

import (
	. "github.com/jacoblister/noisefloor/common"
)

// PatchMultiply contains multiple copies of a patch
type PatchMultiply struct {
	Gain AudioFloat `default:"0.5" min:"0" max:"1"`

	patch [maxChannels]Patch
}

// Start - init multipled patches
func (p *PatchMultiply) Start(sampleRate int) {
	p.Gain = 0.5
	for i := 0; i < maxChannels; i++ {
		p.patch[i].Start(sampleRate)
	}
}

// Process - produce sum off multiplied patches
func (p *PatchMultiply) Process(freqs [maxChannels][3]AudioFloat) (output AudioFloat) {
	output = 0
	for i := 0; i < maxChannels; i++ {
		output += p.patch[i].Process(freqs[i][0], freqs[i][1], freqs[i][2])
	}

	output *= p.Gain
	return
}
