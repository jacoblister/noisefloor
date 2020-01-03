package processor

//Parameter is a single paramter with a range and default
type Parameter struct {
	Name    string
	Default float32
	Min     float32
	Max     float32
	Value   float32
}

func boolTofloat32(value bool) float32 {
	if value {
		return 1
	}
	return 0
}

func float32toBool(value float32) bool {
	if value <= 0 {
		return false
	}
	return true
}

// Definition exports the constant definition
func (c *Constant) Definition() (name string, inputs []string, outputs []string, parameters []Parameter) {
	return "Constant", []string{}, []string{"Out"},
		[]Parameter{
			Parameter{Name: "Value", Min: 0, Max: 10, Default: 1, Value: c.Value},
		}
}

//ProcessArgs calls process with an array of input/output samples
func (c *Constant) ProcessArgs(in []float32) (output []float32) {
	out := c.Process()
	return []float32{out}
}

//ProcessSamples calls process with an array of input/output samples
func (c *Constant) ProcessSamples(in [][]float32, length int) (output [][]float32) {
	output = make([][]float32, 1)
	output[0] = make([]float32, length)
	for i := 0; i < length; i++ {
		output[0][i] = c.Process()
	}
	return
}

//SetParameter set a single processor parameter
func (c *Constant) SetParameter(index int, value float32) {
	switch index {
	case 0:
		c.Value = value
	}
}

// Definition exports the constant definition
func (d *Divide) Definition() (name string, inputs []string, outputs []string, parameters []Parameter) {
	return "Divide", []string{"x", "y"}, []string{"Out"},
		[]Parameter{}
}

//ProcessArgs calls process with an array of input/output samples
func (d *Divide) ProcessArgs(in []float32) (output []float32) {
	out := d.Process(in[0], in[1])
	return []float32{out}
}

//ProcessSamples calls process with an array of input/output samples
func (d *Divide) ProcessSamples(in [][]float32, length int) (output [][]float32) {
	output = make([][]float32, 1)
	output[0] = make([]float32, length)
	for i := 0; i < length; i++ {
		output[0][i] = d.Process(in[0][i], in[1][i])
	}
	return
}

//SetParameter set a single processor parameter
func (d *Divide) SetParameter(index int, value float32) {}

// Definition exports processor definition
func (e *Envelope) Definition() (name string, inputs []string, outputs []string, parameters []Parameter) {
	return "Envelope", []string{"Gte", "Trg"}, []string{"Out"},
		[]Parameter{
			Parameter{Name: "Attack", Min: 0, Max: 100, Default: 2, Value: e.Attack},
			Parameter{Name: "Decay", Min: 0, Max: 1000, Default: 100, Value: e.Decay},
			Parameter{Name: "Sustain", Min: 0, Max: 1.0, Default: 0.75, Value: e.Sustain},
			Parameter{Name: "Release", Min: 0, Max: 1000, Default: 1000, Value: e.Release},
		}
}

//ProcessArgs calls process with an array of input/output samples
func (e *Envelope) ProcessArgs(in []float32) (output []float32) {
	out := e.Process(in[0], in[1])
	return []float32{out}
}

//ProcessSamples calls process with an array of input/output samples
func (e *Envelope) ProcessSamples(in [][]float32, length int) (output [][]float32) {
	output = make([][]float32, 1)
	output[0] = make([]float32, length)
	for i := 0; i < length; i++ {
		output[0][i] = e.Process(in[0][i], in[1][i])
	}
	return
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
	return "Gain", []string{"In", "Gai"}, []string{"Out"}, []Parameter{
		Parameter{Name: "Level", Min: 0, Max: 2, Default: 1, Value: g.Level},
	}
}

//ProcessArgs calls process with with args as an array
func (g *Gain) ProcessArgs(in []float32) (output []float32) {
	out := g.Process(in[0], in[1])
	return []float32{out}
}

//ProcessSamples calls process with an array of input/output samples
func (g *Gain) ProcessSamples(in [][]float32, length int) (output [][]float32) {
	output = make([][]float32, 1)
	output[0] = make([]float32, length)
	for i := 0; i < length; i++ {
		output[0][i] = g.Process(in[0][i], in[1][i])
	}
	return
}

//SetParameter set a single processor parameter
func (g *Gain) SetParameter(index int, value float32) {
	switch index {
	case 0:
		g.Level = value
	}
}

// Definition exports the constant definition
func (m *Multiply) Definition() (name string, inputs []string, outputs []string, parameters []Parameter) {
	return "Multiply", []string{"x", "y"}, []string{"Out"},
		[]Parameter{}
}

