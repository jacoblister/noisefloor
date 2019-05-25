package processor

// Definition exports processor definition
func (e *Envelope) Definition() (name string, inputs []string, outputs []string) {
	return "Envelope", []string{"Gte", "Trg"}, []string{"Out"}
}

//ProcessArray calls process with an array of input/output samples
func (e *Envelope) ProcessArray(in []float32) (output []float32) {
	out := e.Process(in[0], in[1])
	return []float32{out}
}

// Definition exports processor definition
func (g *Gain) Definition() (name string, inputs []string, outputs []string) {
	return "Gain", []string{"In", "Gai"}, []string{"Out"}
}

//ProcessArray calls process with an array of input/output samples
func (g *Gain) ProcessArray(in []float32) (output []float32) {
	out := g.Process(in[0], in[1])
	return []float32{out}
}

// Definition exports processor definition
func (o *Oscillator) Definition() (name string, inputs []string, outputs []string) {
	return "Oscillator", []string{"Frq"}, []string{"Out"}
}

//ProcessArray calls process with an array of input/output samples
func (o *Oscillator) ProcessArray(in []float32) (output []float32) {
	out := o.Process(in[0])
	return []float32{out}
}
