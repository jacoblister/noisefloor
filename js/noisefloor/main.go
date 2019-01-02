package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/jacoblister/noisefloor/engine"
	"github.com/jacoblister/noisefloor/js/frontend"
)

func main() {
	js.Global.Set("noisefloorjs", map[string]interface{}{
		"start":         engine.Start,
		"stop":          engine.Stop,
		"process":       engine.Process,
		"makeProcessor": engine.MakeProcessor,
		"getMIDIEvents": frontend.GetMIDIEvents,
	})
	println("main")
}
