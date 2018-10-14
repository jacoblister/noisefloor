package processor

import (
	"math"

	. "github.com/jacoblister/noisefloor/common"
)

// Oscillator - basic wave function generator
type Oscillator struct {
	sampleRate    AudioFloat
	currentSample AudioFloat
	freq          AudioFloat
	waveform      []AudioFloat
}

// Start - init oscillaor waveforms
func (o *Oscillator) Start(sampleRate int) {
	o.sampleRate = AudioFloat(sampleRate)
	o.waveform = make([]AudioFloat, sampleRate)
	o.freq = 220

	for i := 0; i < sampleRate; i++ {
		// o.waveform[i] = (float32(i) / (float32(sampleRate) / 2)) - 1
		o.waveform[i] = AudioFloat(math.Sin(float64(2*math.Pi*float64(i)) / float64(sampleRate)))
	}
}

// Stop - deallocate oscialltor
func (o *Oscillator) Stop() {
	o.waveform = nil
}

// Process - produce next sample
func (o *Oscillator) Process() AudioFloat {
	var result = o.waveform[int(o.currentSample)]

	o.currentSample += o.freq
	if o.currentSample >= o.sampleRate {
		o.currentSample -= o.sampleRate
	}

	return result
}
