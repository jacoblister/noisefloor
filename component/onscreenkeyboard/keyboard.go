package onscreenkeyboard

import "github.com/jacoblister/noisefloor/midi"

const keyMax = 127
const velocityMax = 127

// Keyboard is the onscreen keyboard processor
type Keyboard struct {
	keydown    [keyMax]bool
	MidiEvents []midi.Event
}

// Start initilized the component, with a specified sampling rate
func (k *Keyboard) Start(sampleRate int) {
}

// Stop suspends the component
func (k *Keyboard) Stop() {
}

// Process processes a block of samples and midi events
func (k *Keyboard) Process(samplesIn [][]float32, midiIn []midi.Event) (samplesOut [][]float32, midiOut []midi.Event) {
	samplesOut = samplesIn
	midiOut = append(midiIn, k.MidiEvents...)
	k.MidiEvents = nil

	return samplesOut, midiOut
}
