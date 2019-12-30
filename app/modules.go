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
		vDividerPos          int
		vDividerMoving       bool
		hDividerPos          int
		hDividerMoving       bool
		dspUIEngineState     dspUI.EngineState
		dspUIFilePickerState dspUI.FilePickerState
	}
}

// Start begin the main application audio processing
func (m *modules) Start(sampleRate int) {
	m.keyboard.Start(sampleRate)
	m.dspEngine.Start(sampleRate)
}

// Stop closes the main application audio processing
func (m *modules) Stop() {
	m.keyboard.Stop()
	m.dspEngine.Stop()
}

// Process process a block of audio/midi
func (m *modules) Process(samplesIn [][]float32, midiIn []midi.Event) (samplesOut [][]float32, midiOut []midi.Event) {
	samples, midi := samplesIn, midiIn

	samples, midi = m.keyboard.Process(samples, midi)
	samples, midi = m.dspEngine.Process(samples, midi)

	return samples, midi
}

func (m *modules) Init() {
	m.state.hDividerPos = 640
	m.state.vDividerPos = 200
}

// Render returns the main view
func (m *modules) Render() vdom.Element {
	engineUI := dspUI.MakeEngine(&m.dspEngine, 1200-m.state.vDividerPos-4, m.state.hDividerPos, &m.state.dspUIEngineState)
	filePickerUI := dspUI.MakeFilePicker(&m.dspEngine, m.state.vDividerPos, m.state.hDividerPos, &m.state.dspUIFilePickerState)
	vSplit := vdomcomp.MakeLayoutVSplit(1200, m.state.hDividerPos, m.state.vDividerPos, 4, &m.state.vDividerMoving,
		filePickerUI,
		engineUI,
		func(pos int) {
			if pos > 100 {
				m.state.vDividerPos = pos
			}
		},
	)

	hSplit := vdomcomp.MakeLayoutHSplit(1200, 768, m.state.hDividerPos, 4, &m.state.hDividerMoving,
		vSplit, onscreenkeyboardUI.MakeKeyboard(&m.keyboard),
		func(pos int) {
			if pos > 100 {
				m.state.hDividerPos = pos
			}
		},
	)

	elem := vdom.MakeElement("svg",
		"id", "root",
		"xmlns", "http://www.w3.org/2000/svg",
		"style", "width:100%;height:100%;position:fixed;top:0;left:0;bottom:0;right:0;",
		hSplit,
	)

	return elem
}
