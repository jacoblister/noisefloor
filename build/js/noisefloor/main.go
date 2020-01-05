package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/jacoblister/noisefloor/app"
	"github.com/jacoblister/noisefloor/app/audiomodule"
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor/processorfactory"
	"github.com/jacoblister/noisefloor/pkg/midi"
)

type noiseFloor struct {
	audioProcessor audiomodule.AudioProcessor
}

var nf noiseFloor

//GetAudioProcessor returns the audioProcessor to external javascript
func GetAudioProcessor() audiomodule.AudioProcessor {
	return nf.audioProcessor
}

func (nf *noiseFloor) Start(hardwareDevices app.HardwareDevices, audioProcessor audiomodule.AudioProcessor) {
	nf.audioProcessor = audioProcessor
}

func (nf *noiseFloor) Stop(hardwareDevices app.HardwareDevices) {
}

// main in the application entry point
func main() {
	js.Global.Set("noisefloorjs", map[string]interface{}{
		"MakeProcessor":     processorfactory.MakeProcessor,
		"MakeComponent":     audiomodule.MakeComponent,
		"MakeMidiEvent":     midi.MakeMidiEvent,
		"GetAudioProcessor": app.GetAudioProcessor,
	})

	app.App(&nf, &fs)
}
