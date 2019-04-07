package app

import (
	"github.com/jacoblister/noisefloor/midi"
	"github.com/jacoblister/noisefloor/vdom"
)

// Start begin the main application
func Start(sampleRate int) {}

// Stop closes the main application
func Stop() {}

// Process process a block of audio/midi
func Process(samplesIn [][]float32, samplesOut [][]float32, midiIn []midi.Event, midiOut *[]midi.Event) {

}

// Render returns the main view
func Render() vdom.Element {
	return vdom.MakeElement("div")
}
