package main

import (
	"time"
	"unsafe"

	"github.com/jacoblister/noisefloor/component"
	"github.com/jacoblister/noisefloor/component/synth"
)

type noiseFloor struct {
	driverAudio    driverAudio
	driverMidi     driverMidi
	audioProcessor component.AudioProcessor
}

// indexPointer is a helper method to dereference a pointer array by index
func indexPointer(ptr unsafe.Pointer, i int) unsafe.Pointer {
	var p uintptr
	var ptrSize = unsafe.Sizeof(&p)

	return unsafe.Pointer(*(**uintptr)(unsafe.Pointer(uintptr(ptr) + uintptr(i)*ptrSize)))
}

func main() {
	// nf := noiseFloor{driverAudio: &driverAudioMock{}, driverMidi: &driverMidiMock{}, audioProcessor: &synth.Engine{}}
	nf := noiseFloor{driverAudio: &driverAudioJack{}, driverMidi: &driverMidiMock{}, audioProcessor: &synth.Engine{}}

	nf.driverAudio.setMidiDriver(nf.driverMidi)
	nf.driverAudio.setAudioProcessor(nf.audioProcessor)
	nf.driverAudio.start()

	time.Sleep(30 * time.Second)

	nf.driverAudio.stop()

	// time.Sleep(3 * time.Second)

}
