package dsp

import (
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"
	"github.com/jacoblister/noisefloor/pkg/midi"
)

type interpretedEngine struct {
	graphExecutor graphExecutor
	vars          []float32
	osc           processor.Oscillator
}

func (g *interpretedEngine) Start(sampleRate int) {
	g.osc.Start(sampleRate)
	for i := 0; i < len(g.graphExecutor.ops); i++ {
		g.graphExecutor.ops[i].processor.Start(sampleRate)
	}
	g.vars = make([]float32, g.graphExecutor.varCount, g.graphExecutor.varCount)
	g.vars[1] = 550
	g.graphExecutor.midiInput.SetMono()
}

func (g *interpretedEngine) Process(samplesIn [][]float32, midiIn []midi.Event) (samplesOut [][]float32, midiOut []midi.Event) {
	g.graphExecutor.midiInput.ProcessMIDI(midiIn)
	g.graphExecutor.outputTerm.SetSamples(samplesIn)

	inArgs := make([]float32, 0, 8)
	var length = len(samplesIn[0])
	for i := 0; i < length; i++ {
		// samplesIn[0][i] = g.osc.Process(440)

		for j := 0; j < len(g.graphExecutor.ops); j++ {
			op := g.graphExecutor.ops[j]
			inArgs := inArgs[:len(op.inVars)]
			for k := 0; k < len(op.inVars); k++ {
				inArgs[k] = g.vars[op.inVars[k]]
			}
			outArgs := g.graphExecutor.ops[j].processor.ProcessArray(inArgs)
			for k := 0; k < len(op.outVars); k++ {
				g.vars[op.outVars[k]] = outArgs[k]
			}
		}
		g.graphExecutor.midiInput.NextSample()
	}

	return samplesIn, midiIn
}

func (g *interpretedEngine) Stop() {
}
