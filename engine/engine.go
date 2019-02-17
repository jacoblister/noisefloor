package engine

import (
    "github.com/jacoblister/noisefloor/common/midi"
	"github.com/jacoblister/noisefloor/engine/processor"
)

// Engine - DSP processing engine
type Engine struct {
}

var midiinput MIDIInput
var patch PatchMultiply
var osc processor.Oscillator

// Start initilized the engine, with a specified sampling rate
func Start(sampleRate int) {
	println("do DSP start, sample rate:", sampleRate)
	midiinput.Start()
	patch.Start(sampleRate)

	osc.Start(sampleRate)
	osc.Waveform = processor.Sin
	osc.Freq = 5
}

// Stop suspends the engine
func Stop() {
	println("do DSP stop")
}

// Process processes a block of samples and midi events
func Process(samplesIn [][]float32, samplesOut [][]float32, midiIn []midi.Event, midiOut *[]midi.Event) {
	midiinput.ProcessMIDI(midiIn)

	var len = len(samplesOut[0])
	for i := 0; i < len; i++ {
		freqs := midiinput.Process()
		var sample = patch.Process(freqs)

		mic := samplesIn[0][i] * 500
		mod := osc.Process()
		if mod < 0 {
			mod = 0
		}
		mic *= mod
		sample += mic

		samplesOut[0][i] = sample
		samplesOut[1][i] = sample
	}
}
