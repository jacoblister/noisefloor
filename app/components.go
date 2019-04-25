package app

import (
	"github.com/jacoblister/noisefloor/component/onscreenkeyboard"
	"github.com/jacoblister/noisefloor/component/synth"
	"github.com/jacoblister/noisefloor/midi"
	"github.com/jacoblister/noisefloor/vdom"
)

type noiseFloor struct {
	keyboard    onscreenkeyboard.Keyboard
	synthEngine synth.Engine
}

// Start begin the main application audio processing
func (nf *noiseFloor) Start(sampleRate int) {}

// Stop closes the main application audio processing
func (nf *noiseFloor) Stop() {}

// Process process a block of audio/midi
func (nf *noiseFloor) Process(samplesIn [][]float32, midiIn []midi.Event) (samplesOut [][]float32, midiOut []midi.Event) {
	samples, midi := samplesIn, midiIn

	samples, midi = nf.keyboard.Process(samples, midi)
	samples, midi = nf.synthEngine.Process(samples, midi)

	return samples, midi
}

// Render returns the main view
func (nf *noiseFloor) Render() vdom.Element {
	return nf.keyboard.Render()
}
