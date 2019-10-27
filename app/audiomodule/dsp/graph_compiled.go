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
	processor Processor
	inVars    []int // index of input variables
	outVars   []int // index of output variables
}

type graphExecutor struct {
	midiInput  *processorbuiltin.MIDIInput // 'special' MIDI input processor
	inputTerm  *processorbuiltin.Terminal  // 'special' Audio input terminal
	outputTerm *processorbuiltin.Terminal  // 'special' Audio output terminal
	varCount   int                         // number of process variables in graph
	ops        []graphOp                   // operations to perform, in order
}

func compileGraphExecutor(graph Graph) graphExecutor {
	graphExecutor := graphExecutor{}
	graphExecutor.midiInput = graph.ProcessorList[0].Processor.(*processorbuiltin.MIDIInput)
	graphExecutor.outputTerm = graph.ProcessorList[4].Processor.(*processorbuiltin.Terminal)

	graphExecutor.varCount = 8
	graphExecutor.ops = append(graphExecutor.ops, graphOp{graph.ProcessorList[0].Processor, []int{}, []int{1, 2, 3, 0, 0, 0, 0}})
	graphExecutor.ops = append(graphExecutor.ops, graphOp{graph.ProcessorList[1].Processor, []int{1}, []int{4}})
	graphExecutor.ops = append(graphExecutor.ops, graphOp{graph.ProcessorList[2].Processor, []int{2, 3}, []int{5}})
	graphExecutor.ops = append(graphExecutor.ops, graphOp{graph.ProcessorList[3].Processor, []int{4, 5}, []int{6}})
	graphExecutor.ops = append(graphExecutor.ops, graphOp{graph.ProcessorList[4].Processor, []int{6, 6}, []int{}})

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
