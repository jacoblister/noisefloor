package processorbasic

import "github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"

// Start - init module
func (r *Add) Start(sampleRate int, connectedMask int) {}

// Stop - release module
func (r *Add) Stop() {}

// Definition exports the constant definition
func (r *Add) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Add", []string{"x","y"}, []string{"Out"},
	[]processor.Parameter{
	}
}

//SetParameter set a single processor parameter
func (r *Add) SetParameter(index int, value float32) {
	switch index {
	} 
}

//ProcessArgs calls process with an array of input/output samples
func (r *Add) ProcessArgs(in []float32) (output []float32) {
	out0 := r.Process(in[0],in[1])
	return []float32{out0}
}

//ProcessSamples calls process with an array of input/output samples
func (r *Add) ProcessSamples(in [][]float32, length int) (out [][]float32) {
	out = make([][]float32, 1)
	out[0] = make([]float32, length)
	for i := 0; i < length; i++ {
		out[0][i] = r.Process(in[0][i], in[1][i])
	}
	return
}

// Start - init module
func (r *Constant) Start(sampleRate int, connectedMask int) {}

// Stop - release module
func (r *Constant) Stop() {}

// Definition exports the constant definition
func (r *Constant) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Constant", []string{}, []string{"Out"},
	[]processor.Parameter{
		processor.Parameter{Name: "Value", Min: 0, Max: 100, Default: 1, Value: float32(r.Value)},
	}
}

//SetParameter set a single processor parameter
func (r *Constant) SetParameter(index int, value float32) {
	switch index {
	case 0:
		r.Value = value
	} 
}

//ProcessArgs calls process with an array of input/output samples
func (r *Constant) ProcessArgs(in []float32) (output []float32) {
	out0 := r.Process()
	return []float32{out0}
}

//ProcessSamples calls process with an array of input/output samples
func (r *Constant) ProcessSamples(in [][]float32, length int) (out [][]float32) {
	out = make([][]float32, 1)
	out[0] = make([]float32, length)
	for i := 0; i < length; i++ {
		out[0][i] = r.Process()
	}
	return
}

// Start - init module
func (r *Divide) Start(sampleRate int, connectedMask int) {}

// Stop - release module
func (r *Divide) Stop() {}

// Definition exports the constant definition
func (r *Divide) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Divide", []string{"x","y"}, []string{"Out"},
	[]processor.Parameter{
	}
}

//SetParameter set a single processor parameter
func (r *Divide) SetParameter(index int, value float32) {
	switch index {
	} 
}

//ProcessArgs calls process with an array of input/output samples
func (r *Divide) ProcessArgs(in []float32) (output []float32) {
	out0 := r.Process(in[0],in[1])
	return []float32{out0}
}

//ProcessSamples calls process with an array of input/output samples
func (r *Divide) ProcessSamples(in [][]float32, length int) (out [][]float32) {
	out = make([][]float32, 1)
	out[0] = make([]float32, length)
	for i := 0; i < length; i++ {
		out[0][i] = r.Process(in[0][i], in[1][i])
	}
	return
}

// Stop - release module
func (r *Envelope) Stop() {}

// Definition exports the constant definition
func (r *Envelope) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Envelope", []string{"Gte","Trg"}, []string{"Out"},
	[]processor.Parameter{
		processor.Parameter{Name: "Attack", Min: 0, Max: 100, Default: 2, Value: float32(r.Attack)},
		processor.Parameter{Name: "Decay", Min: 0, Max: 1000, Default: 100, Value: float32(r.Decay)},
		processor.Parameter{Name: "Sustain", Min: 0, Max: 1, Default: 0.75, Value: float32(r.Sustain)},
		processor.Parameter{Name: "Release", Min: 0, Max: 1000, Default: 1000, Value: float32(r.Release)},
	}
}

//SetParameter set a single processor parameter
func (r *Envelope) SetParameter(index int, value float32) {
	switch index {
	case 0:
		r.Attack = value
	case 1:
		r.Decay = value
	case 2:
		r.Sustain = value
	case 3:
		r.Release = value
	} 
}

//ProcessArgs calls process with an array of input/output samples
func (r *Envelope) ProcessArgs(in []float32) (output []float32) {
	out0 := r.Process(in[0],in[1])
	return []float32{out0}
}

//ProcessSamples calls process with an array of input/output samples
func (r *Envelope) ProcessSamples(in [][]float32, length int) (out [][]float32) {
	out = make([][]float32, 1)
	out[0] = make([]float32, length)
	for i := 0; i < length; i++ {
		out[0][i] = r.Process(in[0][i], in[1][i])
	}
	return
}

