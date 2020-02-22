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

	skipCount     int
	index         int
	samples       [2][]float32
	lastSample    float32
	connectedMask int
}

// Start - init Scope
func (s *Scope) Start(sampleRate int, connectedMask int) {
	s.samples[0] = make([]float32, scopeSamples, scopeSamples)
	s.samples[1] = make([]float32, scopeSamples, scopeSamples)
	s.connectedMask = connectedMask
}

// Process - proccess next sample
func (s *Scope) Process(InA float32, InB float32) (OutA float32, OutB float32) {
	OutA = InA
	OutB = InB

	s.skipCount--
	if s.skipCount >= 0 {
		return
	}
	s.skipCount = s.Skip

	if s.Trigger > 0 && s.index == 0 {
		// wait for zero crossing
		if s.lastSample > 0 || InA < 0 {
			s.lastSample = InA
			return
		}
	}

	if s.index < scopeSamples {
		s.samples[0][s.index] = InA
		s.samples[1][s.index] = InB
		s.index++
	} else {
		if s.Trigger > 0 {
			s.index = 0
		} else {
			s.samples[0] = s.samples[0][1:]
			s.samples[0] = append(s.samples[0], InA)
			s.samples[1] = s.samples[1][1:]
			s.samples[1] = append(s.samples[1], InA)
		}
	}
	s.lastSample = InA

	return
}

// CustomRenderDimentions get the extended dimentions of the scope
func (s *Scope) CustomRenderDimentions() (width int, height int) {
	return 200, 100
}

// Render - render the scope
func (s *Scope) Render() vdom.Element {
	pathElements := []vdom.Element{}

	for i := 0; i < 2; i++ {
		stroke := "blue"
		if i != 0 {
			stroke = "darkcyan"
		}
		if (s.connectedMask & (1 << uint(i))) != 0 {
			path := strings.Builder{}

			path.WriteString("M0.5," + strconv.Itoa(int(s.samples[i][0]*-50)+50) + ".5")
			for j := 1; j < scopeSamples; j++ {
				path.WriteString(" L" + strconv.Itoa(j*2) + ".5," + strconv.Itoa(int(s.samples[i][j]*-50)+50) + ".5")
			}

			pathElement := vdom.MakeElement("path",
				"d", path.String(),
				"stroke", stroke,
				"fill", "none",
			)
			pathElements = append(pathElements, pathElement)
		}
	}

	element := vdom.MakeElement("g",
		pathElements)

	return element
}
