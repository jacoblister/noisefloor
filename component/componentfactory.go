package component

import (
	"github.com/jacoblister/noisefloor/engine"
)

//MakeComponent generates a new compoent by the given name
func MakeComponent(name string) AudioProcessor {
	switch name {
	case "Engine":
		return &engine.Engine{}
	}

	return nil
}
