package processor

// Divide - divide x by y
type Divide struct {
	Dummy int
}

// Start - init module
func (d *Divide) Start(sampleRate int) {}

// Process - produce next sample
func (d *Divide) Process(x float32, y float32) (output float32) {
	if y == 0 {
		return 0
	}
	output = x / y
	return
}
