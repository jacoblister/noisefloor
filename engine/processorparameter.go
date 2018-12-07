package engine

import (
	. "github.com/jacoblister/noisefloor/common"
)

// ProcessorParameter defines a processor setting and its limits
type ProcessorParameter struct {
	name  string
	value AudioFloat
	min   AudioFloat
	max   AudioFloat
	enum  []string
}
