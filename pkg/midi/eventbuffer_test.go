package midi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeEventBuffer(t *testing.T) {
	const (
		time     = 0
		channel  = 2
		note     = 60
		velocity = 100
	)
	eventBuffer := []Event{}
	eventBuffer = append(eventBuffer, &NoteOnEvent{GenericEvent: GenericEvent{Time: time, Channel: channel}, Note: note, Velocity: velocity})
	eventBuffer = append(eventBuffer, &NoteOnEvent{GenericEvent: GenericEvent{Time: time, Channel: channel}, Note: note + 1, Velocity: velocity})

	byteBuffer := EncodeEventBuffer(eventBuffer)

	expectedByteBuffer := []byte{
		3, 0, 0, 0, 0, byte(Note)<<4 | (channel - 1), note, velocity,
		3, 0, 0, 0, 0, byte(Note)<<4 | (channel - 1), note + 1, velocity,
		0,
	}
	assert.Equal(t, expectedByteBuffer, byteBuffer)
}

func TestDecodeByteBuffer(t *testing.T) {
	// Given ... termination byte only
	byteBuffer := []byte{0}
	eventBuffer := DecodeByteBuffer(byteBuffer)

	// When ... Then ...
	expectedEventBuffer := []Event{}
	assert.Equal(t, expectedEventBuffer, eventBuffer)

	const (
		time      = 0
		channel   = 2
		note      = 60
		velocity  = 100
		eventType = Note
	)
	byteBuffer = []byte{
		3, 0, 0, 0, 0, byte(Note)<<4 | (channel - 1), note, velocity,
		3, 0, 0, 0, 0, byte(Note)<<4 | (channel - 1), note + 1, velocity,
		0,
	}

	eventBuffer = DecodeByteBuffer(byteBuffer)

	expectedEventBuffer = []Event{}
	expectedEventBuffer = append(expectedEventBuffer, NoteOnEvent{GenericEvent: GenericEvent{Time: time, Channel: channel}, Note: note, Velocity: velocity})
	expectedEventBuffer = append(expectedEventBuffer, NoteOnEvent{GenericEvent: GenericEvent{Time: time, Channel: channel}, Note: note + 1, Velocity: velocity})

	assert.Equal(t, expectedEventBuffer, eventBuffer)
}

func BenchmarkEncodeEventBuffer(b *testing.B) {
	eventBuffer := []Event{}
	eventBuffer = append(eventBuffer, NoteOnEvent{GenericEvent: GenericEvent{Time: 0, Channel: 2}, Note: 60, Velocity: 100})

	for i := 0; i < b.N; i++ {
		EncodeEventBuffer(eventBuffer)
	}
}

func BenchmarkDecodeByteBuffer(b *testing.B) {
	byteBuffer := []byte{3, 0, 0, 0, 0, byte(Note) << 4, 60, 100, 0}
	for i := 0; i < b.N; i++ {
		DecodeByteBuffer(byteBuffer)
	}
}
