package main

import (
	. "github.com/jacoblister/noisefloor/common"
	"honnef.co/go/js/dom"
)

var midiEvents []MidiEvent

var keyMap = map[string]int{
	"a": 60, "w": 61, "s": 62, "e": 63, "d": 64,
	"f": 65, "t": 66, "g": 67, "y": 68, "h": 69, "u": 70, "j": 71, "k": 72,
}

func keyDown(event dom.Event) {
	keyEvent := event.(*dom.KeyboardEvent)
	midiKey := keyMap[keyEvent.Key]
	if midiKey != 0 {
		midiEvent := MidiEvent{0, midiKey, 127}
		midiEvents = append(midiEvents, midiEvent)
	}
}

func keyUp(event dom.Event) {
	keyEvent := event.(*dom.KeyboardEvent)
	midiKey := keyMap[keyEvent.Key]
	if midiKey != 0 {
		midiEvent := MidiEvent{0, midiKey, 0}
		midiEvents = append(midiEvents, midiEvent)
	}
}

// GetMIDIEvents returns the currently pending MIDI events
func GetMIDIEvents() []MidiEvent {
	result := midiEvents
	midiEvents = midiEvents[:0]
	return result
}

// RenderFrontend renders the frontend into the top level DOM Document
func RenderFrontend() {
	d := dom.GetWindow().Document()
	// div := d.CreateElement("div").(*dom.HTMLDivElement)
	// div.Style().SetProperty("color", "red", "")
	// div.SetTextContent("Noisefloor Frontend")

	root := d.GetElementByID("root")
	root.SetTextContent("Noisefloor Frontend")

	d.AddEventListener("keydown", true, keyDown)
	d.AddEventListener("keyup", true, keyUp)

}
