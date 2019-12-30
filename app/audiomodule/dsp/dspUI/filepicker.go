package dspUI

import (
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp"
	"github.com/jacoblister/noisefloor/app/vdomcomp"
	"github.com/jacoblister/noisefloor/pkg/vdom"
)

//FilePicker is the dsp engine file selector
type FilePicker struct {
	Engine *dsp.Engine
	width  int
	height int
	state  *FilePickerState
}

// FilePickerState is the file picker stateful store
type FilePickerState struct {
	initialized  bool
	item         []string
	selectedItem string
}

//MakeFilePicker create an new Engine File picker UI componenet
func MakeFilePicker(engine *dsp.Engine, width int, height int, filePickerState *FilePickerState) *FilePicker {
	filePicker := FilePicker{Engine: engine, width: width, height: height, state: filePickerState}

	return &filePicker
}

//Render displays the dsp file picker.
func (f *FilePicker) Render() vdom.Element {
	if !f.state.initialized {
		f.state.item = f.Engine.Files()
		f.state.selectedItem = f.Engine.Filename()
		f.state.initialized = true
	}

	list := vdomcomp.MakePickList(0, 0, f.width-1, f.height, f.state.item, f.state.selectedItem, func(item string) {
		f.Engine.Load(item)
		f.state.selectedItem = item
	})

	return vdom.MakeElement("g", &list)
}
