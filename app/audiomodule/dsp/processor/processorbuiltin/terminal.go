package processorbuiltin

import "strconv"

//Terminal is the audio input/output connector to hardware or parent graph
//it is a special case which is not part of the compiled graph
type Terminal struct {
	isInput    bool
	connectors int

	samples     [][]float32
	sampleIndex int
}

// Start - init envelope generator
func (t *Terminal) Start(sampleRate int) {}

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
func (t *Terminal) Definition() (name string, inputs []string, outputs []string) {
	inputs = []string{}
	outputs = []string{}

	for i := 0; i < t.connectors; i++ {
		if t.isInput {
			inputs = append(inputs, "In"+strconv.Itoa(i))
		} else {
			outputs = append(inputs, "Out"+strconv.Itoa(i))
		}
	}

	return "Terminal", inputs, outputs
}

//ProcessArray calls process with an array of input/output samples
func (t *Terminal) ProcessArray(in []float32) (output []float32) {
	//TODO handle input terminal

	for i := 0; i < len(in); i++ {
		t.samples[i][t.sampleIndex] = in[i]
	}
	t.sampleIndex++
	return in
}
