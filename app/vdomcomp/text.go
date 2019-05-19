package vdomcomp

//Text is a simple Text component for testing
import "github.com/jacoblister/noisefloor/pkg/vdom"

//Text is a SVG text element
type Text struct {
	Text string
}

//Render renders the Text component
func (t *Text) Render() vdom.Element {
	elem := vdom.MakeElement("text",
		"alignment-baseline", "text-before-edge",
		"x", 0,
		"y", 0,
		vdom.MakeTextElement(t.Text),
	)

	return elem
}
