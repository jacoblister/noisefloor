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
		op.processor.Start(sampleRate)
		for j := 0; j < len(op.connectorIn); j++ {
			if op.connectorIn[j].samples == nil {
				op.connectorIn[j].samples = make([]float32, frameLength, frameLength)
			}
		}
	}

	g.graphExecutor.midiInput.Start(sampleRate)
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
			inArgs[k] = op.connectorIn[k].samples
		}
		outArgs := g.graphExecutor.ops[j].processor.ProcessSamples(inArgs, length)
		for k := 0; k < len(op.connectorOut); k++ {
			for l := 0; l < len(op.connectorOut[k]); l++ {
				op.connectorOut[k][l].samples = outArgs[k]
				op.connectorOut[k][l].Value = outArgs[k][0]
			}
		}
	}

	return samplesIn, midiIn
}

func (g *interpretedEngine) Stop() {
}
