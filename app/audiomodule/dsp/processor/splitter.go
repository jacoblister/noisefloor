package processor

// Splitter - split signal to multiple outputs
type Splitter struct {
	Dummy int
}

// Start - init splitter
func (s *Splitter) Start(sampleRate int) {}

// Process - produce next sample
func (s *Splitter) Process(input float32) (out1 float32, out2 float32, out3 float32, out4 float32) {
	return input, input, input, input
}
