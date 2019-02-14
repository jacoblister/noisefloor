package component

import (
	. "github.com/jacoblister/noisefloor/common"
)

// MIDILevels contains global current MIDI CC Levels
type MIDILevels [16][128]byte

//Component interface
type Component interface {
	Start(sampleRate int, midiLevels *MIDILevels)
	Stop()
	Process(samplesIn [][][]AudioFloat, samplesOut [][][]AudioFloat, midiIn []MidiEvent, midiOut []MidiEvent)
	Request(endpoint string, request string) string

	// Front end
	// ReactComponent() *ReactComponent
}

// Example stack
//
// -Output
// MIDIOutMapper (Active patch changes, tempo send, etc)
// Mixer         (Stereo downmix)
// Sampler
// Synth(s)
// Sequencer     (MIDI and looper insturctions)
// MIDIInMapper
// InputSelector (device -> 16 channels)
// -Input
