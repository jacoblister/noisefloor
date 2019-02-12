package common

// MidiEvent struct with time and data
type MidiEvent []int

type NewMidiEvent struct {
	Time int
	Data []byte
}

func MakeMidiEvent(Time int, Data []byte) (result *NewMidiEvent) {
	return &NewMidiEvent{Time, Data}
}
