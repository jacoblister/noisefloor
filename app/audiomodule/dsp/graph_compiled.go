package dsp

import (
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor/processorbuiltin"
	"github.com/jacoblister/noisefloor/pkg/midi"
)

// CompileTarget enumerated type
type CompileTarget int

// CompileTarget implementation
const (
	CompileInterpreted CompileTarget = iota
	CompileGolang
	CompileJavascript
	CompileWasm
	CompileCPP
)

// compiledGraph the graph after compilation
type compiledGraph interface {
	Start(sampleRate int)
	Stop()
	Process(samplesIn [][]float32, midiIn []midi.Event) (samplesOut [][]float32, midiOut []midi.Event)
}

type graphOp struct {
	processor    Processor
	connectorIn  []*Connector
	connectorOut [][]*Connector
}

type graphExecutor struct {
	midiInput  *processorbuiltin.MIDIInput // 'special' MIDI input processor
	inputTerm  *processorbuiltin.Terminal  // 'special' Audio input terminal
	outputTerm *processorbuiltin.Terminal  // 'special' Audio output terminal
	ops        []graphOp                   // operations to perform, in order
}

func compileGraphExecutor(graph Graph) graphExecutor {
	// init graphExecutor and provide default 'special' processors
	graphExecutor := graphExecutor{}
	graphExecutor.midiInput = &processorbuiltin.MIDIInput{}
	graphExecutor.outputTerm = &processorbuiltin.Terminal{}

	for i := 0; i < len(graph.ProcessorList); i++ {
		// check for 'special' processors
		midiInput, ok := graph.ProcessorList[i].Processor.(*processorbuiltin.MIDIInput)
		if ok {
			graphExecutor.midiInput = midiInput
		}
		terminal, ok := graph.ProcessorList[i].Processor.(*processorbuiltin.Terminal)
		if ok {
			graphExecutor.outputTerm = terminal
		}

		graphExecutor.ops = append(graphExecutor.ops,
			graphOp{
				graph.ProcessorList[i].Processor,
				graph.inputConnectorsForProcessor(graph.ProcessorList[i].Processor),
				graph.outputConnectorsForProcessor(graph.ProcessorList[i].Processor),
			},
		)
	}

	return graphExecutor
}

// compileProcessorGraph compiles a graph, and returns a function to run it
func compileProcessorGraph(graph Graph, target CompileTarget) compiledGraph {
	graphExecutor := compileGraphExecutor(graph)

	switch target {
	case CompileInterpreted:
		compiledGraph := interpretedEngine{graphExecutor: graphExecutor}
		return &compiledGraph

	}
	panic("unsupported target")
}
