package dsp

import (
	"github.com/jacoblister/noisefloor/pkg/midi"
)

type interpretedEngine struct {
	graphExecutor graphExecutor
}

const frameLength = 4096

func (g *interpretedEngine) Start(sampleRate int) {
	for i := 0; i < len(g.graphExecutor.ops); i++ {
		op := &g.graphExecutor.ops[i]
		connectedMask := 0
		for j := 0; j < len(op.connectorIn); j++ {
			if op.connectorIn[j].Samples() == nil {
				op.connectorIn[j].SetSamples(make([]float32, frameLength, frameLength))
			}
			if op.connectorIn[j].FromProcessor != nil {
				connectedMask |= (1 << uint(j))
			}
		}
		op.processor.Start(sampleRate, connectedMask)
	}
	g.graphExecutor.midiInput.Start(sampleRate, 0)
	g.graphExecutor.midiInput.SetMono()
}

func (g *interpretedEngine) Process(samplesIn [][]float32, midiIn []midi.Event) (samplesOut [][]float32, midiOut []midi.Event) {
	g.graphExecutor.midiInput.ProcessMIDI(midiIn)
	g.graphExecutor.outputTerm.SetSamples(samplesIn)

	var length = len(samplesIn[0])
	inArgs := make([][]float32, 0, 8)
	for j := 0; j < len(g.graphExecutor.ops); j++ {
		op := g.graphExecutor.ops[j]
		inArgs := inArgs[:len(op.connectorIn)]
		for k := 0; k < len(op.connectorIn); k++ {
			inArgs[k] = op.connectorIn[k].Samples()
		}
		outArgs := g.graphExecutor.ops[j].processor.ProcessSamples(inArgs, length)
		for k := 0; k < len(op.connectorOut); k++ {
			for l := 0; l < len(op.connectorOut[k]); l++ {
				op.connectorOut[k][l].SetSamples(outArgs[k])
				op.connectorOut[k][l].Value = outArgs[k][0]
			}
		}
	}

	return samplesIn, midiIn
}

func (g *interpretedEngine) Stop() {
}
