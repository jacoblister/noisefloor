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

// Definition exports processor definition
func (s *Splitter) Definition() (name string, inputs []string, outputs []string) {
	return "Splitter", []string{"In"}, []string{"Out1", "Out2", "Out3", "Out4"}
}

//ProcessArray calls process with an array of input/output samples
func (s *Splitter) ProcessArray(in []float32) (output []float32) {
	out1, out2, out3, out4 := s.Process(in[0])
	return []float32{out1, out2, out3, out4}
}

// Definition exports processor definition
func (s *Scope) Definition() (name string, inputs []string, outputs []string) {
	return "Scope", []string{"In"}, []string{}
}

//ProcessArray calls process with an array of input/output samples
func (s *Scope) ProcessArray(in []float32) (output []float32) {
	s.Process(in[0])
	return []float32{}
}
