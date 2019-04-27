package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/jacoblister/noisefloor/app"
	"github.com/jacoblister/noisefloor/component"
	"github.com/jacoblister/noisefloor/component/synth"
	"github.com/jacoblister/noisefloor/midi"
)

type noiseFloor struct {
	audioProcessor component.AudioProcessor
}

var nf noiseFloor

//GetAudioProcessor returns the audioProcessor to external javascript
func GetAudioProcessor() component.AudioProcessor {
	return nf.audioProcessor
}

func (nf *noiseFloor) Start(hardwareDevices app.HardwareDevices, audioProcessor component.AudioProcessor) {
	nf.audioProcessor = audioProcessor
	// js.Global.Call("startAudio")
}

func (nf *noiseFloor) Stop(hardwareDevices app.HardwareDevices) {
	// js.Global.Call("stopAudio")
}

// main in the application entry point
func main() {
	js.Global.Set("noisefloorjs", map[string]interface{}{
		"MakeProcessor":     synth.MakeProcessor,
		"MakeComponent":     component.MakeComponent,
		"MakeMidiEvent":     midi.MakeMidiEvent,
		"GetAudioProcessor": app.GetAudioProcessor,
	})

	app.App(&nf)

	println("main")
}
