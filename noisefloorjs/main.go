package main

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	. "github.com/jacoblister/noisefloor/common"
	"github.com/jacoblister/noisefloor/engine/processor"
)

var oscillator processor.Oscillator

func start(sampleRate int) {
	fmt.Println("do DSP start, sample rate:", sampleRate)
	oscillator.Start(sampleRate)
}

func stop() {
	fmt.Println("do DSP stop")
}

func process(samplesIn *js.Object, samplesOut *js.Object, midiIn *js.Object, midiOut *js.Object) {
	midiEvent := MidiEvent{Time: 1000, Data: []byte{1, 2, 3}}
	fmt.Println(midiEvent)

	var len int = samplesOut.Index(0).Length()
	for i := 0; i < len; i++ {
		var sample = oscillator.Process()
		samplesOut.Index(0).SetIndex(i, sample)
		samplesOut.Index(1).SetIndex(i, sample)
	}
}

func main() {
	js.Global.Set("noisefloorjs", map[string]interface{}{
		"start":   start,
		"stop":    stop,
		"process": process,
	})
	fmt.Println("main")
}
