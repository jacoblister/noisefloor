package nf

import (
	"unsafe"

	"github.com/jacoblister/noisefloor/audiomodule"
	"github.com/jacoblister/noisefloor/midi"
)

type driverMidi interface {
	start()
	stop()
	readEvents() []midi.Event
	writeEvents([]midi.Event)
}

type driverAudio interface {
	setMidiDriver(driverMidi driverMidi)
	setAudioProcessor(audioProcessor audiomodule.AudioProcessor)
	samplingRate() int
	start()
	stop()
}

// indexPointer is a helper method to dereference a pointer array by index
func indexPointer(ptr unsafe.Pointer, i int) unsafe.Pointer {
	var p uintptr
	var ptrSize = unsafe.Sizeof(&p)

	return unsafe.Pointer(*(**uintptr)(unsafe.Pointer(uintptr(ptr) + uintptr(i)*ptrSize)))
}
