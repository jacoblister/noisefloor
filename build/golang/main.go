package nf

import (
	"runtime/debug"

	"github.com/jacoblister/noisefloor/app"
	"github.com/jacoblister/noisefloor/app/audiomodule"
)

type noiseFloor struct {
	driverAudio driverAudio
	driverMidi  driverMidi
}

func (nf *noiseFloor) Start(hardwareDevices app.HardwareDevices, audioProcessor audiomodule.AudioProcessor) {
	nf.driverAudio.setDriverMidi(nf.driverMidi)
	nf.driverAudio.setAudioProcessor(audioProcessor)
	nf.driverMidi.start()
	nf.driverAudio.start()
	audioProcessor.Start(nf.driverAudio.samplingRate())
}

func (nf *noiseFloor) Stop(hardwareDevices app.HardwareDevices) {
	nf.driverAudio.stop()
	nf.driverMidi.stop()
}

// Main is the application entry point
func Main() {
	debug.SetGCPercent(-1)

	// nf := noiseFloor{driverAudio: &driverAudioMock{}, driverMidi: &driverMidiMock{}}
	// nf := noiseFloor{driverAudio: &driverAudioWASAPI{}, driverMidi: &driverMidiMock{}}
	// nf := noiseFloor{driverAudio: &driverAudioASIO{}, driverMidi: &driverMidiWDM{}}
	// nf := noiseFloor{driverAudio: &driverAudioASIO{}, driverMidi: &driverMidiMock{}}
	nf := noiseFloor{driverAudio: &driverAudioWASAPI{}, driverMidi: &driverMidiMock{}}
	// nf := noiseFloor{driverAudio: &driverAudioMock{}, driverMidi: &driverMidiWDM{}}
	// nf := noiseFloor{driverAudio: &driverAudioJack{}, driverMidi: &driverMidiJack{}}
	app.App(&nf)
}
