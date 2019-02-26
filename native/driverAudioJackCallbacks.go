package main

import "C"
import (
	"reflect"
	"unsafe"

	"github.com/jacoblister/noisefloor/common/midi"
)

//export goProcess
func goProcess(arg unsafe.Pointer, blockSize int,
	channelInCount int, channelIn unsafe.Pointer,
	channelOutCount int, channelOut unsafe.Pointer,
	midiInCount int, MidiIn unsafe.Pointer,
	midiOutCount int, MidiOut unsafe.Pointer) {
	samplesInSlice := make([][]float32, channelInCount, channelInCount)
	samplesOutSlice := make([][]float32, channelOutCount, channelOutCount)
	midiInSlice := make([]midi.Event, 0, 0)
	midiOutSlice := make([]midi.Event, 0, 0)

	for i := 0; i < channelInCount; i++ {
		samplesIn := indexPointer(channelIn, i)
		h := &reflect.SliceHeader{Data: uintptr(samplesIn), Len: blockSize, Cap: blockSize}
		s := *(*[]float32)(unsafe.Pointer(h))
		samplesInSlice[i] = s
	}

	for i := 0; i < channelOutCount; i++ {
		samplesOut := indexPointer(channelOut, i)
		h := &reflect.SliceHeader{Data: uintptr(samplesOut), Len: blockSize, Cap: blockSize}
		s := *(*[]float32)(unsafe.Pointer(h))
		samplesOutSlice[i] = s
	}

	dp := *(*driverAudioJack)(arg)
	dp.audioProcessor.Process(samplesInSlice, samplesOutSlice, midiInSlice, &midiOutSlice)

	println(blockSize)
}
