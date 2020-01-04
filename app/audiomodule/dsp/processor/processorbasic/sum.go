package processorbasic

// Sum - add inputs
type Sum struct {
	Dummy int
}

// Process - produce next sample
func (s *Sum) Process(In1 float32, In2 float32, In3 float32, In4 float32) (Out float32) {
	return In1 + In2 + In3 + In4
}
