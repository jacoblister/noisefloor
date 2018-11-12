package frontend

import (
	"github.com/bep/gr"
	. "github.com/jacoblister/noisefloor/common"
	"github.com/jacoblister/noisefloor/js/frontend/component"
)

// GetMIDIEvents returns the currently pending MIDI events
func GetMIDIEvents() []MidiEvent {
	return component.GetMIDIEvents()
}

func init() {
	println("init frontend main")

	keyboard := gr.New(new(component.Keyboard))

	keyboard.Render("react", gr.Props{})
}
