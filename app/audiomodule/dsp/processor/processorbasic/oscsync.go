package processorbasic

import (
	"math"
)

// OscSync - basic Oscillator with sync, phase offset,
type OscSync struct {
	Waveform Waveform `default:"0" min:"0" max:"3"`

	sampleRate    float32
	currentSample float32
	waveTable     [maxWaveform][maxSamples]float32
}

// Start - init oscillaor waveforms
func (o *OscSync) Start(sampleRate int, connectedMask int) {
	if sampleRate > maxSamples {
		panic("sample rate is out of range")
	}

	o.sampleRate = float32(sampleRate)

	for i := 0; i < sampleRate; i++ {
		o.waveTable[Sin][i] = float32(math.Sin(float64(2*math.Pi*float64(i)) / float64(sampleRate)))
		o.waveTable[Saw][i] = (float32(i) / (float32(sampleRate) / 2)) - 1
		o.waveTable[Square][i] = -1
		if i < sampleRate/2 {
			o.waveTable[Square][i] = 1
		}

		if i < sampleRate/2 {
			o.waveTable[Triangle][i] = (float32(i) / (float32(sampleRate) / 4)) - 1
		}
	}
}

// Process - produce next sample
func (o *OscSync) Process(Frq float32, Syn float32, Pse float32) (Out float32) {
	if Pse > 1 {
		Pse = 1
	} else if Pse < -1 {
		Pse = -1
	}

	index := o.currentSample + Pse*o.sampleRate
	if index >= o.sampleRate {
		index -= o.sampleRate
	} else if index < 0 {
		index += o.sampleRate
	}

	Out = o.waveTable[o.Waveform][int(index)]

	o.currentSample += Frq
	if o.currentSample >= o.sampleRate {
		o.currentSample -= o.sampleRate
		// fmt.Println(o.currentSample, Pse*Frq)
	}

	return
}
