package processorbuiltin

import "strconv"

//Terminal is the audio input/output connector to hardware or parent graph
//it is a special case which is not part of the compiled graph
type Terminal struct {
	isInput    bool
	connectors int
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

// Definition exports the terminal connectors, given the input/output type and count
func (t *Terminal) Definition() (name string, inputs []string, outputs []string) {
	inputs = []string{}
	outputs = []string{}

	for i := 0; i < t.connectors; i++ {
		if t.isInput {
			inputs = append(inputs, "in"+strconv.Itoa(i))
		} else {
			outputs = append(inputs, "out"+strconv.Itoa(i))
		}
	}

	return "Terminal", inputs, outputs
}

//ProcessArray calls process with an array of input/output samples
//is is a no-op for Terminal, which is not compiled in the graph
func (t *Terminal) ProcessArray(in []float32) (output []float32) {
	return
}
