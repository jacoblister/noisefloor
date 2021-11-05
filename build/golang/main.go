package nf

import (
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
	// debug.SetGCPercent(-1)

	nf := noiseFloor{driverAudio: driverAudioDefault, driverMidi: driverMidiDefault}
	// nf := noiseFloor{driverAudio: driverAudioDefault, driverMidi: &driverMidiMock{}}

	app.App(&nf, fs)
}
