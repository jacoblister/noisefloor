package dsp

import (
	"github.com/jacoblister/noisefloor/pkg/midi"
)

type interpretedEngine struct {
	graphExecutor graphExecutor
	inArgs        [][]float32
	outArgs       [][]float32
}

const maxFrameLength = 4096
const maxArgs = 8

func (g *interpretedEngine) Start(sampleRate int) {
	g.inArgs = make([][]float32, maxArgs)
	for i := 0; i < 8; i++ {
		g.inArgs[i] = make([]float32, maxFrameLength)
	}

	g.outArgs = make([][]float32, maxArgs)
	for i := 0; i < 8; i++ {
		g.outArgs[i] = make([]float32, maxFrameLength)
	}

	for i := 0; i < len(g.graphExecutor.connectors); i++ {
		g.graphExecutor.connectors[i].SetSamples(make([]float32, maxFrameLength))
	}

	for i := 0; i < len(g.graphExecutor.ops); i++ {
		op := &g.graphExecutor.ops[i]
		connectedMask := 0
		for j := 0; j < len(op.connectorIn); j++ {
			if op.connectorIn[j].Samples() == nil {
				op.connectorIn[j].SetSamples(make([]float32, maxFrameLength))
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

	for j := 0; j < len(g.graphExecutor.ops); j++ {
		op := g.graphExecutor.ops[j]

		g.inArgs = g.inArgs[:len(op.connectorIn)]
		for k := 0; k < len(op.connectorIn); k++ {
			g.inArgs[k] = op.connectorIn[k].Samples()
		}
		g.outArgs = g.outArgs[:len(op.connectorOut)]
		for k := 0; k < len(op.connectorOut); k++ {
			for l := 0; l < len(op.connectorOut[k]); l++ {
				g.outArgs[k] = op.connectorOut[k][l].Samples()
				op.connectorOut[k][l].Value = g.outArgs[k][0]
			}
		}

		g.graphExecutor.ops[j].processor.ProcessSamples(g.inArgs, g.outArgs, length)
	}

	return samplesIn, midiIn
}

func (g *interpretedEngine) Stop() {
}
