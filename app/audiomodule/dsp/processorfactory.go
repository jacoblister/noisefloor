package dsp

import (
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor/processorbuiltin"
)

// ListProcessors returns a list of available processors
func ListProcessors() []string {
	return []string{"MIDIInput", "Terminal", "Constant", "Envelope", "Gain", "Multiply", "Oscillator", "Scope", "Splitter", "Sum"}
}

//MakeProcessor generates a new processor by the given processor name
func MakeProcessor(name string) Processor {
	var proc Processor

	switch name {
	case "MIDIInput":
		proc = &processorbuiltin.MIDIInput{}
	case "Terminal":
		// TODO - consider alternative terminal parameters
		terminal := &processorbuiltin.Terminal{}
		terminal.SetParameters(true, 2)
		proc = terminal
	case "Constant":
		proc = &processor.Constant{}
	case "Envelope":
		proc = &processor.Envelope{}
	case "Gain":
		proc = &processor.Gain{}
	case "Multiply":
		proc = &processor.Multiply{}
	case "Oscillator":
		proc = &processor.Oscillator{}
	case "Scope":
		proc = &processor.Scope{}
	case "Splitter":
		proc = &processor.Splitter{}
	case "Sum":
		proc = &processor.Sum{}
	}

	setProcessorDefaults(proc)

	return proc
}
