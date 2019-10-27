package dsp

import "github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor/processorbuiltin"

// PatchMultiply contains multiple copies of a patch
type PatchMultiply struct {
	Gain float32 `default:"0.5" min:"0" max:"1"`

	patch [processorbuiltin.MaxChannels]Patch
}

// Start - init multipled patches
func (p *PatchMultiply) Start(sampleRate int) {
	p.Gain = 0.5
	for i := 0; i < processorbuiltin.MaxChannels; i++ {
		p.patch[i].Start(sampleRate)
	}
}

// Process - produce sum off multiplied patches
func (p *PatchMultiply) Process(midiInput *processorbuiltin.MIDIInput) (output float32) {
	output = 0
	for i := 0; i < processorbuiltin.MaxChannels; i++ {
		freq, gate, trigger, _, _, _, _ := midiInput.Process(i)
		output += p.patch[i].Process(freq, gate, trigger)
	}
	midiInput.NextSample()

	output *= p.Gain
	return
}
