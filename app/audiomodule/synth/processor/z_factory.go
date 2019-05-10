package processor

// Definition exports processor definition
func (e *Envelope) Definition() (name string, inputs []string, outputs []string) {
	return "Envelope", []string{"gate", "trigger"}, []string{"output"}
}

//ProcessArray calls process with an array of input/output samples
func (e *Envelope) ProcessArray(in []float32) (output []float32) {
	out := e.Process(in[0], in[1])
	return []float32{out}
}

// Definition exports processor definition
func (g *Gain) Definition() (name string, inputs []string, outputs []string) {
	return "Gain", []string{"input", "gain"}, []string{"output"}
}

//ProcessArray calls process with an array of input/output samples
func (g *Gain) ProcessArray(in []float32) (output []float32) {
	out := g.Process(in[0], in[1])
	return []float32{out}
}

// Definition exports processor definition
func (o *Oscillator) Definition() (name string, inputs []string, outputs []string) {
	return "Oscillator", []string{}, []string{"output"}
}

//ProcessArray calls process with an array of input/output samples
func (o *Oscillator) ProcessArray(in []float32) (output []float32) {
	out := o.Process()
	return []float32{out}
}
