package dsp

import (
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor/processorbuiltin"
	"github.com/jacoblister/noisefloor/pkg/midi"
)

type interpretedEngine struct {
	midiInput   processorbuiltin.MIDIInput
	processor   []Processor
	inputNodes  []int
	outputNodes []int
	vars        []float32
	ops         []graphOp
}

func (g *interpretedEngine) Start(sampleRate int) {
	g.processor = []Processor{&processor.Oscillator{}, &processor.Envelope{}, &processor.Gain{}}
	for i := 0; i < len(g.processor); i++ {
		g.processor[i].Start(sampleRate)
	}

	g.vars = make([]float32, 8, 8)
	g.inputNodes = []int{0, 1, 2}
	g.outputNodes = []int{7, 7}
	g.ops = append(g.ops, graphOp{g.processor[0], []int{0}, []int{3}})
	g.ops = append(g.ops, graphOp{g.processor[1], []int{1, 2}, []int{4}})
	g.ops = append(g.ops, graphOp{g.processor[2], []int{5, 6}, []int{7}})
}

func (g *interpretedEngine) Process(samplesIn [][]float32, midiIn []midi.Event) (samplesOut [][]float32, midiOut []midi.Event) {
	inArgs := make([]float32, 0, 8)

	g.midiInput.ProcessMIDI(midiIn)
	// update mapped parameters

	var length = len(samplesIn[0])
	for i := 0; i < length; i++ {
		frequency, gate, trigger, _ := g.midiInput.Process(0)
		g.vars[g.inputNodes[0]] = frequency
		g.vars[g.inputNodes[1]] = gate
		g.vars[g.inputNodes[2]] = trigger

		for j := 0; j < len(g.ops); j++ {
			op := g.ops[j]
			inArgs := inArgs[:len(op.inArgs)]
			for k := 0; k < len(op.inArgs); k++ {
				inArgs[k] = g.vars[op.inArgs[k]]
			}
			outArgs := g.ops[j].processor.ProcessArray(inArgs)
			for k := 0; k < len(op.outArgs); k++ {
				g.vars[op.outArgs[k]] = outArgs[k]
			}
		}
		samplesIn[0][i] = g.vars[g.outputNodes[0]]
		samplesIn[1][i] = g.vars[g.outputNodes[1]]
	}
	return samplesIn, midiIn
}

func (g *interpretedEngine) Stop() {
}
