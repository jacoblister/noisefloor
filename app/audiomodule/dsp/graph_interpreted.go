package dsp

import (
	"github.com/jacoblister/noisefloor/pkg/midi"
)

type interpretedEngine struct {
	graphExecutor graphExecutor
}

func (g *interpretedEngine) Start(sampleRate int) {
	for i := 0; i < len(g.graphExecutor.ops); i++ {
		g.graphExecutor.ops[i].processor.Start(sampleRate)
	}

	g.graphExecutor.midiInput.Start(sampleRate)
	g.graphExecutor.midiInput.SetMono()
}

func (g *interpretedEngine) Process(samplesIn [][]float32, midiIn []midi.Event) (samplesOut [][]float32, midiOut []midi.Event) {
	g.graphExecutor.midiInput.ProcessMIDI(midiIn)
	g.graphExecutor.outputTerm.SetSamples(samplesIn)

	inArgs := make([]float32, 0, 8)
	var length = len(samplesIn[0])
	for i := 0; i < length; i++ {
		for j := 0; j < len(g.graphExecutor.ops); j++ {
			op := g.graphExecutor.ops[j]
			inArgs := inArgs[:len(op.connectorIn)]
			for k := 0; k < len(op.connectorIn); k++ {
				inArgs[k] = op.connectorIn[k].Value
			}
			outArgs := g.graphExecutor.ops[j].processor.ProcessArray(inArgs)
			for k := 0; k < len(op.connectorOut); k++ {
				for l := 0; l < len(op.connectorOut[k]); l++ {
					op.connectorOut[k][l].Value = outArgs[k]
				}
			}
		}
		if g.graphExecutor.midiInput != nil {
			g.graphExecutor.midiInput.NextSample()
		}
	}

	return samplesIn, midiIn
}

func (g *interpretedEngine) Stop() {
}
