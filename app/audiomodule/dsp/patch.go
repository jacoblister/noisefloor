package dsp

import (
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor/processorbasic"
)

// Patch is a simple minimal example patch
type Patch struct {
	oscillator processorbasic.Oscillator
	envelope   processorbasic.Envelope
	gain       processorbasic.Gain
}

// Start - init patch
func (p *Patch) Start(sampleRate int) {
	p.oscillator.Waveform = processorbasic.Square
	p.oscillator.Start(sampleRate, 0)

	p.envelope.Attack = 2
	p.envelope.Decay = 100
	p.envelope.Sustain = 0.75
	p.envelope.Release = 1000
	p.envelope.Start(sampleRate, 0)

	p.gain.Level = 1
	p.gain.Start(sampleRate, 0)
}

// Process - produce next sample
func (p *Patch) Process(freq float32, gate float32, trigger float32) (output float32) {
	sample := p.oscillator.Process(freq)
	env := p.envelope.Process(gate, trigger)

	output = p.gain.Process(sample, env)
	return
}
