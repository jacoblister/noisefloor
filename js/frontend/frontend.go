package frontend

import (
	"github.com/bep/gr"
	. "github.com/jacoblister/noisefloor/common"
	"github.com/jacoblister/noisefloor/js/frontend/component"
)

var keyboard *component.Keyboard

// GetMIDIEvents returns the currently pending MIDI events
func GetMIDIEvents() []MidiEvent {
	return keyboard.GetMIDIEvents()
}

func init() {
	keyboard = new(component.Keyboard)
	react := gr.New(keyboard)

	react.Render("react", gr.Props{})
}
