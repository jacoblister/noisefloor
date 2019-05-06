package synth

import (
	"math"

	"github.com/jacoblister/noisefloor/midi"
)

const maxChannels = 8
const maxNote = 127
const pitchBendRange = 12

// MIDIInput - MIDI to CV converter
type MIDIInput struct {
	channelNotes [maxChannels][2]int     // note number, active
	channelData  [maxChannels][3]float32 // freq, level, trigger
	pitchBend    float64
	noteChannels map[int]int
	nextChannel  int
	triggerClear int
}

// Start initilizes the MIDIInput struct
func (m *MIDIInput) Start() {
	m.noteChannels = make(map[int]int)
}

// ProcessMIDI coverts MIDI events to CV signals
func (m *MIDIInput) ProcessMIDI(midiIn []midi.Event) {
	len := len(midiIn)
	for i := 0; i < len; i++ {
		switch event := midiIn[i].(type) {
		case midi.NoteOnEvent:
			note := event.Note
			velocity := event.Velocity

			// Allocate next free channel
			targetChannel := m.nextChannel
			for m.channelNotes[targetChannel][1] != 0 {
				targetChannel++
				if targetChannel >= maxChannels {
					targetChannel = 0
				}

				// If all channels active use current target
				if targetChannel == m.nextChannel {
					m.channelNotes[targetChannel][1] = 0
					delete(m.noteChannels, m.channelNotes[targetChannel][0])
				}
			}

			// set next channel, round robin
			m.nextChannel = targetChannel + 1
			if m.nextChannel >= maxChannels {
				m.nextChannel = 0
			}

			// Calculate frequency and level for note
			frequency := 220.0 * math.Pow(2.0, ((float64(note)-57+m.pitchBend)/12.0))
			level := float32(velocity) / 127.0

			// set channel active
			m.channelNotes[targetChannel][0] = note
			m.channelNotes[targetChannel][1] = 1
			m.channelData[targetChannel][0] = float32(frequency)
			m.channelData[targetChannel][1] = float32(level)
			m.channelData[targetChannel][2] = float32(level)
			m.noteChannels[note] = targetChannel
		case midi.NoteOffEvent:
			note := event.Note

			// note release - update allocated channel
			noteChannel, ok := m.noteChannels[note]
			if ok {
				m.channelNotes[noteChannel][1] = 0
				m.channelData[noteChannel][1] = 0
				m.channelData[noteChannel][2] = -1
				delete(m.noteChannels, note)
			}
		case midi.PitchBendEvent:
			m.pitchBend = event.Normailzed() * pitchBendRange
			for i := 0; i < maxChannels; i++ {
				note := m.channelNotes[i][0]
				frequency := 220.0 * math.Pow(2.0, ((float64(note)-57+m.pitchBend)/12.0))
				m.channelData[i][0] = float32(frequency)
			}
		}
	}
	m.triggerClear = 2
}

// Process returns to next sample CV data
func (m *MIDIInput) Process() *[maxChannels][3]float32 {
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
