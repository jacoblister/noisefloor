package app

import (
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp"
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/dspUI"
	"github.com/jacoblister/noisefloor/app/audiomodule/onscreenkeyboard"
	"github.com/jacoblister/noisefloor/app/audiomodule/onscreenkeyboard/onscreenkeyboardUI"
	"github.com/jacoblister/noisefloor/app/vdomcomp"
	"github.com/jacoblister/noisefloor/pkg/midi"
	"github.com/jacoblister/noisefloor/pkg/vdom"
)

type modules struct {
	keyboard  onscreenkeyboard.Keyboard
	dspEngine dsp.Engine
	state     struct {
		vDividerPos      int
		vDividerMoving   bool
		dspUIEngineState dspUI.EngineState
	}
}

// Start begin the main application audio processing
func (c *modules) Start(sampleRate int) {
	c.keyboard.Start(sampleRate)
	c.dspEngine.Start(sampleRate)
}

// Stop closes the main application audio processing
func (c *modules) Stop() {
	c.keyboard.Stop()
	c.dspEngine.Stop()
}

// Process process a block of audio/midi
func (c *modules) Process(samplesIn [][]float32, midiIn []midi.Event) (samplesOut [][]float32, midiOut []midi.Event) {
	samples, midi := samplesIn, midiIn

	samples, midi = c.keyboard.Process(samples, midi)
	samples, midi = c.dspEngine.Process(samples, midi)

	return samples, midi
}

func (c *modules) Init() {
	if c.state.vDividerPos == 0 {
		c.state.vDividerPos = 500
	}
}

// Render returns the main view
func (c *modules) Render() vdom.Element {
	// elem := vdom.MakeElement("svg",
	// 	"id", "root",
	// 	"xmlns", "http://www.w3.org/2000/svg",
	// 	"style", "width:100%;height:100%;position:fixed;top:0;left:0;bottom:0;right:0;",
	// 	onscreenkeyboardUI.MakeKeyboard(&c.keyboard),
	// )

	// dspUI := dspUI.MakeEngine(&c.dspEngine, &c.state.dspUIEngineState)
	// vSplit := vdomcomp.MakeLayoutVSplit(1024, 768, c.state.vDividerPos, &c.state.vDividerMoving,
	// 	&vdomcomp.Text{Text: "menus"}, dspUI,
	// 	func(pos int) {
	// 		if pos > 100 {
	// 			c.state.vDividerPos = pos
	// 		}
	// 	},
	// )

	dspUI := dspUI.MakeEngine(&c.dspEngine, &c.state.dspUIEngineState)
	vSplit := vdomcomp.MakeLayoutHSplit(1024, 768, c.state.vDividerPos, &c.state.vDividerMoving,
		dspUI, onscreenkeyboardUI.MakeKeyboard(&c.keyboard),
		func(pos int) {
			if pos > 100 {
				c.state.vDividerPos = pos
			}
		},
	)

	elem := vdom.MakeElement("svg",
		"id", "root",
		"xmlns", "http://www.w3.org/2000/svg",
		"style", "width:100%;height:100%;position:fixed;top:0;left:0;bottom:0;right:0;",
		vSplit,
	)

	return elem
}
