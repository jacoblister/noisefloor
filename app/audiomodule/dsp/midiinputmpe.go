package dsp

import (
	"math"

	"github.com/jacoblister/noisefloor/pkg/midi"
)

const mpePitchBendRange = 6

// MIDIInputMPE - MIDI to CV converter for MPR instruments
type MIDIInputMPE struct {
	channelNotes [maxChannels][2]int     // note number, active
	channelData  [maxChannels][3]float32 // freq, level, trigger
	pitchBend    float64
	triggerClear int
}

// Start initilizes the MIDIInput struct
func (m *MIDIInputMPE) Start() {
}

// ProcessMIDI coverts MIDI events to CV signals
func (m *MIDIInputMPE) ProcessMIDI(midiIn []midi.Event) {
	len := len(midiIn)
	for i := 0; i < len; i++ {
		targetChannel := midiIn[i].Generic().Channel - 1
		if targetChannel >= maxChannels {
			continue
		}

		switch event := midiIn[i].(type) {
		case midi.NoteOnEvent:
			note := event.Note
			velocity := event.Velocity

			// Calculate frequency and level for note
			frequency := 220.0 * math.Pow(2.0, ((float64(note)-57+m.pitchBend)/12.0))
			level := float32(velocity) / 127.0

			// set channel active
			m.channelNotes[targetChannel][0] = note
			m.channelNotes[targetChannel][1] = 1
			m.channelData[targetChannel][0] = float32(frequency)
			m.channelData[targetChannel][1] = float32(level)
			m.channelData[targetChannel][2] = float32(level)
		case midi.NoteOffEvent:
			// note release - update allocated channel
			m.channelNotes[targetChannel][1] = 0
			m.channelData[targetChannel][1] = 0
			m.channelData[targetChannel][2] = -1
		case midi.PitchBendEvent:
			m.pitchBend = event.Normailzed() * pitchBendRange

			note := m.channelNotes[targetChannel][0]
			frequency := 220.0 * math.Pow(2.0, ((float64(note)-57+m.pitchBend)/12.0))
			m.channelData[targetChannel][0] = float32(frequency)
		}
	}
	m.triggerClear = 2
}

// Process returns to next sample CV data
func (m *MIDIInputMPE) Process() *[maxChannels][3]float32 {
	if m.triggerClear > 0 {
		m.triggerClear--
		if m.triggerClear == 0 {
			// Clear triggers
			for i := 0; i < maxChannels; i++ {
				m.channelData[i][2] = 0
			}
		}
	}
	return &m.channelData
}
