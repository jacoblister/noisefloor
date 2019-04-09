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

var nf noiseFloor

// Start begin the main application audio processing
func Start(sampleRate int) {}

// Stop closes the main application audio processing
func Stop() {}

// Process process a block of audio/midi
func Process(samplesIn [][]float32, samplesOut [][]float32, midiIn []midi.Event, midiOut *[]midi.Event) {
	nf.keyboard.Process(samplesIn, samplesOut, midiIn, midiOut)
	nf.synthEngine.Process(samplesIn, samplesOut, midiIn, midiOut)
}

// Render returns the main view
func Render() vdom.Element {
	return nf.keyboard.Render()
}
