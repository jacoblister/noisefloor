package processorbuiltin

import "github.com/jacoblister/noisefloor/pkg/midi"

//MIDIInput is the MIDI to CV Converter
type MIDIInput struct{}

// Start - init envelope generator
func (m *MIDIInput) Start(sampleRate int) {}

// ProcessMIDI coverts MIDI events to CV signals
func (m *MIDIInput) ProcessMIDI(midiIn []midi.Event) {}

// Process - produce next sample
func (m *MIDIInput) Process(i int) (frequency float32, gate float32, trigger float32, channel float32) {
	return
}

// Definition exports processor definition
func (m *MIDIInput) Definition() (name string, inputs []string, outputs []string) {
	return "MIDIInput", []string{}, []string{"frequency", "level", "trigger", "channel"}
}

//ProcessArray calls process with an array of input/output samples
func (m *MIDIInput) ProcessArray(in []float32) (output []float32) {
	return
}
