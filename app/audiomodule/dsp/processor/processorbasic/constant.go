package processorbasic

// Constant - single specified value
type Constant struct {
	Value float32 `default:"1" min:"0" max:"10"`
}

// Process - produce next sample
func (c *Constant) Process() (Out float32) {
	Out = c.Value
	return
}
