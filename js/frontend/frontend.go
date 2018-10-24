package frontend

import (
	"honnef.co/go/js/dom"
)

// RenderFrontend renders the frontend into the top level DOM Document
func RenderFrontend() {
	d := dom.GetWindow().Document()
	div := d.CreateElement("div").(*dom.HTMLDivElement)
	div.Style().SetProperty("color", "red", "")
	div.SetTextContent("Noisefloor Frontend")
}
