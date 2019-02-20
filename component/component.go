package component

import (
	"github.com/jacoblister/noisefloor/common/midi"
)

// MIDILevels contains global current MIDI CC Levels
type MIDILevels [16][128]byte

// AudioProcessor is a frame based audio/midi processor
type AudioProcessor interface {
	Start(sampleRate int)
	Stop()
	Process(samplesIn [][]float32, samplesOut [][]float32, midiIn []midi.Event, midiOut *[]midi.Event)
}

// AudioBusProcessor is a frame based audio/midi bus processor (multi channel mono/stereo)
type AudioBusProcessor interface {
	Start(sampleRate int, midiLevels *MIDILevels)
	Stop()
	Process(samplesIn [][][]float32, samplesOut [][][]float32, midiIn []midi.Event, midiOut *[]midi.Event)
}

// AudioProcessorFrontend is the front end to an AudioProcessor (belongs elsewhere)
type AudioProcessorFrontend interface {
	// Front end
	// ReactComponent() *ReactComponent ???
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
