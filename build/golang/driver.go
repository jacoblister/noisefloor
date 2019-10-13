package nf

import (
	"unsafe"

	"C"

	"github.com/jacoblister/noisefloor/app/audiomodule"
	"github.com/jacoblister/noisefloor/pkg/midi"
)
import "reflect"

type driverMidi interface {
	start()
	stop()
	readEvents() []midi.Event
	writeEvents([]midi.Event)
}

type driverAudio interface {
	getDriverMidi() driverMidi
	setDriverMidi(driverMidi driverMidi)
	getAudioProcessor() audiomodule.AudioProcessor
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

func goAudioCallback(driverAudio driverAudio, blockLength int,
	channelInCount int, channelIn unsafe.Pointer,
	channelOutCount int, channelOut unsafe.Pointer) {

	samplesIn := make([][]float32, channelInCount, channelInCount)
	blockLengthInt := int(blockLength)

	for i := 0; i < int(channelInCount); i++ {
		samplesInData := indexPointer(channelIn, i)
		h := &reflect.SliceHeader{Data: uintptr(samplesInData), Len: blockLengthInt, Cap: blockLengthInt}
		s := *(*[]float32)(unsafe.Pointer(h))
		samplesIn[i] = s
	}

	midiIn := driverAudio.getDriverMidi().readEvents()
	samplesOutSlice, midiOut := driverAudio.getAudioProcessor().Process(samplesIn, midiIn)
	driverAudio.getDriverMidi().writeEvents(midiOut)

	for i := 0; i < int(channelOutCount); i++ {
		samplesOutData := indexPointer(channelOut, i)
		h := &reflect.SliceHeader{Data: uintptr(samplesOutData), Len: blockLengthInt, Cap: blockLengthInt}
		s := *(*[]float32)(unsafe.Pointer(h))
		copy(s[:], samplesOutSlice[i][:])
	}
}
