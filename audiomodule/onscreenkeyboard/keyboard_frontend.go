package onscreenkeyboard

import (
	"strconv"

	"github.com/jacoblister/noisefloor/midi"
	"github.com/jacoblister/noisefloor/vdom"
)

func (k *Keyboard) noteEventFromUI(keyNumber int, keyDown bool) {
	if k.keydown[keyNumber] == keyDown {
		// return early if key already is same state
		return
	}

	k.keydown[keyNumber] = keyDown

	var midiEvent midi.Event
	if keyDown {
		midiEvent = midi.NoteOnEvent{GenericEvent: midi.GenericEvent{Time: 0, Channel: 1},
			Note: keyNumber, Velocity: velocityMax}
	} else {
		midiEvent = midi.NoteOffEvent{GenericEvent: midi.GenericEvent{Time: 0, Channel: 1},
			Note: keyNumber}
	}
	k.MidiEvents = append(k.MidiEvents, midiEvent)

	vdom.UpdateComponent(k)
}

func (k *Keyboard) noteEventFromProcess(keyNumber int, keyDown bool) {
	k.keydown[keyNumber] = keyDown

	vdom.UpdateComponentBackground(k)
}

func (k *Keyboard) renderKey(keyNumber int, isBlack bool, xPosition int, depressed bool) vdom.Element {
	stroke := "black"
	fill := "white"

	keyType := "key-white"
	width := 40
	height := 160
	if isBlack {
		fill = "black"
		keyType = "key-black"
		xPosition += 28
		width = 26
		height = 120
	}

	var depressedElem vdom.Attr
	if depressed {
		depressedElem = vdom.Attr{Name: "style", Value: "depressed"}
		fill = "lightcyan"
	}

	key := vdom.MakeElement("rect",
		"id", "key-"+strconv.Itoa(keyNumber),
		"class", keyType,
		depressedElem,
		"x", xPosition,
		"y", 10,
		"width", width,
		"height", height,
		"stroke", stroke,
		"fill", fill,
		vdom.MakeEventHandler(vdom.MouseDown, func(element *vdom.Element, event *vdom.Event) {
			k.noteEventFromUI(keyNumber, true)
		}),
		vdom.MakeEventHandler(vdom.MouseUp, func(element *vdom.Element, event *vdom.Event) {
			k.noteEventFromUI(keyNumber, false)
		}),
		// vdom.MakeEventHandler(vdom.MouseEnter, func(element *vdom.Element, event *vdom.Event) {
		// 	buttons, _ := strconv.Atoi(event.Data)
		// 	if buttons > 0 {
		// 		k.noteEvent(keyNumber, true)
		// 	}
		// }),
		// vdom.MakeEventHandler(vdom.MouseLeave, func(element *vdom.Element, event *vdom.Event) {
		// 	k.noteEvent(keyNumber, false)
		// }),
		// vdom.MakeEventHandler(vdom.TouchStart, func(element *vdom.Element, event *vdom.Event) {
		// 	k.noteEvent(keyNumber, true)
		// }),
		// vdom.MakeEventHandler(vdom.TouchEnd, func(element *vdom.Element, event *vdom.Event) {
		// 	k.noteEvent(keyNumber, false)
		// }),
	)
	return key
}

func isBlackKey(n int) bool {
	return n == 1 || n == 3 || n == 6 || n == 8 || n == 10
}

func (k *Keyboard) renderOctave(parent *vdom.Element, keyStart int, xStart int) {
	for noteType := 0; noteType < 2; noteType++ {
		xPos := xStart
		for keyNumber := 0; keyNumber < 12; keyNumber++ {
			if keyNumber > 0 && !isBlackKey(keyNumber) {
				xPos += 40
			}
			if noteType == 0 && !isBlackKey(keyNumber) {
				key := k.renderKey(keyNumber+keyStart, false, xPos, k.keydown[keyNumber+keyStart])
				parent.AppendChild(key)
			}
			if noteType == 1 && isBlackKey(keyNumber) {
				key := k.renderKey(keyNumber+keyStart, true, xPos, k.keydown[keyNumber+keyStart])
				parent.AppendChild(key)
			}
		}
	}
}

// Render displays the keyboard.
func (k *Keyboard) Render() vdom.Element {
	println("render keyboard")
	elem := vdom.MakeElement("g")
	for octave := 0; octave < 3; octave++ {
		k.renderOctave(&elem, 48+octave*12, (40*7*octave)+1)
	}

	return elem
}

// // ComponentDidMount registers DOM event handler for physical keyboard actions
// func (k *KeyboardFrontend) ComponentDidMount() {
// 	var keyMap = map[string]int{
// 		"a": 60, "w": 61, "s": 62, "e": 63, "d": 64,
// 		"f": 65, "t": 66, "g": 67, "y": 68, "h": 69, "u": 70, "j": 71, "k": 72,
// 	}
//
// 	doc := js.Global.Get("document")
// 	doc.Call("addEventListener", "keydown", func(event *js.Object) {
// 		midiKey := keyMap[event.Get("key").Call("toLowerCase").String()]
// 		if midiKey != 0 {
// 			k.noteEvent(midiKey, true)
// 		}
// 	})
// 	doc.Call("addEventListener", "keyup", func(event *js.Object) {
// 		midiKey := keyMap[event.Get("key").Call("toLowerCase").String()]
// 		if midiKey != 0 {
// 			k.noteEvent(midiKey, false)
// 		}
// 	})
// }
