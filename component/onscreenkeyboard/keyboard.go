package onscreenkeyboard

import "github.com/jacoblister/noisefloor/common/midi"

type KeyboardViewer interface {
	SendEvent(event midi.EventData)
}

// Keyboard is the onscreen keyboard processor
type Keyboard struct {
	MidiEvents []midi.Event
}

// Start initilized the component, with a specified sampling rate
func (k *Keyboard) Start(sampleRate int) {
}

// Stop suspends the component
func (k *Keyboard) Stop() {
}

// Process processes a block of samples and midi events
func (k *Keyboard) Process(samplesIn [][][]float32, samplesOut [][][]float32, midiIn []midi.Event, midiOut []midi.Event) {
	midiOut = k.MidiEvents
	k.MidiEvents = nil
}

func (k *Keyboard) SendEvent(event midi.EventData) {
	k.MidiEvents = append(k.MidiEvents, midi.MakeMidiEvent(event.Time, event.Data))
}
