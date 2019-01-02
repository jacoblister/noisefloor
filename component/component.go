package component

import (
	. "github.com/jacoblister/noisefloor/common"
)

//Component interface
type Component interface {
	Start(sampleRate int)
	Stop()
	Process(samplesIn [][]AudioFloat, samplesOut [][]AudioFloat, midiIn []MidiEvent, midiOut []MidiEvent)
}
