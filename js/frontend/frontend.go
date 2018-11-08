package main

import (
	"github.com/bep/gr"
	"github.com/gopherjs/gopherjs/js"
	"github.com/jacoblister/noisefloor/js/frontend/component"
)

func main() {
	// println("run main")

	js.Global.Set("noisefloorjs", map[string]interface{}{
		"getMIDIEvents": component.GetMIDIEvents,
	})
	keyboard := gr.New(new(component.Keyboard))

	// gr.RenderLoop(func() {
	keyboard.Render("react", gr.Props{})
	// })
}
