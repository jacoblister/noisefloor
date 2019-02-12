package onscreenkeyboard

import (
	. "github.com/jacoblister/noisefloor/common"
)

// Keyboard is the onscreen keyboard processor
type Keyboard struct {
}

// Start initilized the component, with a specified sampling rate
func (k *Keyboard) Start(sampleRate int) {
}

// Stop suspends the component
func (k *Keyboard) Stop() {
}

// Process processes a block of samples and midi events
func (k *Keyboard) Process(samplesIn [][][]AudioFloat, samplesOut [][][]AudioFloat, midiIn []MidiEvent, midiOut []MidiEvent) {
}
