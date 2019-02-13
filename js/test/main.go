package main

// type MidiEvent struct {
// 	Time int
// 	Data []byte
// }

import (
	"github.com/jacoblister/noisefloor/common/midi"
)

// func MakeMidiEvent(Time int, Data []byte) (result *common.NewMidiEvent) {
// 	return &common.NewMidiEvent{Time, Data}
// }

// func testMidiEvent(midiEvent []NewMidiEvent) {
// 	newEvent := MidiEvent{123, make([]byte, 3)}
// 	newEvent.Data[0] = 80
//
// 	midiEvent[len(midiEvent)+1] = newEvent
// }

const midiArraySize = 128
const arraySize = 128

func SumArray(data [arraySize]int) (result int) {
	var total int
	// data = append(data, 1)

	for i := 0; i < len(data); i++ {
		total += (data)[i]
	}
	return total
}

func SumArrayPtr(data *[arraySize]int) (result int) {
	var total int
	// data = append(data, 1)

	for i := 0; i < len(data); i++ {
		total += (data)[i]
	}
	return total
}

func SumSlice(data []int) (result int) {
	var total int
	// data = append(data, 30)

	for i := 0; i < len(data); i++ {
		total += (data)[i]
	}
	return total
}

func SumSlicePtr(data *[]int) (result int) {
	var total int
	// *data = append(*data, 30)

	for i := 0; i < len(*data); i++ {
		total += (*data)[i]
	}
	return total
}

func main() {
	// midiEvent := make([]MidiEvent, 0)
	// testMidiEvent(midiEvent)

	arrayValues := [arraySize]int{10, 20}

	total := SumArray(arrayValues)
	println("array total:", total)
	println("array len:", len(arrayValues))
	println()

	total = SumArrayPtr(&arrayValues)
	println("array ptr total:", total)
	println("array ptr len:", len(arrayValues))
	println()

	sliceValues := arrayValues[:]
	sliceValues = sliceValues[:1]

	total = SumSlice(sliceValues)
	println("slice total:", total)
	println("slice len:", len(sliceValues))
	println()

	total = SumSlicePtr(&sliceValues)
	println("slice ptr total:", total)
	println("slice ptr len:", len(sliceValues))
	println()

	midiEvent := midi.MakeMidiEventData(0, []byte{127, 0, 0})
	println(midiEvent.Time)

	midiEventLinearA := []int{0, 127, 0, 0}
	midiEventLinearA[0] = 0

	midiEventLinearB := [4]int{0, 127, 0, 0}
	midiEventLinearB[0] = 0

	// values =
	// testslice()
	println("maininator")

	// js.Global.Set("noisefloorjs", map[string]interface{}{
	// 	"MakeMidiEvent": common.MakeMidiEvent,
	// })

	// js.Global.Set("noisefloorjs", map[string]interface{}{
	// 	"testslice": testslice,
	// })
}
