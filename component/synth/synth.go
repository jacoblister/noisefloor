package synth

import (
	"github.com/jacoblister/noisefloor/component/synth/processor"
	"github.com/jacoblister/noisefloor/midi"
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
	e.osc.Freq = 440
}

// Stop suspends the engine
func (e *Engine) Stop() {
	println("do DSP stop")
}

// Process processes a block of samples and midi events
func (e *Engine) Process(samplesIn [][]float32, midiIn []midi.Event) (samplesOut [][]float32, midiOut []midi.Event) {
	e.midiinput.ProcessMIDI(midiIn)

	var len = len(samplesIn[0])
	for i := 0; i < len; i++ {
		// var sample = e.osc.Process()
		freqs := e.midiinput.Process()
		var sample = e.patch.Process(freqs)
		sample += samplesIn[0][i]

		// mic := samplesIn[0][i] * 500
		// mod := e.osc.Process()
		// if mod < 0 {
		// 	mod = 0
		// }
		// mic *= mod
		// sample += mic
		// sample = e.osc.Process()

		samplesIn[0][i] = sample
		samplesIn[1][i] = sample
	}

	return samplesIn, midiIn
}
