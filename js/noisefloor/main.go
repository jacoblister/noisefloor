package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/jacoblister/noisefloor/js/engine"
)

func main() {
	js.Global.Set("noisefloorjs", map[string]interface{}{
		"makeMidiEvent": engine.MakeMidiEvent,
		"jsMidiEvent":   engine.JSMidiEvent,
		"start":         engine.Start,
		"stop":          engine.Stop,
		"process":       engine.Process,
	})
	println("main")
}
