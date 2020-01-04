package processorbasic

// Add - add two values
type Add struct {
	Dummy int
}

// Process - produce next sample
func (a *Add) Process(x float32, y float32) (Out float32) {
	Out = x + y
	return
}
