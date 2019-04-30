package midi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeMidiEventData(t *testing.T) {
	// Given ... dummy time and data
	time := 123
	data := []byte{1, 2, 3}

	// When ...
	eventData := MakeMidiEventData(time, data)

	// Then ... dummy time and data back
	assert.Equal(t, time, eventData.Time)
	assert.Equal(t, data, eventData.Data)
}

func TestMakeMidiEvent_NoType(t *testing.T) {
	// event := MakeMidiEvent(123, []byte{0, 60, 0})
	// assert.Nil(t, event)
	assert.Panics(t, func() { MakeMidiEvent(123, []byte{0, 60, 0}) })
}

func TestMakeMidiEvent_NoteOff(t *testing.T) {
	// Given ... Note off event
	const (
		time      = 123
		channel   = 2
		note      = 60
		velocity  = 0
		eventType = NoteOff
	)
	data := []byte{byte(eventType)<<4 | (channel - 1), note, velocity}

	// When ...
	event := MakeMidiEvent(time, data)

	// Then ... event back, with matching parameters/generic parameters
	noteOffEvent := event.(NoteOffEvent)
	assert.Equal(t, time, noteOffEvent.Time)
	assert.Equal(t, channel, noteOffEvent.Channel)
	assert.Equal(t, note, noteOffEvent.Note)
	assert.Equal(t, velocity, noteOffEvent.Velocity)

	assert.Equal(t, time, event.Data().Time)
	assert.Equal(t, data, event.Data().Data)
	assert.Equal(t, time, event.Generic().Time)
	assert.Equal(t, channel, event.Generic().Channel)
}

func TestMakeMidiEvent_NoteOn(t *testing.T) {
	// Given ... Note on event
	const (
		time      = 123
		channel   = 2
		note      = 60
		velocity  = 100
		eventType = Note
	)
	data := []byte{byte(eventType)<<4 | (channel - 1), note, velocity}

	// When ...
	event := MakeMidiEvent(time, data)

	// Then ... event back, with matching parameters/generic parameters
	noteOnEvent := event.(NoteOnEvent)
	assert.Equal(t, time, noteOnEvent.Time)
	assert.Equal(t, channel, noteOnEvent.Channel)
	assert.Equal(t, note, noteOnEvent.Note)
	assert.Equal(t, velocity, noteOnEvent.Velocity)

	assert.Equal(t, time, event.Data().Time)
	assert.Equal(t, data, event.Data().Data)
	assert.Equal(t, time, event.Generic().Time)
	assert.Equal(t, channel, event.Generic().Channel)
}

func TestMakeMidiEvent_AfterTouch(t *testing.T) {
	// Given ... Aftertouch event
	const (
		time      = 123
		channel   = 2
		note      = 60
		pressure  = 64
		eventType = AfterTouch
	)
	data := []byte{byte(eventType)<<4 | (channel - 1), note, pressure}

	// When ...
	event := MakeMidiEvent(time, data)

	// Then ... event back, with matching parameters/generic parameters
	afterTouchEvent := event.(AfterTouchEvent)
	assert.Equal(t, time, afterTouchEvent.Time)
	assert.Equal(t, channel, afterTouchEvent.Channel)
	assert.Equal(t, note, afterTouchEvent.Note)
	assert.Equal(t, pressure, afterTouchEvent.Pressure)

	assert.Equal(t, time, event.Data().Time)
	assert.Equal(t, data, event.Data().Data)
	assert.Equal(t, time, event.Generic().Time)
	assert.Equal(t, channel, event.Generic().Channel)
}

func TestMakeMidiEvent_Control(t *testing.T) {
	// Given ... Aftertouch event
	const (
		time      = 123
		channel   = 2
		number    = 10
		value     = 11
		eventType = Control
	)
	data := []byte{byte(eventType)<<4 | (channel - 1), number, value}

	// When ...
	event := MakeMidiEvent(time, data)

	// Then ... event back, with matching parameters/generic parameters
	controlEvent := event.(ControlEvent)
	assert.Equal(t, time, controlEvent.Time)
	assert.Equal(t, channel, controlEvent.Channel)
	assert.Equal(t, number, controlEvent.Number)
	assert.Equal(t, value, controlEvent.Value)

	assert.Equal(t, time, event.Data().Time)
	assert.Equal(t, data, event.Data().Data)
	assert.Equal(t, time, event.Generic().Time)
	assert.Equal(t, channel, event.Generic().Channel)
}

