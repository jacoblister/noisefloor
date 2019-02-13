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
func getTypeFromData(data []byte) int {
	return int(data[0]&0x70) >> 4
}

//getChannelFromData gets the MIDI event channel for an event
func getChannelFromData(data []byte) int {
	return int(data[0]&0x0F) + 1
}

//Event interface specifies a generic MIDI event
type Event interface {
	AsMidiEventData() EventData
}

// MakeMidiEventData constructs a new midi event from time and data bytes
func MakeMidiEventData(time int, data []byte) *EventData {
	return &EventData{time, data}
}

// MakeMidiEvent returns an interface to a decoded MIDI event
func MakeMidiEvent(time int, data []byte) Event {
	return &NoteOnEvent{Time: time, Channel: getChannelFromData(data), Note: int(data[1])}
}

// NoteOnEvent is a MIDI note on event with velocity
type NoteOnEvent struct {
	Time     int
	Channel  int
	Note     int
	Velocity int
}

func (e NoteOnEvent) AsMidiEventData() EventData {
	return EventData{Time: e.Time}
}
