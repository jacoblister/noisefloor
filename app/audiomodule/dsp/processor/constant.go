package processor

// Constant - single specified value
type Constant struct {
	Value float32
}

// Start - init module
func (c *Constant) Start(sampleRate int) {}

// Process - produce next sample
func (c *Constant) Process() (output float32) {
	output = c.Value
	return
}