func TestMakeMidiEvent_ProgramChange(t *testing.T) {
	// Given ... ChannelPressure event
	const (
		time      = 123
		channel   = 2
		number    = 2
		eventType = Program
	)
	data := []byte{byte(eventType)<<4 | (channel - 1), number}

	// When ...
	event := MakeMidiEvent(time, data)

	// Then ... event back, with matching parameters/generic parameters
	programEvent := event.(ProgramEvent)
	assert.Equal(t, time, programEvent.Time)
	assert.Equal(t, channel, programEvent.Channel)
	assert.Equal(t, number, programEvent.Number)

	assert.Equal(t, time, event.Data().Time)
	assert.Equal(t, data, event.Data().Data)
	assert.Equal(t, time, event.Generic().Time)
	assert.Equal(t, channel, event.Generic().Channel)
}

func TestMakeMidiEvent_ChannelPressure(t *testing.T) {
	// Given ... ChannelPressure event
	const (
		time      = 123
		channel   = 2
		pressure  = 22
		eventType = ChannelPressure
	)
	data := []byte{byte(eventType)<<4 | (channel - 1), pressure}

	// When ...
	event := MakeMidiEvent(time, data)

	// Then ... event back, with matching parameters/generic parameters
	channelPressureEvent := event.(ChannelPressureEvent)
	assert.Equal(t, time, channelPressureEvent.Time)
	assert.Equal(t, channel, channelPressureEvent.Channel)
	assert.Equal(t, pressure, channelPressureEvent.Pressure)

	assert.Equal(t, time, event.Data().Time)
	assert.Equal(t, data, event.Data().Data)
	assert.Equal(t, time, event.Generic().Time)
	assert.Equal(t, channel, event.Generic().Channel)
}

func TestMakeMidiEvent_PitchBend(t *testing.T) {
	// Given ... Pitch Bend Event
	const (
		time      = 123
		channel   = 2
		eventType = PitchBend
	)
	data := []byte{byte(eventType)<<4 | (channel - 1), 0, 0}

	// When ...
	event := MakeMidiEvent(time, data)

	// Then ... pitch bend event back, with matching parameters/generic parameters
	pitchBendEvent := event.(PitchBendEvent)
	assert.Equal(t, time, pitchBendEvent.Time)
	assert.Equal(t, channel, pitchBendEvent.Channel)
	assert.Equal(t, 0, pitchBendEvent.Value)
	assert.Equal(t, -1.0, pitchBendEvent.Normailzed())

	assert.Equal(t, time, event.Data().Time)
	assert.Equal(t, data, event.Data().Data)
	assert.Equal(t, time, event.Generic().Time)
	assert.Equal(t, channel, event.Generic().Channel)

	// Given ... Mid Pitch Bend Range
	data = []byte{byte(eventType)<<4 | (channel - 1), 0x00, 0x40}

	// When ...
	event = MakeMidiEvent(time, data)

	// Then ...
	pitchBendEvent = event.(PitchBendEvent)
	assert.Equal(t, 4096, pitchBendEvent.Value)
	assert.Equal(t, 0.0, pitchBendEvent.Normailzed())

	// Given ... Mid Pitch Bend Range
	data = []byte{byte(eventType)<<4 | (channel - 1), 0x7F, 0x7F}

	// When ...
	event = MakeMidiEvent(time, data)

	// Then ...
	pitchBendEvent = event.(PitchBendEvent)
	assert.Equal(t, 8191, pitchBendEvent.Value)
	assert.Equal(t, 1.0, pitchBendEvent.Normailzed())
}
