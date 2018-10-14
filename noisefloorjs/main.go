package main

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/jacoblister/noisefloor/engine/processor"
)

var oscillator processor.Oscillator

func start(sampleRate int) {
	fmt.Println("do DSP start:", sampleRate)
	oscillator.Start(sampleRate)
	fmt.Println("done DSP start")
}

func stop() {
	fmt.Println("do DSP stop")
}

func process(samplesIn *js.Object, samplesOut *js.Object) {
	fmt.Println("do DSP process", samplesOut.Length())
	for i := 0; i < samplesOut.Length(); i++ {
		samplesOut.SetIndex(i, oscillator.Process())
	}
	fmt.Println("done DSP process")
}

func main() {
	js.Global.Set("noisefloorjs", map[string]interface{}{
		"start":   start,
		"stop":    stop,
		"process": process,
	})
	fmt.Println("main")
}
