package midi

import (
	"testing"

	"gotest.tools/assert"
)

func TestMakeMidiEvent(t *testing.T) {
	event := MakeMidiEvent(123, []byte{1, 100, 0})

	assert.Equal(t, event.AsMidiEventData().Time, 123)
	noteOnEvent := event.(*NoteOnEvent)
	assert.Equal(t, noteOnEvent.Channel, 2)
	assert.Equal(t, noteOnEvent.Note, 100)
}
