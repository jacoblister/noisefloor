package engine

// ProcessorParameter defines a processor setting and its limits
type ProcessorParameter struct {
	name  string
	value float32
	min   float32
	max   float32
	enum  []string
}
