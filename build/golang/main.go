package nf

import (
	"runtime/debug"

	"github.com/jacoblister/noisefloor/app"
	"github.com/jacoblister/noisefloor/component"
)

type noiseFloor struct {
	driverAudio driverAudio
	driverMidi  driverMidi
	// audioProcessor component.AudioProcessor
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

// Main is the applicatoin entry point
func Main() {
	debug.SetGCPercent(-1)

	nf := noiseFloor{driverAudio: &driverAudioMock{}, driverMidi: &driverMidiMock{}}
	// nf := noiseFloor{driverAudio: &driverAudioASIO{}, driverMidi: &driverMidiWDM{}}

	app.App(&nf)
}

// //Main is the native build mainline
// func Main() {
// 	debug.SetGCPercent(-1)
//
// 	// nf := noiseFloor{driverAudio: &driverAudioMock{}, driverMidi: &driverMidiMock{}, audioProcessor: &synth.Engine{}}
// 	// nf := noiseFloor{driverAudio: &driverAudioMock{}, driverMidi: &driverMidiWDM{}, audioProcessor: &synth.Engine{}}
// 	nf := noiseFloor{driverAudio: &driverAudioASIO{}, driverMidi: &driverMidiWDM{}, audioProcessor: &synth.Engine{}}
// 	// nf := noiseFloor{driverAudio: &driverAudioJack{}, driverMidi: &driverMidiJack{}, audioProcessor: &synth.Engine{}}
// 	// nf := noiseFloor{driverAudio: &driverAudioMock{}, driverMidi: &driverMidiJack{}, audioProcessor: &synth.Engine{}}
//
// 	nf.driverAudio.setMidiDriver(nf.driverMidi)
// 	nf.driverAudio.setAudioProcessor(nf.audioProcessor)
// 	nf.driverMidi.start()
// 	nf.driverAudio.start()
// 	nf.audioProcessor.Start(nf.driverAudio.samplingRate())
//
// 	signalChannel := make(chan os.Signal, 2)
// 	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
// 	<-signalChannel
//
// 	nf.driverAudio.stop()
// 	nf.driverMidi.stop()
//
// 	// time.Sleep(3 * time.Second)
//
// }
