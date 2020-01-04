package processorbasic

// Start - init module
func (r *Constant) Start(sampleRate int) {}

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

//ProcessArgs calls process with an array of input/output samples
func (r *Scope) ProcessArgs(in []float32) (output []float32) {
	out0 := r.Process(in[0])
	return []float32{out0}
}

//ProcessSamples calls process with an array of input/output samples
func (r *Scope) ProcessSamples(in [][]float32, length int) (out [][]float32) {
	out = make([][]float32, 1)
	out[0] = make([]float32, length)
	for i := 0; i < length; i++ {
		out[0][i] = r.Process(in[0][i])
	}
	return
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

