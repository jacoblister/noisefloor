package dsp

import (
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"
	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor/processorbuiltin"
	"github.com/jacoblister/noisefloor/pkg/midi"
	"github.com/jacoblister/noisefloor/pkg/vfs"
)

const workdir = "workspace"

// ProcessEventFunc is a callback on update of DSP processing
type ProcessEventFunc func()

// Engine - DSP processing engine
type Engine struct {
	midiinput        processorbuiltin.MIDIInput
	patch            PatchMultiply
	osc              processor.Oscillator
	Graph            Graph
	compiledGraph    compiledGraph
	processEventSkip int
	processEventFunc ProcessEventFunc
	filename         string
}

// SetProcessEventFunc sets a notify callback when a process update occurs
func (e *Engine) SetProcessEventFunc(processEventFunc ProcessEventFunc) {
	e.processEventFunc = processEventFunc
}

// Start initilized the engine, with a specified sampling rate
func (e *Engine) Start(sampleRate int) {
	println("do DSP start, sample rate:", sampleRate)
	// e.compiledGraph.Start(sampleRate)

	e.osc.Start(sampleRate)
	e.midiinput.Start(sampleRate)
	e.patch.Start(sampleRate)
	e.osc.Waveform = processor.Sin
}

// Stop suspends the engine
func (e *Engine) Stop() {
	println("do DSP stop")
}

// Process processes a block of samples and midi events
func (e *Engine) Process(samplesIn [][]float32, midiIn []midi.Event) (samplesOut [][]float32, midiOut []midi.Event) {
	// e.midiinput.ProcessMIDI(midiIn)
	// var len = len(samplesIn[0])
	// for i := 0; i < len; i++ {
	// 	// var sample = e.osc.Process()
	// 	// freqs := e.midiinput.Process()
	// 	var sample = e.patch.Process(&e.midiinput)
	// 	sample += samplesIn[0][i]
	//
	// 	// mic := samplesIn[0][i] * 500
	// 	// mod := e.osc.Process()
	// 	// if mod < 0 {
	// 	// 	mod = 0
	// 	// }
	// 	// mic *= mod
	// 	// sample += mic
	// 	// sample = e.osc.Process()
	//
	// 	samplesIn[0][i] = sample
	// 	samplesIn[1][i] = sample
	// }

	// if e.compiledGraph != nil {
	// 	samplesIn, midiIn = e.compiledGraph.Process(samplesIn, midiIn)
	// }

	// notify front end if registered
	if e.processEventFunc != nil {
		e.processEventSkip--
		if e.processEventSkip <= 0 {
			e.processEventSkip = 3
			e.processEventFunc()
		}
	}

	len := len(samplesIn[0])
	for i := 0; i < len; i++ {
		sample := e.osc.Process(440)
		samplesIn[0][i] = sample
		samplesIn[1][i] = sample
	}

	return samplesIn, midiIn
}

// recompileGraph recompiles the current graph
func (e *Engine) recompileGraph() {
	println("Recompile graph")
	e.compiledGraph = nil

	// TODO - totally wrong place for this - avoid race conditon in engine startup
	compiledGraph := compileProcessorGraph(e.Graph, CompileInterpretedSingleSample)
	compiledGraph.Start(48000)

	e.compiledGraph = compiledGraph
}

// GraphChange is called when the graph changes, with indication if recompile is required
func (e *Engine) GraphChange(recompile bool) {
	if recompile {
		e.recompileGraph()
	}
	e.Save(workdir + "/" + e.filename)
}

// Filename returns the filename of the currently loaded graph
func (e *Engine) Filename() string {
	return e.filename
}

// Files returns a list of files in the working directory
func (e *Engine) Files() []string {
	result := []string{}

	dir, _ := vfs.DefaultFS().Open(workdir)
	fileInfo, _ := dir.Readdir(-1)

	for i := 0; i < len(fileInfo); i++ {
		result = append(result, fileInfo[i].Name())
	}

	return result
}

// Load loads a graph into the synthengine from file
func (e *Engine) Load(filename string) {
	// e.Graph = exampleGraph()

	fullname := workdir + "/" + filename

	file, _ := vfs.DefaultFS().Open(fullname)
	graph, err := loadProcessorGraph(file)
	if err != nil {
		println("Error loading", fullname, ":", err.Error())
	} else {
		e.Graph = graph
		e.filename = filename
	}

	e.recompileGraph()
}

// Save saves the graph to the specified file
func (e *Engine) Save(filename string) {
	// file, _ := os.Create("out.xml")
	// saveProcessorGraph(e.Graph, file)
	// file.Close()
}
