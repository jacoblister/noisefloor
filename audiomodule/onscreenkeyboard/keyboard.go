package onscreenkeyboard

import "github.com/jacoblister/noisefloor/pkg/midi"

const keyMax = 127
const velocityMax = 127

// NoteEventFunc is a callback for handling note on/off events
type NoteEventFunc func(keyNumber int, keyDown bool)

// Keyboard is the onscreen keyboard processor
type Keyboard struct {
	MidiEvents    []midi.Event
	Keydown       [keyMax]bool
	noteEventFunc NoteEventFunc
}

// SetNoteEventFunc sets a notify callback when notes occurs
func (k *Keyboard) SetNoteEventFunc(noteEventFunc NoteEventFunc) {
	k.noteEventFunc = noteEventFunc
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

	// notify front end if registered
	if k.noteEventFunc != nil {
		for i := 0; i < len(midiIn); i++ {
			switch event := midiIn[i].(type) {
			case midi.NoteOnEvent:
				k.noteEventFunc(event.Note, true)
			case midi.NoteOffEvent:
				k.noteEventFunc(event.Note, false)
			}
		}
	}

	midiOut = append(midiIn, k.MidiEvents...)
	k.MidiEvents = nil

	return samplesOut, midiOut
}
