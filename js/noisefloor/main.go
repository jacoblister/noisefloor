package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/jacoblister/noisefloor/common/midi"
	"github.com/jacoblister/noisefloor/component"
	"github.com/jacoblister/noisefloor/engine"
	"github.com/jacoblister/noisefloor/js/frontend"
)

// main exports functions to javascript
// (only for dead code elimination, exports are in export.inc.js)
func main() {
	js.Global.Set("noisefloorjs", map[string]interface{}{
		"MakeProcessor": engine.MakeProcessor,
		"MakeComponent": component.MakeComponent,
		"MakeMidiEvent": midi.MakeMidiEvent,
		"GetMIDIEvents": frontend.GetMIDIEvents,
	})
	println("main")
}