// Start - init module
func (r *Gain) Start(sampleRate int, connectedMask int) {}

// Stop - release module
func (r *Gain) Stop() {}

// Definition exports the constant definition
func (r *Gain) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Gain", []string{"In","Gai"}, []string{"Out"},
	[]processor.Parameter{
		processor.Parameter{Name: "Level", Min: 0, Max: 2, Default: 1, Value: float32(r.Level)},
	}
}

//SetParameter set a single processor parameter
func (r *Gain) SetParameter(index int, value float32) {
	switch index {
	case 0:
		r.Level = value
	} 
}

//ProcessArgs calls process with an array of input/output samples
func (r *Gain) ProcessArgs(in []float32) (output []float32) {
	out0 := r.Process(in[0],in[1])
	return []float32{out0}
}

//ProcessSamples calls process with an array of input/output samples
func (r *Gain) ProcessSamples(in [][]float32, length int) (out [][]float32) {
	out = make([][]float32, 1)
	out[0] = make([]float32, length)
	for i := 0; i < length; i++ {
		out[0][i] = r.Process(in[0][i], in[1][i])
	}
	return
}

// Start - init module
func (r *Multiply) Start(sampleRate int, connectedMask int) {}

// Stop - release module
func (r *Multiply) Stop() {}

// Definition exports the constant definition
func (r *Multiply) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Multiply", []string{"x","y"}, []string{"Out"},
	[]processor.Parameter{
	}
}

//SetParameter set a single processor parameter
func (r *Multiply) SetParameter(index int, value float32) {
	switch index {
	} 
}

//ProcessArgs calls process with an array of input/output samples
func (r *Multiply) ProcessArgs(in []float32) (output []float32) {
	out0 := r.Process(in[0],in[1])
	return []float32{out0}
}

//ProcessSamples calls process with an array of input/output samples
func (r *Multiply) ProcessSamples(in [][]float32, length int) (out [][]float32) {
	out = make([][]float32, 1)
	out[0] = make([]float32, length)
	for i := 0; i < length; i++ {
		out[0][i] = r.Process(in[0][i], in[1][i])
	}
	return
}

// Stop - release module
func (r *Oscillator) Stop() {}

// Definition exports the constant definition
func (r *Oscillator) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Oscillator", []string{"Frq"}, []string{"Out"},
	[]processor.Parameter{
		processor.Parameter{Name: "Waveform", Min: 0, Max: 3, Default: 0, Value: float32(r.Waveform)},
	}
}

//SetParameter set a single processor parameter
func (r *Oscillator) SetParameter(index int, value float32) {
	switch index {
	case 0:
		r.Waveform = Waveform(value + 0.5)
	} 
}

//ProcessArgs calls process with an array of input/output samples
func (r *Oscillator) ProcessArgs(in []float32) (output []float32) {
	out0 := r.Process(in[0])
	return []float32{out0}
}

//ProcessSamples calls process with an array of input/output samples
func (r *Oscillator) ProcessSamples(in [][]float32, length int) (out [][]float32) {
	out = make([][]float32, 1)
	out[0] = make([]float32, length)
	for i := 0; i < length; i++ {
		out[0][i] = r.Process(in[0][i])
	}
	return
}

// Stop - release module
func (r *OscSync) Stop() {}

// Definition exports the constant definition
func (r *OscSync) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "OscSync", []string{"Frq","Syn","Pse"}, []string{"Out"},
	[]processor.Parameter{
		processor.Parameter{Name: "Waveform", Min: 0, Max: 3, Default: 0, Value: float32(r.Waveform)},
	}
}

//SetParameter set a single processor parameter
func (r *OscSync) SetParameter(index int, value float32) {
	switch index {
	case 0:
		r.Waveform = Waveform(value + 0.5)
	} 
}

//ProcessArgs calls process with an array of input/output samples
func (r *OscSync) ProcessArgs(in []float32) (output []float32) {
	out0 := r.Process(in[0],in[1],in[2])
	return []float32{out0}
}

//ProcessSamples calls process with an array of input/output samples
func (r *OscSync) ProcessSamples(in [][]float32, length int) (out [][]float32) {
	out = make([][]float32, 1)
	out[0] = make([]float32, length)
	for i := 0; i < length; i++ {
		out[0][i] = r.Process(in[0][i], in[1][i], in[2][i])
	}
	return
}

// Stop - release module
func (r *Scope) Stop() {}

