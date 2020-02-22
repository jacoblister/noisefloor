package processorbasic

// Splitter - split signal to multiple outputs
type Splitter struct {
	Dummy int
}

// Process - produce next sample
func (s *Splitter) Process(In float32) (Out0 float32, Out1 float32, Out2 float32, Out3 float32) {
	return In, In, In, In
}
