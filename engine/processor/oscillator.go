package processor

import (
	"math"

	. "github.com/jacoblister/noisefloor/common"
)

// Waveform enumerated type
type Waveform int

// Basic wave shapes
const (
	Sin      Waveform = 0
	Saw      Waveform = 1
	Square   Waveform = 2
	Triangle Waveform = 3
)

// Oscillator - basic wave function generator
type Oscillator struct {
	Waveform Waveform   `value:"Sin"`
	Freq     AudioFloat `value:"440" min:"20" max:"20000"`

	sampleRate    AudioFloat
	currentSample AudioFloat
	waveTable     [4][]AudioFloat
}

// Start - init oscillaor waveforms
func (o *Oscillator) Start(sampleRate int) {
	o.sampleRate = AudioFloat(sampleRate)
	o.waveTable[Sin] = make([]AudioFloat, sampleRate)
	o.waveTable[Saw] = make([]AudioFloat, sampleRate)
	o.Freq = 220

	for i := 0; i < sampleRate; i++ {
		o.waveTable[Sin][i] = AudioFloat(math.Sin(float64(2*math.Pi*float64(i)) / float64(sampleRate)))
		o.waveTable[Saw][i] = (AudioFloat(i) / (AudioFloat(sampleRate) / 2)) - 1
	}
}

// Stop - deallocate oscialltor
func (o *Oscillator) Stop() {
	o.waveTable[Sin] = nil
}

// Process - produce next sample
func (o *Oscillator) Process() AudioFloat {
	var result = o.waveTable[o.Waveform][int(o.currentSample)]

	o.currentSample += o.Freq
	if o.currentSample >= o.sampleRate {
		o.currentSample -= o.sampleRate
	}

	return result
}
