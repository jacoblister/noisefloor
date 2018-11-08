package component

import (
	. "github.com/jacoblister/noisefloor/common"

	"github.com/bep/gr"
	"github.com/bep/gr/evt"
	"github.com/bep/gr/svg"
	"github.com/bep/gr/svga"
)

type Keyboard struct {
	*gr.This
}

// Implements the StateInitializer interface.
func (k Keyboard) GetInitialState() gr.State {
	return gr.State{"keydown": false}
}

// Implements the Renderer interface.
func (k Keyboard) Render() gr.Component {
	var depressed gr.Modifier
	if k.State().Bool("keydown") {
		depressed = gr.CSS("depressed")
	}

	elem := svg.Rect(
		gr.CSS("key-white"),
		depressed,
		svga.X(10),
		svga.Y(10),
		svga.Width(100),
		svga.Height(40),
		evt.MouseDown(func(event *gr.Event) {
			k.SetState(gr.State{"keydown": true})
		}),
		evt.MouseUp(func(event *gr.Event) {
			k.SetState(gr.State{"keydown": false})
		}),
	)
	return elem
}

// Implements the ShouldComponentUpdate interface.
func (k Keyboard) ShouldComponentUpdate(next gr.Cops) bool {
	return k.State().HasChanged(next.State, "keydown")
}

// GetMIDIEvents returns the currently pending MIDI events
func GetMIDIEvents() []MidiEvent {
	// result := midiEvents
	// midiEvents = midiEvents[:0]
	var midiEvents []MidiEvent
	return midiEvents
}
