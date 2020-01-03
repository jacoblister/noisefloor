package processorbasic

// Sum - add inputs
type Sum struct {
	Dummy int
}

// Start - init splitter
func (s *Sum) Start(sampleRate int) {}

// Process - produce next sample
func (s *Sum) Process(in1 float32, in2 float32, in3 float32, in4 float32) (out float32) {
	return in1 + in2 + in3 + in4
}