//ProcessArgs calls process with an array of input/output samples
func (m *Multiply) ProcessArgs(in []float32) (output []float32) {
	out := m.Process(in[0], in[1])
	return []float32{out}
}

//ProcessSamples calls process with an array of input/output samples
func (m *Multiply) ProcessSamples(in [][]float32, length int) (output [][]float32) {
	output = make([][]float32, 1)
	output[0] = make([]float32, length)
	for i := 0; i < length; i++ {
		output[0][i] = m.Process(in[0][i], in[1][i])
	}
	return
}

//SetParameter set a single processor parameter
func (m *Multiply) SetParameter(index int, value float32) {}

// Definition exports processor definition
func (o *Oscillator) Definition() (name string, inputs []string, outputs []string, parameters []Parameter) {
	return "Oscillator", []string{"Frq"}, []string{"Out"},
		[]Parameter{Parameter{Name: "Wave", Min: 0, Max: 3, Value: float32(o.Waveform)}}
}

//ProcessArgs calls process with args as an array
func (o *Oscillator) ProcessArgs(in []float32) (output []float32) {
	out := o.Process(in[0])
	return []float32{out}
}

//ProcessSamples calls process with an array of input/output samples
func (o *Oscillator) ProcessSamples(in [][]float32, length int) (output [][]float32) {
	output = make([][]float32, 1)
	output[0] = make([]float32, length)
	for i := 0; i < length; i++ {
		output[0][i] = o.Process(in[0][i])
	}
	return
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

//ProcessArgs calls process with args as an array
func (s *Splitter) ProcessArgs(in []float32) (output []float32) {
	out1, out2, out3, out4 := s.Process(in[0])
	return []float32{out1, out2, out3, out4}
}

//ProcessSamples calls process with an array of input/output samples
func (s *Splitter) ProcessSamples(in [][]float32, length int) (output [][]float32) {
	output = make([][]float32, 4)
	output[0] = make([]float32, length)
	output[1] = make([]float32, length)
	output[2] = make([]float32, length)
	output[3] = make([]float32, length)
	for i := 0; i < length; i++ {
		output[0][i], output[1][i], output[2][i], output[3][i] = s.Process(in[0][i])
	}
	return
}

//SetParameter set a single processor parameter
func (s *Splitter) SetParameter(index int, value float32) {}

// Definition exports processor definition
func (s *Sum) Definition() (name string, inputs []string, outputs []string, parameters []Parameter) {
	return "Sum", []string{"In1", "In2", "In3", "In4"}, []string{"Out"}, []Parameter{}
}

//ProcessArgs calls process with args as an array
func (s *Sum) ProcessArgs(in []float32) (output []float32) {
	out := s.Process(in[0], in[1], in[2], in[3])
	return []float32{out}
}

//ProcessSamples calls process with an array of input/output samples
func (s *Sum) ProcessSamples(in [][]float32, length int) (output [][]float32) {
	output = make([][]float32, 1)
	output[0] = make([]float32, length)
	for i := 0; i < length; i++ {
		output[0][i] = s.Process(in[0][i], in[1][i], in[2][i], in[3][i])
	}
	return
}

//SetParameter set a single processor parameter
func (s *Sum) SetParameter(index int, value float32) {}

// Definition exports processor definition
func (s *Scope) Definition() (name string, inputs []string, outputs []string, parameters []Parameter) {
	return "Scope", []string{"In"}, []string{}, []Parameter{
		Parameter{Name: "Trigger", Min: 0, Max: 1, Default: 1, Value: boolTofloat32(s.Trigger)},
		Parameter{Name: "Skip", Min: 0, Max: 200, Default: 4, Value: float32(s.Skip)},
	}
}

//ProcessArgs calls process with args as an array
func (s *Scope) ProcessArgs(in []float32) (output []float32) {
	s.Process(in[0])
	return []float32{}
}

//ProcessSamples calls process with an array of input/output samples
func (s *Scope) ProcessSamples(in [][]float32, length int) (output [][]float32) {
	output = make([][]float32, 0)
	for i := 0; i < length; i++ {
		s.Process(in[0][i])
	}
	return
}

//SetParameter set a single processor parameter
func (s *Scope) SetParameter(index int, value float32) {
	switch index {
	case 0:
		value = value - 0.5
		s.Trigger = float32toBool(value)
	case 1:
		s.Skip = int(value)
	}
}