// Definition exports the constant definition
func (r *Scope) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Scope", []string{"InA","InB"}, []string{"OutA","OutB"},
	[]processor.Parameter{
		processor.Parameter{Name: "Trigger", Min: 0, Max: 1, Default: 1, Value: float32(r.Trigger)},
		processor.Parameter{Name: "Skip", Min: 0, Max: 200, Default: 4, Value: float32(r.Skip)},
	}
}

//SetParameter set a single processor parameter
func (r *Scope) SetParameter(index int, value float32) {
	switch index {
	case 0:
		r.Trigger = int(value + 0.5)
	case 1:
		r.Skip = int(value + 0.5)
	} 
}

//ProcessArgs calls process with an array of input/output samples
func (r *Scope) ProcessArgs(in []float32) (output []float32) {
	out0,out1 := r.Process(in[0],in[1])
	return []float32{out0,out1}
}

//ProcessSamples calls process with an array of input/output samples
func (r *Scope) ProcessSamples(in [][]float32, length int) (out [][]float32) {
	out = make([][]float32, 2)
	out[0] = make([]float32, length)
	out[1] = make([]float32, length)
	for i := 0; i < length; i++ {
		out[0][i], out[1][i] = r.Process(in[0][i], in[1][i])
	}
	return
}

// Start - init module
func (r *Select) Start(sampleRate int, connectedMask int) {}

// Stop - release module
func (r *Select) Stop() {}

// Definition exports the constant definition
func (r *Select) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Select", []string{"a","b"}, []string{"Out"},
	[]processor.Parameter{
		processor.Parameter{Name: "Input", Min: 0, Max: 1, Default: 0, Value: float32(r.Input)},
	}
}

//SetParameter set a single processor parameter
func (r *Select) SetParameter(index int, value float32) {
	switch index {
	case 0:
		r.Input = int(value + 0.5)
	} 
}

//ProcessArgs calls process with an array of input/output samples
func (r *Select) ProcessArgs(in []float32) (output []float32) {
	out0 := r.Process(in[0],in[1])
	return []float32{out0}
}

//ProcessSamples calls process with an array of input/output samples
func (r *Select) ProcessSamples(in [][]float32, length int) (out [][]float32) {
	out = make([][]float32, 1)
	out[0] = make([]float32, length)
	for i := 0; i < length; i++ {
		out[0][i] = r.Process(in[0][i], in[1][i])
	}
	return
}

// Start - init module
func (r *Splitter) Start(sampleRate int, connectedMask int) {}

// Stop - release module
func (r *Splitter) Stop() {}

// Definition exports the constant definition
func (r *Splitter) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Splitter", []string{"In"}, []string{"Out0","Out1","Out2","Out3"},
	[]processor.Parameter{
	}
}

//SetParameter set a single processor parameter
func (r *Splitter) SetParameter(index int, value float32) {
	switch index {
	} 
}

//ProcessArgs calls process with an array of input/output samples
func (r *Splitter) ProcessArgs(in []float32) (output []float32) {
	out0,out1,out2,out3 := r.Process(in[0])
	return []float32{out0,out1,out2,out3}
}

//ProcessSamples calls process with an array of input/output samples
func (r *Splitter) ProcessSamples(in [][]float32, length int) (out [][]float32) {
	out = make([][]float32, 4)
	out[0] = make([]float32, length)
	out[1] = make([]float32, length)
	out[2] = make([]float32, length)
	out[3] = make([]float32, length)
	for i := 0; i < length; i++ {
		out[0][i], out[1][i], out[2][i], out[3][i] = r.Process(in[0][i])
	}
	return
}

// Start - init module
func (r *Sum) Start(sampleRate int, connectedMask int) {}

// Stop - release module
func (r *Sum) Stop() {}

// Definition exports the constant definition
func (r *Sum) Definition() (name string, inputs []string, outputs []string, parameters []processor.Parameter) {
	return "Sum", []string{"In0","In1","In2","In3"}, []string{"Out"},
	[]processor.Parameter{
	}
}

//SetParameter set a single processor parameter
func (r *Sum) SetParameter(index int, value float32) {
	switch index {
	} 
}

//ProcessArgs calls process with an array of input/output samples
func (r *Sum) ProcessArgs(in []float32) (output []float32) {
	out0 := r.Process(in[0],in[1],in[2],in[3])
	return []float32{out0}
}

//ProcessSamples calls process with an array of input/output samples
func (r *Sum) ProcessSamples(in [][]float32, length int) (out [][]float32) {
	out = make([][]float32, 1)
	out[0] = make([]float32, length)
	for i := 0; i < length; i++ {
		out[0][i] = r.Process(in[0][i], in[1][i], in[2][i], in[3][i])
	}
	return
}

