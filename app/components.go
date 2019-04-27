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
	return c.keyboard.Render()
}
