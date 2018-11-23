package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/jacoblister/noisefloor/engine"
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

func main() {
	js.Global.Set("noisefloorjs", map[string]interface{}{
		"start":         engine.Start,
		"stop":          engine.Stop,
		"process":       engine.Process,
		"makeProcessor": engine.MakeProcessor,
	})

	println("loaded engine.main")
}
