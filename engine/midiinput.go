package engine

import (
	"math"

	. "github.com/jacoblister/noisefloor/common"
)

const maxChannels = 4
const maxNote = 127

// MIDIInput - MIDI to CV converter
type MIDIInput struct {
	channelNotes [maxNote]int
	channelData  [maxChannels][3]AudioFloat
	noteChannels map[int]int
	nextChannel  int
	triggerClear int
}

// Start initilizes the MIDIInput struct
func (m *MIDIInput) Start() {
	m.noteChannels = make(map[int]int)
}

// ProcessMIDI coverts MIDI events to CV signals
func (m *MIDIInput) ProcessMIDI(midiIn []MidiEvent) {
	len := len(midiIn)
	for i := 0; i < len; i++ {
		note := midiIn[i][1]
		velocity := midiIn[i][2]

		// note release or new note - free allocated channel
		noteChannel, ok := m.noteChannels[note]
		if ok {
			m.channelNotes[noteChannel] = 0
			m.channelData[noteChannel][1] = 0
			m.channelData[noteChannel][2] = -1
			delete(m.noteChannels, note)
		}

		if velocity > 0 {
			// Calculate frequency and level for note
			frequency := 220.0 * math.Pow(2.0, ((float64(note)-57)/12.0))
			level := velocity / 127.0

			// Allocate next free channel
			targetChannel := m.nextChannel
			for m.channelNotes[targetChannel] != 0 {
				targetChannel++
				if targetChannel >= maxChannels {
					targetChannel = 0
				}

				// If all channels active use current target
				if targetChannel == m.nextChannel {
					m.channelNotes[targetChannel] = 0
					delete(m.noteChannels, m.channelNotes[targetChannel])
				}
			}

			// set next channel, round robin
			m.nextChannel = targetChannel + 1
			if m.nextChannel >= maxChannels {
				m.nextChannel = 0
			}

			// set channel active
			m.channelNotes[targetChannel] = note
			m.channelData[targetChannel][0] = AudioFloat(frequency)
			m.channelData[targetChannel][1] = AudioFloat(level)
			m.channelData[targetChannel][2] = AudioFloat(level)
			m.noteChannels[note] = targetChannel
		}
	}
	m.triggerClear = 2
}

// Process returns to next sample CV data
func (m *MIDIInput) Process() [maxChannels][3]AudioFloat {
	if m.triggerClear > 0 {
		m.triggerClear--
		if m.triggerClear == 0 {
			// Clear triggers
			for i := 0; i < maxChannels; i++ {
				m.channelData[i][2] = 0
			}
		}
	}
	return m.channelData
}
