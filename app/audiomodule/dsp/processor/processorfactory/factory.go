package processorfactory

import (
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor/processorbuiltin"
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor/processorbasic"
)

// ListProcessors returns a list of available processors
func ListProcessors() []string {
	return []string {"MIDIInput","Terminal","Add","Constant","Divide","Envelope","Gain","Multiply","Oscillator","Scope","Splitter","Sum"}
}

// MakeProcessor generates a new processor by the given processor name
func MakeProcessor(name string) processor.Processor {
	var proc processor.Processor

	switch name {
	case "MIDIInput":
		proc = &processorbuiltin.MIDIInput{}
	case "Terminal":
		proc = &processorbuiltin.Terminal{}
	case "Add":
		proc = &processorbasic.Add{}
	case "Constant":
		proc = &processorbasic.Constant{}
	case "Divide":
		proc = &processorbasic.Divide{}
	case "Envelope":
		proc = &processorbasic.Envelope{}
	case "Gain":
		proc = &processorbasic.Gain{}
	case "Multiply":
		proc = &processorbasic.Multiply{}
	case "Oscillator":
		proc = &processorbasic.Oscillator{}
	case "Scope":
		proc = &processorbasic.Scope{}
	case "Splitter":
		proc = &processorbasic.Splitter{}
	case "Sum":
		proc = &processorbasic.Sum{}
	}

	processor.SetProcessorDefaults(proc)

	return proc
}
