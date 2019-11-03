package processor

//Parameter is a single paramter with a range and default
type Parameter struct {
	Name    string
	Default float32
	Min     float32
	Max     float32
	Value   float32
}

// Definition exports processor definition
func (e *Envelope) Definition() (name string, inputs []string, outputs []string, parameters []Parameter) {
	return "Envelope", []string{"Gte", "Trg"}, []string{"Out"},
		[]Parameter{
			Parameter{Name: "Attack", Min: 0, Max: 100, Value: e.Attack},
			Parameter{Name: "Decay", Min: 0, Max: 1000, Value: e.Decay},
			Parameter{Name: "Sustain", Min: 0, Max: 1.0, Value: e.Sustain},
			Parameter{Name: "Release", Min: 0, Max: 1000, Value: e.Release},
		}
}

//ProcessArray calls process with an array of input/output samples
func (e *Envelope) ProcessArray(in []float32) (output []float32) {
	out := e.Process(in[0], in[1])
	return []float32{out}
}

//SetParameter set a single processor parameter
func (e *Envelope) SetParameter(index int, value float32) {
	switch index {
	case 0:
		e.Attack = value
	case 1:
		e.Decay = value
	case 2:
		e.Sustain = value
	case 3:
		e.Release = value
	}
}

// Definition exports processor definition
func (g *Gain) Definition() (name string, inputs []string, outputs []string, parameters []Parameter) {
	return "Gain", []string{"In", "Gai"}, []string{"Out"}, []Parameter{}
}

//ProcessArray calls process with an array of input/output samples
func (g *Gain) ProcessArray(in []float32) (output []float32) {
	out := g.Process(in[0], in[1])
	return []float32{out}
}

//SetParameter set a single processor parameter
func (g *Gain) SetParameter(index int, value float32) {}

// Definition exports processor definition
func (o *Oscillator) Definition() (name string, inputs []string, outputs []string, parameters []Parameter) {
	return "Oscillator", []string{"Frq"}, []string{"Out"},
		[]Parameter{Parameter{Name: "Wave", Min: 0, Max: 3, Value: float32(o.Waveform)}}
}

//ProcessArray calls process with an array of input/output samples
func (o *Oscillator) ProcessArray(in []float32) (output []float32) {
	out := o.Process(in[0])
	return []float32{out}
}

//SetParameter set a single processor parameter
func (o *Oscillator) SetParameter(index int, value float32) {
	value = value * 4 / 3
	switch index {
	case 0:
		o.Waveform = Waveform(value)
	}
}

// Definition exports processor definition
func (s *Splitter) Definition() (name string, inputs []string, outputs []string, parameters []Parameter) {
	return "Splitter", []string{"In"}, []string{"Out1", "Out2", "Out3", "Out4"}, []Parameter{}
}

//ProcessArray calls process with an array of input/output samples
func (s *Splitter) ProcessArray(in []float32) (output []float32) {
	out1, out2, out3, out4 := s.Process(in[0])
	return []float32{out1, out2, out3, out4}
}

//SetParameter set a single processor parameter
func (s *Splitter) SetParameter(index int, value float32) {}

// Definition exports processor definition
func (s *Scope) Definition() (name string, inputs []string, outputs []string, parameters []Parameter) {
	return "Scope", []string{"In"}, []string{}, []Parameter{}
}

//ProcessArray calls process with an array of input/output samples
func (s *Scope) ProcessArray(in []float32) (output []float32) {
	s.Process(in[0])
	return []float32{}
}

//SetParameter set a single processor parameter
func (s *Scope) SetParameter(index int, value float32) {}
