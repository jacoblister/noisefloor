package engine

import (
	"github.com/jacoblister/noisefloor/common/midi"
	"github.com/jacoblister/noisefloor/engine/processor"
)

// Engine - DSP processing engine
type Engine struct {
	midiinput MIDIInput
	patch     PatchMultiply
	osc       processor.Oscillator
}

// Start initilized the engine, with a specified sampling rate
func (e *Engine) Start(sampleRate int) {
	println("do DSP start, sample rate:", sampleRate)
	e.midiinput.Start()
	e.patch.Start(sampleRate)

	e.osc.Start(sampleRate)
	e.osc.Waveform = processor.Sin
	e.osc.Freq = 5
}

// Stop suspends the engine
func (e *Engine) Stop() {
	println("do DSP stop")
}

// Process processes a block of samples and midi events
func (e *Engine) Process(samplesIn [][]float32, samplesOut [][]float32, midiIn []midi.Event, midiOut *[]midi.Event) {
	e.midiinput.ProcessMIDI(midiIn)

	var len = len(samplesOut[0])
	for i := 0; i < len; i++ {
		freqs := e.midiinput.Process()
		var sample = e.patch.Process(freqs)

		mic := samplesIn[0][i] * 500
		mod := e.osc.Process()
		if mod < 0 {
			mod = 0
		}
		mic *= mod
		sample += mic

		samplesOut[0][i] = sample
		samplesOut[1][i] = sample
	}
}
