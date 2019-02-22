package synth

import "github.com/jacoblister/noisefloor/component/synth/processor"

//MakeProcessor generates a new processor by the given processor name
func MakeProcessor(name string) Processor {
	switch name {
	case "Envelope":
		return &processor.Envelope{}
	case "Gain":
		return &processor.Gain{}
	case "Oscillator":
		return &processor.Oscillator{}
	}

	return nil
}
