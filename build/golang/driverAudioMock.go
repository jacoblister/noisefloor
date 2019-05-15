package nf

import (
	"time"

	"github.com/jacoblister/noisefloor/app/audiomodule"
)

type driverAudioMock struct {
	driverMidi     driverMidi
	audioProcessor audiomodule.AudioProcessor
	stopchan       chan bool
	stoppedchan    chan bool
}

func (d *driverAudioMock) mockProcess() {
	defer close(d.stoppedchan)

	println("Mock Audio Start")
	for {
		select {
		case <-d.stopchan:
			println("Mock Audio Stop")
			return
		case <-time.After(1 * time.Second):
			// println("Mock Audio Process...")
			samples := [][]float32{{}}
			midiIn := d.driverMidi.readEvents()
			_, midiOut := d.audioProcessor.Process(samples, midiIn)
			d.driverMidi.writeEvents(midiOut)
			// println()
		}
	}
}

func (d *driverAudioMock) setMidiDriver(driverMidi driverMidi) {
	d.driverMidi = driverMidi
}

func (d *driverAudioMock) setAudioProcessor(audioProcessor audiomodule.AudioProcessor) {
	d.audioProcessor = audioProcessor
}

func (d *driverAudioMock) start() {
	d.stopchan = make(chan bool)
	d.stoppedchan = make(chan bool)

	go d.mockProcess()
}

func (d *driverAudioMock) stop() {
	println("stop")
	close(d.stopchan)
	<-d.stoppedchan
}

func (d *driverAudioMock) samplingRate() int {
	return 44100
}
