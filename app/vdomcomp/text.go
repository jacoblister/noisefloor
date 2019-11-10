package vdomcomp

//Text is a simple Text component for testing
import "github.com/jacoblister/noisefloor/pkg/vdom"

//Text is a SVG text element
type Text struct {
	Text string
	X    int
	Y    int
}

//Render renders the Text component
func (t *Text) Render() vdom.Element {
	elem := vdom.MakeElement("text",
		"alignment-baseline", "text-before-edge",
		"x", t.X,
		"y", t.Y,
		vdom.MakeTextElement(t.Text),
	)

	return elem
}
