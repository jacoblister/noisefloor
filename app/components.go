package app

import (
	"github.com/jacoblister/noisefloor/component/onscreenkeyboard"
	"github.com/jacoblister/noisefloor/component/synth"
	"github.com/jacoblister/noisefloor/midi"
	"github.com/jacoblister/noisefloor/vdom"
)

type components struct {
	keyboard    onscreenkeyboard.Keyboard
	synthEngine synth.Engine
}

// Start begin the main application audio processing
func (c *components) Start(sampleRate int) {
	c.keyboard.Start(sampleRate)
	c.synthEngine.Start(sampleRate)
}

// Stop closes the main application audio processing
func (c *components) Stop() {
	c.keyboard.Stop()
	c.synthEngine.Stop()
}

// Process process a block of audio/midi
func (c *components) Process(samplesIn [][]float32, midiIn []midi.Event) (samplesOut [][]float32, midiOut []midi.Event) {
	samples, midi := samplesIn, midiIn

	samples, midi = c.keyboard.Process(samples, midi)
	samples, midi = c.synthEngine.Process(samples, midi)

	return samples, midi
}

// Render returns the main view
func (c *components) Render() vdom.Element {
	elem := vdom.MakeElement("svg",
		"id", "root",
		"xmlns", "http://www.w3.org/2000/svg",
		"style", "width:100%;height:100%;position:fixed;top:0;left:0;bottom:0;right:0;",
		vdom.MakeElement("g", "transform", "translate(0,0)", &c.keyboard),
	)
	return elem
}
