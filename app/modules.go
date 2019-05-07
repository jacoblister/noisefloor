package app

import (
	"github.com/jacoblister/noisefloor/audiomodule/onscreenkeyboard"
	"github.com/jacoblister/noisefloor/audiomodule/synth"
	"github.com/jacoblister/noisefloor/midi"
	"github.com/jacoblister/noisefloor/vdom"
)

type modules struct {
	keyboard    onscreenkeyboard.Keyboard
	synthEngine synth.Engine
}

// Start begin the main application audio processing
func (c *modules) Start(sampleRate int) {
	c.keyboard.Start(sampleRate)
	c.synthEngine.Start(sampleRate)
}

// Stop closes the main application audio processing
func (c *modules) Stop() {
	c.keyboard.Stop()
	c.synthEngine.Stop()
}

// Process process a block of audio/midi
func (c *modules) Process(samplesIn [][]float32, midiIn []midi.Event) (samplesOut [][]float32, midiOut []midi.Event) {
	samples, midi := samplesIn, midiIn

	samples, midi = c.keyboard.Process(samples, midi)
	samples, midi = c.synthEngine.Process(samples, midi)

	return samples, midi
}

// Render returns the main view
func (c *modules) Render() vdom.Element {
	elem := vdom.MakeElement("svg",
		"id", "root",
		"xmlns", "http://www.w3.org/2000/svg",
		"style", "width:100%;height:100%;position:fixed;top:0;left:0;bottom:0;right:0;",
		// vdom.MakeElement("g", "transform", "translate(0,0)", &c.synthEngine),
		vdom.MakeElement("g", "transform", "translate(0,0)", &c.keyboard),
	)
	return elem
}
