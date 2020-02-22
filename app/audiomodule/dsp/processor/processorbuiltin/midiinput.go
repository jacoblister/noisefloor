package processorbuiltin

import (
	"math"

	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"
	"github.com/jacoblister/noisefloor/pkg/midi"
)

// MaxChannels it the maximum number of polyphonic channels
const MaxChannels = 8

const maxNote = 127
const pitchBendRange = 12

// MIDIInputMode enumerated type
type MIDIInputMode int

// polyphonic mode
const (
	MIDIInputModePoly MIDIInputMode = iota
	MIDIInputModeMono
	MIDIInputModeMPE
)

//MIDIInput is the MIDI to CV Converter
type MIDIInput struct {
	mode         MIDIInputMode
	channelNotes [MaxChannels][2]int     // note number, active
	channelData  [MaxChannels][3]float32 // freq, level, trigger
	pitchBend    float64
	noteChannels map[int]int
	nextChannel  int
	triggerClear int
}

// SetMono - force mono mode
func (m *MIDIInput) SetMono() {
	m.mode = MIDIInputModeMono
}

// Start - init midi input
func (m *MIDIInput) Start(sampleRate int, connectedMask int) {
	m.noteChannels = make(map[int]int)
}

// ProcessMIDI coverts MIDI events to CV signals
func (m *MIDIInput) ProcessMIDI(midiIn []midi.Event) {
	monoOffEvent := false

	// Clear all triggers
	for i := 0; i < MaxChannels; i++ {
		m.channelData[i][2] = 0
	}

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
				if targetChannel >= MaxChannels {
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
			if m.nextChannel >= MaxChannels {
				m.nextChannel = 0
			}

			// Calculate frequency and level for note
			frequency := 220.0 * math.Pow(2.0, ((float64(note)-57+m.pitchBend)/12.0))
			level := float32(velocity) / 127.0
			trigger := level

			// mono mode - force first channel and don't retrigger if active
			if m.mode == MIDIInputModeMono {
				targetChannel = 0
				if monoOffEvent {
					trigger = 0
				}
			}

			// set channel active
			m.channelNotes[targetChannel][0] = note
			m.channelNotes[targetChannel][1] = 1
			m.channelData[targetChannel][0] = float32(frequency)
			m.channelData[targetChannel][1] = float32(level)
			m.channelData[targetChannel][2] = float32(trigger)
			m.noteChannels[note] = targetChannel
		case midi.NoteOffEvent:
			note := event.Note

			// note release - update allocated channel
			noteChannel, ok := m.noteChannels[note]
			if ok {
				m.channelNotes[noteChannel][1] = 0
				m.channelData[noteChannel][1] = 0
				m.channelData[noteChannel][2] = 0
				delete(m.noteChannels, note)
			}

			monoOffEvent = true
		case midi.PitchBendEvent:
			m.pitchBend = event.Normailzed() * pitchBendRange
			for i := 0; i < MaxChannels; i++ {
				note := m.channelNotes[i][0]
				frequency := 220.0 * math.Pow(2.0, ((float64(note)-57+m.pitchBend)/12.0))
				m.channelData[i][0] = float32(frequency)
			}
		}
	}
	m.triggerClear = 2
}

// Process - produce next sample
func (m *MIDIInput) Process(i int) (frequency float32, gate float32, trigger float32,
	aftertouch float32, slide float32, release float32, channel float32) {
	channelData := &m.channelData[i]
	return channelData[0], channelData[1], channelData[2], 0.0, 0.0, 0.0, float32(i)
}

// Definition exports processor definition
func (m *MIDIInput) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "MIDIInput", []string{}, []string{"Frq", "Lvl", "Trg", "Aft", "Sld", "Rel", "Chn"},
		[]processor.Parameter{}
}

//ProcessArgs calls process with args as an array
func (m *MIDIInput) ProcessArgs(in []float32) (output []float32) {
	frequency, gate, trigger, aftertouch, slide, release, channel := m.Process(0)
	return []float32{frequency, gate, trigger, aftertouch, slide, release, channel}
}

//ProcessSamples calls process with an array of input/output samples
func (m *MIDIInput) ProcessSamples(in [][]float32, length int) (output [][]float32) {
	output = make([][]float32, 7)
	output[0] = make([]float32, length)
	output[1] = make([]float32, length)
	output[2] = make([]float32, length)
	output[3] = make([]float32, length)
	output[4] = make([]float32, length)
	output[5] = make([]float32, length)
	output[6] = make([]float32, length)
	for i := 0; i < length; i++ {
		output[0][i], output[1][i], output[2][i], output[3][i], output[4][i], output[5][i], output[6][i] = m.Process(0)
	}
	return
}

//SetParameter set a single processor parameter
func (m *MIDIInput) SetParameter(index int, value float32) {}
