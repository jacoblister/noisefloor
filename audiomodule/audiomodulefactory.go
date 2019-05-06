package audiomodule

import "github.com/jacoblister/noisefloor/audiomodule/synth"

//MakeComponent generates a new compoent by the given name
func MakeComponent(name string) AudioProcessor {
	switch name {
	case "SynthEngine":
		return &synth.Engine{}
	}

	return nil
}
