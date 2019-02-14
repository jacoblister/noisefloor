package midi

import (
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeMidiEventData(t *testing.T) {
	time := 123
	data := []byte{1, 2, 3}

	eventData := MakeMidiEventData(time, data)
	assert.Equal(t, eventData.Time, time)
	assert.Equal(t, eventData.Data, data)
}

func TestMakeMidiEvent_NoType(t *testing.T) {
	assert.Panics(t, func() { MakeMidiEvent(123, []byte{0, 60, 0}) })
}

func TestMakeMidiEvent_NoteOn(t *testing.T) {
	const (
		time      = 123
		channel   = 2
		note      = 60
		velocity  = 100
		eventType = Note
	)
	data := []byte{byte(eventType)<<4 | (channel - 1), note, velocity}
	event := MakeMidiEvent(time, data)

	noteOnEvent := event.(*NoteOnEvent)
	assert.Equal(t, noteOnEvent.Time, time)
	assert.Equal(t, noteOnEvent.Channel, channel)
	assert.Equal(t, noteOnEvent.Note, note)
	assert.Equal(t, noteOnEvent.Velocity, velocity)

	assert.Equal(t, event.Data().Time, time)
	assert.Equal(t, event.Data().Data, data)
	assert.Equal(t, event.Generic().Time, time)
	assert.Equal(t, event.Generic().Channel, channel)
}

// Benchmarks
func BenchmarkMakeArrayMidiEvent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		midiEvent := [...]int{0, 127, 0, 0}
		midiEvent[0] = 0
	}
}

func BenchmarkMakeSliceMidiEvent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		midiEvent := []int{0, 127, 0, 0}
		midiEvent[0] = 0
	}
}

func BenchmarkMakeMidiEventData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeMidiEventData(0, []byte{127, 0, 0})
	}
}

func BenchmarkMakeMidiEvent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeMidiEvent(0, []byte{9 << 4, 0, 0})
	}
}

// Array/Slice/Pointer benchmarks for MidiEvent vector
const arraySize = 128

var testEventArray [arraySize]EventData
var testEventSlice []EventData

func BenchmarkMidiAppendSlice(b *testing.B) {
	midiEvent := MakeMidiEventData(0, []byte{127, 0, 0})

	for i := 0; i < b.N; i++ {
		midiEvents := make([]*EventData, 0, arraySize)
		for j := 0; j < arraySize; j++ {
			midiEvents = append(midiEvents, midiEvent)
		}
	}
}

func BenchmarkMidiAppendArray(b *testing.B) {
	midiEvent := MakeMidiEventData(0, []byte{127, 0, 0})

	for i := 0; i < b.N; i++ {
		midiEvents := [arraySize]*EventData{}
		for j := 0; j < arraySize; j++ {
			midiEvents[j] = midiEvent
		}
	}
}

func TestMain(m *testing.M) {
	for i := 0; i < arraySize; i++ {
		time := rand.Intn(10000)
		testEventArray[i] = *MakeMidiEventData(time, []byte{0, 0, 0})
	}
	testEventSlice = testEventArray[:]
	os.Exit(m.Run())
}

func MaxTimeArray(event [arraySize]EventData) int {
	var maxTime int

	for i := 0; i < len(event); i++ {
		if event[i].Time > maxTime {
			maxTime = event[i].Time
		}
	}
	return maxTime
}

func BenchmarkMaxTimeArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxTimeArray(testEventArray)
	}
}

func MaxTimeArrayPtr(event *[arraySize]EventData) int {
	var maxTime int

	for i := 0; i < len(event); i++ {
		if event[i].Time > maxTime {
			maxTime = event[i].Time
		}
	}
	return maxTime
}

func BenchmarkMaxTimeArrayPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxTimeArrayPtr(&testEventArray)
	}
}

func MaxTimeSlice(event []EventData) int {
	var maxTime int

	for i := 0; i < len(event); i++ {
		if event[i].Time > maxTime {
			maxTime = event[i].Time
		}
	}
	return maxTime
}

func BenchmarkMaxTimeSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxTimeSlice(testEventSlice)
	}
}

func MaxTimeSlicePtr(event *[]EventData) int {
	var maxTime int

	for i := 0; i < len(*event); i++ {
		if (*event)[i].Time > maxTime {
			maxTime = (*event)[i].Time
		}
	}
	return maxTime
}

func BenchmarkMaxTimeSlicePtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxTimeSlicePtr(&testEventSlice)
	}
}
