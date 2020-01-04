package processorbasic

import (
	"strconv"
	"strings"

	"github.com/jacoblister/noisefloor/pkg/vdom"
)

const scopeSamples = 100

// Scope - display signal
type Scope struct {
	Trigger int `default:"1" min:"0" max:"1"`
	Skip    int `default:"4" min:"0" max:"200"`

	skipCount  int
	index      int
	samples    []float32
	lastSample float32
}

// Start - init Scope
func (s *Scope) Start(sampleRate int) {
	s.samples = make([]float32, scopeSamples, scopeSamples)
}

// Process - proccess next sample
func (s *Scope) Process(In float32) (Out float32) {
	Out = In

	s.skipCount--
	if s.skipCount >= 0 {
		return
	}
	s.skipCount = s.Skip

	if s.Trigger > 0 && s.index == 0 {
		// wait for zero crossing
		if s.lastSample > 0 || In < 0 {
			s.lastSample = In
			return
		}
	}

	if s.index < scopeSamples {
		s.samples[s.index] = In
		s.index++
	} else {
		if s.Trigger > 0 {
			s.index = 0
		} else {
			s.samples = s.samples[1:]
			s.samples = append(s.samples, In)
		}
	}
	s.lastSample = In

	return
}

// CustomRenderDimentions get the extended dimentions of the scope
func (s *Scope) CustomRenderDimentions() (width int, height int) {
	return 200, 100
}

// Render - render the scope
func (s *Scope) Render() vdom.Element {
	path := strings.Builder{}
	path.WriteString("M0.5," + strconv.Itoa(int(s.samples[0]*-50)+50) + ".5")
	for i := 1; i < scopeSamples; i++ {
		path.WriteString(" L" + strconv.Itoa(i*2) + ".5," + strconv.Itoa(int(s.samples[i]*-50)+50) + ".5")
	}

	pathElement := vdom.MakeElement("path",
		"d", path.String(),
		"stroke", "blue",
		"fill", "none",
	)

	return pathElement
}
