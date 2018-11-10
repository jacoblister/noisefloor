package component

import (
	"github.com/gopherjs/gopherjs/js"
	. "github.com/jacoblister/noisefloor/common"

	"github.com/bep/gr"
	"github.com/bep/gr/evt"
	"github.com/bep/gr/svg"
	"github.com/bep/gr/svga"
)

// Keyboard is an onscreen MIDI keyboard
type Keyboard struct {
	*gr.This
	keydown [12]bool
}

var keyMap = map[string]int{
	"a": 60, "w": 61, "s": 62, "e": 63, "d": 64,
	"f": 65, "t": 66, "g": 67, "y": 68, "h": 69, "u": 70, "j": 71, "k": 72,
}

// GetInitialState sets up the keyboard state.
func (k *Keyboard) GetInitialState() gr.State {
	return gr.State{"keydown": k.keydown}
}

func (k *Keyboard) renderKey(keyNumber int, depressed bool) *gr.Element {
	var depressedElem gr.Modifier
	if depressed {
		depressedElem = gr.CSS("depressed")
	}

	key := svg.Rect(
		gr.CSS("key-white"),
		depressedElem,
		svga.X(keyNumber*20),
		svga.Y(10),
		svga.Width(20),
		svga.Height(40),
		evt.MouseDown(func(event *gr.Event) {
			k.keydown[keyNumber] = true
			k.SetState(gr.State{"keydown": k.keydown})
		}),
		evt.MouseUp(func(event *gr.Event) {
			k.keydown[keyNumber] = false
			k.SetState(gr.State{"keydown": k.keydown})
		}),
	)
	return key
}

// Render displays the keyboard.
func (k *Keyboard) Render() gr.Component {
	elem := svg.G()
	for i := 0; i < 12; i++ {
		key := k.renderKey(i, k.keydown[i])
		key.Modify(elem)
	}

	return elem
}

// ComponentDidMount registers DOM event handler for physical keyboard actions
func (k *Keyboard) ComponentDidMount() {
	doc := js.Global.Get("document")
	doc.Call("addEventListener", "keydown", func(event *js.Object) {
		midiKey := keyMap[event.Get("key").String()]
		if midiKey != 0 {
			println("key pressed", midiKey)
		}
	})
}

// GetMIDIEvents returns the currently pending MIDI events
func GetMIDIEvents() []MidiEvent {
	// result := midiEvents
	// midiEvents = midiEvents[:0]
	var midiEvents []MidiEvent
	return midiEvents
}
