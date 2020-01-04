package processorbasic

// Gain - linear of exponential gain
type Gain struct {
	Level       float32 `default:"1" min:"0" max:"2"`
	Exponential bool
}

// Process - produce next sample
func (g *Gain) Process(In float32, Gai float32) (Out float32) {
	Out = In * Gai * g.Level
	return
}
