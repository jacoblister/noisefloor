package processor

import (
	. "github.com/jacoblister/noisefloor/common"
)

// Phase enumerated type
type Phase int

// Basic wave shapes
const (
	Inactive Phase = iota
	Attack
	Decay
	Sustain
	Release
)

// Envelope - ADSR envelope generator
type Envelope struct {
	Attack  AudioFloat `default:"2" min:"0" max:"10000"`
	Decay   AudioFloat `default:"100" min:"0" max:"10000"`
	Sustain AudioFloat `default:"0.5" min:"0" max:"1"`
	Release AudioFloat `default:"1000" min:"0" max:"10000"`

	sampleRate AudioFloat
	output     AudioFloat
	phase      Phase
	delta      AudioFloat
}

// Start - init envelope generator
func (e *Envelope) Start(sampleRate int) {
	e.sampleRate = AudioFloat(sampleRate)
	e.Attack = 2
	e.Decay = 100
	e.Sustain = 0.5
	e.Release = 1000
}

// Process - produce next sample
func (e *Envelope) Process(gate AudioFloat, trigger AudioFloat) (output AudioFloat) {
	if trigger > 0 {
		e.output = 0
		e.delta = (1000 / e.Attack) / e.sampleRate
		e.phase = Attack
	}

	switch phase := e.phase; phase {
	case Attack:
		e.output += e.delta
		if e.output > 1 {
			e.delta = (1000 / e.Decay) / e.sampleRate
			e.phase = Decay
		}
	case Decay:
		e.output -= e.delta
		if e.output < e.Sustain {
			e.phase = Sustain
		}
	case Sustain:
		if gate == 0 {
			e.delta = (1000 / e.Release) / e.sampleRate
			e.phase = Release
		}
	case Release:
		e.output -= e.delta
		if e.output < 0 {
			e.output = 0
			e.phase = Inactive
		}
	}

	output = e.output
	return
}
