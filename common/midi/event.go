package midi

// EventData struct with time and data
type EventData struct {
	Time int
	Data []byte
}

// EventType is the basic MIDI event type
type EventType int

// Basic MIDI event types
const (
	NoteOff    EventType = 8
	Note       EventType = 9
	AfterTouch EventType = 10
	Control    EventType = 11
	Program    EventType = 12
	Pressure   EventType = 13
	PitchWheel EventType = 14
	System     EventType = 15
)

//getTypeFromData gets the basic MIDI event Type for an event
func getTypeFromData(data []byte) EventType {
	return EventType(data[0] >> 4)
}

//getChannelFromData gets the MIDI event channel for an event
func getChannelFromData(data []byte) int {
	return int(data[0]&0x0F) + 1
}

//statusByte returns a MIDI status by eventType and Channel
func statusByte(eventType EventType, channel int) byte {
	return byte(eventType)<<4 | byte(channel-1)&0x0F
}

//Event interface specifies a generic MIDI event
type Event interface {
	MidiEventData() EventData
	GenericEvent() genericEvent
}

// MakeMidiEventData constructs a new midi event from time and data bytes
func MakeMidiEventData(time int, data []byte) *EventData {
	return &EventData{time, data}
}

// MakeMidiEvent returns an interface to a decoded MIDI event
func MakeMidiEvent(time int, data []byte) Event {
	switch getTypeFromData(data) {
	case Note:
		// return &NoteOnEvent{Time: time, Channel: getChannelFromData(data), Note: int(data[1]), Velocity: int(data[2])}
		// return &NoteOnEvent{GenericEvent{Time: time, Channel: getChannelFromData(data)}, Note: int(data[1]), Velocity: int(data[2])}
		return &NoteOnEvent{genericEvent: genericEvent{Time: time, Channel: getChannelFromData(data)},
			Note: int(data[1]), Velocity: int(data[2])}
	}
	panic("Could not make midi event from data")
}

type genericEvent struct {
	Time    int
	Channel int
}

// NoteOnEvent is a MIDI note on event with velocity
type NoteOnEvent struct {
	genericEvent
	Note     int
	Velocity int
}

// GenericEvent returns the genericEvent for the NoteOnEvent type
func (e NoteOnEvent) GenericEvent() genericEvent { return e.genericEvent }

// MidiEventData returns MidiEventData for the NoteOnEvent type
func (e NoteOnEvent) MidiEventData() EventData {
	return EventData{Time: e.Time, Data: []byte{statusByte(Note, e.Channel), byte(e.Note), byte(e.Velocity)}}
}
