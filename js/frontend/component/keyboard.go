package component

import (
	"github.com/gopherjs/gopherjs/js"
	. "github.com/jacoblister/noisefloor/common"

	"github.com/bep/gr"
	"github.com/bep/gr/evt"
	"github.com/bep/gr/svg"
	"github.com/bep/gr/svga"
)

const keyMax = 127
const velocityMax = 127

// Keyboard is an onscreen MIDI keyboard
type Keyboard struct {
	*gr.This
	keydown [keyMax]bool
}

var midiEvents []MidiEvent

// GetMIDIEvents returns the currently pending MIDI events
func GetMIDIEvents() []MidiEvent {
	return midiEvents
}

func (k *Keyboard) noteEvent(keyNumber int, keyDown bool) {
	k.keydown[keyNumber] = keyDown
	k.SetState(gr.State{"keydown": k.keydown})

	velocity := 0
	if keyDown {
		velocity = velocityMax
	}
	midiEvent := MidiEvent{0, 0, keyNumber, velocity}

	midiEvents = append(midiEvents, midiEvent)
}

// GetInitialState sets up the keyboard state.
func (k *Keyboard) GetInitialState() gr.State {
	return gr.State{"keydown": k.keydown}
}

func (k *Keyboard) renderKey(keyNumber int, isBlack bool, xPosition int, depressed bool) *gr.Element {
	var depressedElem gr.Modifier
	if depressed {
		depressedElem = gr.CSS("depressed")
	}

	keyType := "key-white"
	width := 40
	height := 160
	if isBlack {
		keyType = "key-black"
		xPosition += 28
		width = 26
		height = 120
	}

	key := svg.Rect(
		gr.CSS(keyType),
		depressedElem,
		svga.X(xPosition),
		svga.Y(10),
		svga.Width(width),
		svga.Height(height),
		evt.MouseDown(func(event *gr.Event) {
			k.noteEvent(keyNumber, true)
		}),
		evt.MouseUp(func(event *gr.Event) {
			k.noteEvent(keyNumber, false)
		}),
		evt.MouseOut(func(event *gr.Event) {
			k.noteEvent(keyNumber, false)
		}),
		evt.MouseEnter(func(event *gr.Event) {
			if event.Get("buttons").Int() != 0 {
				k.noteEvent(keyNumber, true)
			}
		}),
		evt.TouchStart(func(event *gr.Event) {
			k.noteEvent(keyNumber, true)
		}),
		evt.TouchEnd(func(event *gr.Event) {
			k.noteEvent(keyNumber, false)
		}),
	)
	return key
}

func isBlackKey(n int) bool {
	return n == 1 || n == 3 || n == 6 || n == 8 || n == 10
}

func (k *Keyboard) renderOctave(elem *gr.Element, keyStart int, xStart int) *gr.Element {
	for noteType := 0; noteType < 2; noteType++ {
		xPos := xStart
		for keyNumber := 0; keyNumber < 12; keyNumber++ {
			if keyNumber > 0 && !isBlackKey(keyNumber) {
				xPos += 40
			}
			if noteType == 0 && !isBlackKey(keyNumber) {
				key := k.renderKey(keyNumber+keyStart, false, xPos, k.keydown[keyNumber+keyStart])
				key.Modify(elem)
			}
			if noteType == 1 && isBlackKey(keyNumber) {
				key := k.renderKey(keyNumber+keyStart, true, xPos, k.keydown[keyNumber+keyStart])
				key.Modify(elem)
			}
		}
	}

	return elem
}

// Render displays the keyboard.
func (k *Keyboard) Render() gr.Component {
	elem := svg.G()
	// k.renderOctave(elem, 60, 0)
	for octave := 0; octave < 3; octave++ {
		k.renderOctave(elem, 48+octave*12, (40*7*octave)+1)
	}

	return elem
}

var keyMap = map[string]int{
	"a": 60, "w": 61, "s": 62, "e": 63, "d": 64,
	"f": 65, "t": 66, "g": 67, "y": 68, "h": 69, "u": 70, "j": 71, "k": 72,
}

// ComponentDidMount registers DOM event handler for physical keyboard actions
func (k *Keyboard) ComponentDidMount() {
	doc := js.Global.Get("document")
	doc.Call("addEventListener", "keydown", func(event *js.Object) {
		midiKey := keyMap[event.Get("key").Call("toLowerCase").String()]
		if midiKey != 0 {
			k.noteEvent(midiKey, true)
		}
	})
	doc.Call("addEventListener", "keyup", func(event *js.Object) {
		midiKey := keyMap[event.Get("key").Call("toLowerCase").String()]
		if midiKey != 0 {
			k.noteEvent(midiKey, false)
		}
	})
}
