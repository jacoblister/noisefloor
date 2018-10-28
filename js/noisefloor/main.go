package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/jacoblister/noisefloor/engine"
	"github.com/jacoblister/noisefloor/js/frontend"
)

func main() {
	js.Global.Set("noisefloorjs", map[string]interface{}{
		// "makeMidiEvent": engine.MakeMidiEvent,
		// "logMidiEvent":  engine.LogMidiEvent,
		// "jsMidiEvent":   engine.JSMidiEvent,
		"start":          engine.Start,
		"stop":           engine.Stop,
		"process":        engine.Process,
		"getMIDIEvents":  frontend.GetMIDIEvents,
		"renderFrontEnd": frontend.RenderFrontend,
	})
	println("main")
}
