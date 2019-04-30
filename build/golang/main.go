package nf

import (
	"runtime/debug"

	"github.com/jacoblister/noisefloor/app"
	"github.com/jacoblister/noisefloor/component"
)

type noiseFloor struct {
	driverAudio driverAudio
	driverMidi  driverMidi
}

func (nf *noiseFloor) Start(hardwareDevices app.HardwareDevices, audioProcessor component.AudioProcessor) {
	nf.driverAudio.setMidiDriver(nf.driverMidi)
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

	nf := noiseFloor{driverAudio: &driverAudioMock{}, driverMidi: &driverMidiMock{}}
	// nf := noiseFloor{driverAudio: &driverAudioASIO{}, driverMidi: &driverMidiWDM{}}
	// nf := noiseFloor{driverAudio: &driverAudioMock{}, driverMidi: &driverMidiWDM{}}

	app.App(&nf)
}
