package processorbasic

// Select - between input lines
type Select struct {
	Input int `default:"0" min:"0" max:"1"`
}

// Process - produce next sample
func (s *Select) Process(a float32, b float32) (Out float32) {
	switch s.Input {
	case 0:
		Out = a
	case 1:
		Out = b
	}
	return
}
