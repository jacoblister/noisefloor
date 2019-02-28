package main

import (
	"runtime/debug"
	"time"

	"github.com/jacoblister/noisefloor/component"
	"github.com/jacoblister/noisefloor/component/synth"
)

type noiseFloor struct {
	driverAudio    driverAudio
	driverMidi     driverMidi
	audioProcessor component.AudioProcessor
}

func main() {
	debug.SetGCPercent(-1)

	// nf := noiseFloor{driverAudio: &driverAudioJack{}, driverMidi: &driverMidiMock{}, audioProcessor: &synth.Engine{}}
	nf := noiseFloor{driverAudio: &driverAudioJack{}, driverMidi: &driverMidiJack{}, audioProcessor: &synth.Engine{}}
	// nf := noiseFloor{driverAudio: &driverAudioMock{}, driverMidi: &driverMidiJack{}, audioProcessor: &synth.Engine{}}

	nf.driverAudio.setMidiDriver(nf.driverMidi)
	nf.driverAudio.setAudioProcessor(nf.audioProcessor)
	nf.driverMidi.start()
	nf.driverAudio.start()
	nf.audioProcessor.Start(nf.driverAudio.samplingRate())

	time.Sleep(1000 * time.Second)

	nf.driverAudio.stop()

	// time.Sleep(3 * time.Second)

}
