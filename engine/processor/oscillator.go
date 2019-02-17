package processor

import (
	"math"
)

// Waveform enumerated type
type Waveform int

// Basic wave shapes
const (
	Sin Waveform = iota
	Saw
	Square
	Triangle
	maxWaveform
)

const maxSamples = 48000

// Oscillator - basic wave function generator
type Oscillator struct {
	Waveform Waveform   `default:"Sin"`
	Freq     float32 `default:"440" min:"20" max:"20000"`

	sampleRate    float32
	currentSample float32
	waveTable     [maxWaveform][maxSamples]float32
}

// Start - init oscillaor waveforms
func (o *Oscillator) Start(sampleRate int) {
	if sampleRate > maxSamples {
		panic("sample rate is out of range")
	}

	o.sampleRate = float32(sampleRate)
	o.Freq = 220

	for i := 0; i < sampleRate; i++ {
		o.waveTable[Sin][i] = float32(math.Sin(float64(2*math.Pi*float64(i)) / float64(sampleRate)))
		o.waveTable[Saw][i] = (float32(i) / (float32(sampleRate) / 2)) - 1
		o.waveTable[Square][i] = -1
		if i < sampleRate/2 {
			o.waveTable[Square][i] = 1
		}
	}
}

// Stop - deallocate oscialltor
func (o *Oscillator) Stop() {}

// Process - produce next sample
func (o *Oscillator) Process() (output float32) {
	output = o.waveTable[o.Waveform][int(o.currentSample)]

	o.currentSample += o.Freq
	if o.currentSample >= o.sampleRate {
		o.currentSample -= o.sampleRate
	}

	return
}
