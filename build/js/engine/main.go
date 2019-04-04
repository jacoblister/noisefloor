package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/jacoblister/noisefloor/component"
	"github.com/jacoblister/noisefloor/component/synth"
	"github.com/jacoblister/noisefloor/midi"
)

// func ProcessJs(samplesIn *js.Object, samplesOut *js.Object, midiIn *js.Object, midiOut *js.Object) {
// 	// midiEvent := MidiEvent{Time: 1000, Data: []byte{1, 2, 3}}
//
// 	var len int = samplesOut.Index(0).Length()
// 	for i := 0; i < len; i++ {
// 		var sample = oscillator.Process()
// 		samplesOut.Index(0).SetIndex(i, sample)
// 		samplesOut.Index(1).SetIndex(i, sample)
// 	}
// }

// TestCallProcess is a dummy call to process to evaluate transpiled javascript
func TestCallProcess() {
	var Engine synth.Engine

	var samplesIn [][]float32
	var samplesOut [][]float32
	var midiIn []midi.Event
	var midiOut []midi.Event

	Engine.Process(samplesIn, samplesOut, midiIn, &midiOut)
}

// main exports functions to javascript
// (only for dead code elimination, exports are in export.inc.js)
func main() {
	js.Global.Set("noisefloorjs", map[string]interface{}{
		"TestCallProcess": TestCallProcess,
		"MakeProcessor":   synth.MakeProcessor,
		"MakeComponent":   component.MakeComponent,
		"MakeMidiEvent":   midi.MakeMidiEvent,
	})

	slice := []byte{1, 2, 3, 4}
	slice[0] = 0

	println("loaded engine.main")
}
