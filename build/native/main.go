package nf

import (
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/jacoblister/noisefloor/component"
	"github.com/jacoblister/noisefloor/component/synth"
)

type noiseFloor struct {
	driverAudio    driverAudio
	driverMidi     driverMidi
	audioProcessor component.AudioProcessor
}

//Main is the native build mainline
func Main() {
	debug.SetGCPercent(-1)

	// nf := noiseFloor{driverAudio: &driverAudioASIO{}, driverMidi: &driverMidiMock{}, audioProcessor: &synth.Engine{}}
	nf := noiseFloor{driverAudio: &driverAudioMock{}, driverMidi: &driverMidiMock{}, audioProcessor: &synth.Engine{}}
	// nf := noiseFloor{driverAudio: &driverAudioJack{}, driverMidi: &driverMidiJack{}, audioProcessor: &synth.Engine{}}
	// nf := noiseFloor{driverAudio: &driverAudioMock{}, driverMidi: &driverMidiJack{}, audioProcessor: &synth.Engine{}}

	nf.driverAudio.setMidiDriver(nf.driverMidi)
	nf.driverAudio.setAudioProcessor(nf.audioProcessor)
	nf.driverMidi.start()
	nf.driverAudio.start()
	nf.audioProcessor.Start(nf.driverAudio.samplingRate())

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	<-signalChannel

	nf.driverAudio.stop()

	// time.Sleep(3 * time.Second)

}
