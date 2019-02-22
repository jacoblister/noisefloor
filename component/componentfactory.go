package component

import "github.com/jacoblister/noisefloor/component/synth"

//MakeComponent generates a new compoent by the given name
func MakeComponent(name string) AudioProcessor {
	switch name {
	case "SynthEngine":
		return &synth.Engine{}
	}

	return nil
}
