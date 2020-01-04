package processorbasic

// Multiply - multiply two values
type Multiply struct {
	Dummy int
}

// Process - produce next sample
func (m *Multiply) Process(x float32, y float32) (Out float32) {
	Out = x * y
	return
}
