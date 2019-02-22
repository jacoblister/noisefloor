package main

import "C"
import (
	"reflect"
	"unsafe"
)

//export goProcess
func goProcess(arg unsafe.Pointer, blockSize int,
	channelsInCount int, samplesIn unsafe.Pointer,
	channelsOutCount int, samplesOut unsafe.Pointer,
	midiInCount int, MidiIn unsafe.Pointer,
	midiOutCount int, MidiOut unsafe.Pointer) {
	samplesInSlice := [][]float32{}

	for i := 0; i < channelsInCount; i++ {
		h := &reflect.SliceHeader{Data: uintptr(samplesIn), Len: 2, Cap: 2}
		NewSlice := *(*[]float32)(unsafe.Pointer(h))
		samplesInSlice[i] = NewSlice
		println(NewSlice[0], NewSlice[1])
	}
	// println(NewSlice[0], NewSlice[1])

	dp := *(*driverAudioJack)(arg)
	dp.audioProcessor.Start(1000)

	println(blockSize)
}
