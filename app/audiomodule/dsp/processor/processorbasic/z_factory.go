package processorbasic

import "github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"

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
func (c *Constant) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Constant", []string{}, []string{"Out"},
		[]processor.Parameter{
			processor.Parameter{Name: "Value", Min: 0, Max: 10, Default: 1, Value: c.Value},
		}
}

//SetParameter set a single processor parameter
func (c *Constant) SetParameter(index int, value float32) {
	switch index {
	case 0:
		c.Value = value
	}
}

// Definition exports the constant definition
func (d *Divide) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Divide", []string{"x", "y"}, []string{"Out"},
		[]processor.Parameter{}
}

//SetParameter set a single processor parameter
func (d *Divide) SetParameter(index int, value float32) {}

// Definition exports processor definition
func (e *Envelope) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Envelope", []string{"Gte", "Trg"}, []string{"Out"},
		[]processor.Parameter{
			processor.Parameter{Name: "Attack", Min: 0, Max: 100, Default: 2, Value: e.Attack},
			processor.Parameter{Name: "Decay", Min: 0, Max: 1000, Default: 100, Value: e.Decay},
			processor.Parameter{Name: "Sustain", Min: 0, Max: 1.0, Default: 0.75, Value: e.Sustain},
			processor.Parameter{Name: "Release", Min: 0, Max: 1000, Default: 1000, Value: e.Release},
		}
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
func (g *Gain) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Gain", []string{"In", "Gai"}, []string{"Out"}, []processor.Parameter{
		processor.Parameter{Name: "Level", Min: 0, Max: 2, Default: 1, Value: g.Level},
	}
}

//SetParameter set a single processor parameter
func (g *Gain) SetParameter(index int, value float32) {
	switch index {
	case 0:
		g.Level = value
	}
}

// Definition exports the constant definition
func (m *Multiply) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Multiply", []string{"x", "y"}, []string{"Out"},
		[]processor.Parameter{}
}

//SetParameter set a single processor parameter
func (m *Multiply) SetParameter(index int, value float32) {}

// Definition exports processor definition
func (o *Oscillator) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Oscillator", []string{"Frq"}, []string{"Out"},
		[]processor.Parameter{processor.Parameter{Name: "Wave", Min: 0, Max: 3, Value: float32(o.Waveform)}}
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
func (s *Splitter) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Splitter", []string{"In"}, []string{"Out1", "Out2", "Out3", "Out4"}, []processor.Parameter{}
}

//SetParameter set a single processor parameter
func (s *Splitter) SetParameter(index int, value float32) {}

// Definition exports processor definition
func (s *Sum) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Sum", []string{"In1", "In2", "In3", "In4"}, []string{"Out"}, []processor.Parameter{}
}

//SetParameter set a single processor parameter
func (s *Sum) SetParameter(index int, value float32) {}

// Definition exports processor definition
func (s *Scope) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Scope", []string{"In"}, []string{"Out"}, []processor.Parameter{
		processor.Parameter{Name: "Trigger", Min: 0, Max: 1, Default: 1, Value: float32(s.Trigger)},
		processor.Parameter{Name: "Skip", Min: 0, Max: 200, Default: 4, Value: float32(s.Skip)},
	}
}

//SetParameter set a single processor parameter
func (s *Scope) SetParameter(index int, value float32) {
	switch index {
	case 0:
		value = value - 0.5
		s.Trigger = int(value)
	case 1:
		s.Skip = int(value)
	}
}
