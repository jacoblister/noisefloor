package main

import (
	"time"

	"github.com/jacoblister/noisefloor/common/midi"
	"github.com/jacoblister/noisefloor/component"
	"github.com/jacoblister/noisefloor/engine"
)

type driverMidi interface {
	start()
	stop()
	readEvents() []midi.Event
	writeEvents([]midi.Event)
}

type driverAudio interface {
	setMidiDriver(driverMidi driverMidi)
	setAudioProcessor(audioProcessor component.AudioProcessor)
	start()
	stop()
}

type noiseFloor struct {
	driverAudio    driverAudio
	driverMidi     driverMidi
	audioProcessor component.AudioProcessor
}

func main() {
	nf := noiseFloor{driverAudio: &driverAudioMock{}, driverMidi: &driverMidiMock{}, audioProcessor: &engine.Engine{}}

	nf.driverAudio.setMidiDriver(nf.driverMidi)
	nf.driverAudio.setAudioProcessor(nf.audioProcessor)
	nf.driverAudio.start()

	time.Sleep(3 * time.Second)

	nf.driverAudio.stop()

	// time.Sleep(3 * time.Second)

}
