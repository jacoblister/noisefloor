package processorbuiltin

import (
	"strconv"

	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"
)

//Terminal is the audio input/output connector to hardware or parent graph
//it is a special case which is not part of the compiled graph
type Terminal struct {
	isInput    bool
	connectors int

	samples     [][]float32
	sampleIndex int
}

// Start - init processor
func (t *Terminal) Start(sampleRate int, connectedMask int) {}

// Process - produce next sample
func (t *Terminal) Process() {
	return
}

// SetParameters sets the internal params
func (t *Terminal) SetParameters(isInput bool, connectors int) {
	t.isInput = isInput
	t.connectors = connectors
}

//SetSamples resets the input/output sample buffer
func (t *Terminal) SetSamples(samples [][]float32) {
	t.samples = samples
	t.sampleIndex = 0
}

// Definition exports the terminal connectors, given the input/output type and count
func (t *Terminal) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	if t.connectors == 0 {
		t.isInput = true
		t.connectors = 2
	}

	inputs = []string{}
	outputs = []string{}

	for i := 0; i < t.connectors; i++ {
		if t.isInput {
			inputs = append(inputs, "In"+strconv.Itoa(i))
		} else {
			outputs = append(inputs, "Out"+strconv.Itoa(i))
		}
	}

	return "Terminal", inputs, outputs, []processor.Parameter{}
}

//ProcessArgs calls process with an array of input/output samples
func (t *Terminal) ProcessArgs(in []float32) (output []float32) {
	//TODO handle input terminal

	for i := 0; i < len(in); i++ {
		t.samples[i][t.sampleIndex] = in[i]
	}
	t.sampleIndex++
	return in
}

//ProcessSamples calls process with an array of input/output samples
func (t *Terminal) ProcessSamples(in [][]float32, out [][]float32, length int) {
	for i := 0; i < len(in); i++ {
		for j := 0; j < length; j++ {
			t.samples[i][j] = in[i][j]
		}
	}
}

//SetParameter set a single processor parameter
func (t *Terminal) SetParameter(index int, value float32) {}
