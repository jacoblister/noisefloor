package engine

import (
	. "github.com/jacoblister/noisefloor/common"
)

//Processor interface
type Processor interface {
	Start(sampleRate int)
	Stop()
	Process(args ...AudioFloat) AudioFloat
}
