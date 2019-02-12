package main

import (
	"testing"

	"github.com/jacoblister/noisefloor/common"
)

func BenchmarkMakeLinearMidiEvent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < midiArraySize; j++ {
			midiEvent := []int{0, 127, 0, 0}
			midiEvent[0] = 0
		}
	}
}

func BenchmarkMakeMidiEvent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < midiArraySize; j++ {
			midiEvent := common.MakeMidiEvent(0, []byte{127, 0, 0})
			midiEvent.Time = 0
		}
	}
}

func BenchmarkMidiAppend(b *testing.B) {
	midiEvents := make([]*common.NewMidiEvent, midiArraySize, midiArraySize)

	midiEvent := common.MakeMidiEvent(0, []byte{127, 0, 0})
	for i := 0; i < b.N; i++ {
		for j := 0; j < midiArraySize; j++ {
			midiEvents = append(midiEvents, midiEvent)
		}
	}
}

func BenchmarkMidiAppendArray(b *testing.B) {
	midiEvents := [arraySize]*common.NewMidiEvent{}

	midiEvent := common.MakeMidiEvent(0, []byte{127, 0, 0})
	for i := 0; i < b.N; i++ {
		for j := 0; j < arraySize; j++ {
			midiEvents[j] = midiEvent
		}
	}
}

func BenchmarkSumArray(b *testing.B) {
	arrayValues := [arraySize]int{10, 20}

	for i := 0; i < b.N; i++ {
		SumArray(arrayValues)
	}
}

func BenchmarkSumArrayPtr(b *testing.B) {
	arrayValues := [arraySize]int{10, 20}

	for i := 0; i < b.N; i++ {
		SumArrayPtr(&arrayValues)
	}
}

func BenchmarkSumSlice(b *testing.B) {
	sliceValues := make([]int, arraySize, arraySize)
	sliceValues[0] = 10
	sliceValues[1] = 20

	for i := 0; i < b.N; i++ {
		SumSlice(sliceValues)
	}
}

func BenchmarkSumSlicePtr(b *testing.B) {
	sliceValues := make([]int, arraySize, arraySize)
	sliceValues[0] = 10
	sliceValues[1] = 20

	for i := 0; i < b.N; i++ {
		SumSlicePtr(&sliceValues)
	}
}
