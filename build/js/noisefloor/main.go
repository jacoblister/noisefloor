package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/jacoblister/noisefloor/build/js/frontend"
	"github.com/jacoblister/noisefloor/component"
	"github.com/jacoblister/noisefloor/component/synth"
	"github.com/jacoblister/noisefloor/midi"
)

// main exports functions to javascript
// (only for dead code elimination, exports are in export.inc.js)
func main() {
	js.Global.Set("noisefloorjs", map[string]interface{}{
		"MakeProcessor": synth.MakeProcessor,
		"MakeComponent": component.MakeComponent,
		"MakeMidiEvent": midi.MakeMidiEvent,
		"GetMIDIEvents": frontend.GetMIDIEvents,
	})
	println("main")
}
