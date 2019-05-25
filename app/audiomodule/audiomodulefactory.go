package audiomodule

import "github.com/jacoblister/noisefloor/app/audiomodule/dsp"

//MakeComponent generates a new compoent by the given name
func MakeComponent(name string) AudioProcessor {
	switch name {
	case "DSPEngine":
		return &dsp.Engine{}
	}

	return nil
}
