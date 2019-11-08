package dsp

import (
	"reflect"
	"testing"

	"github.com/jacoblister/noisefloor/app/audiomodule/dsp/processor"
	"github.com/jacoblister/noisefloor/pkg/midi"
)

func BenchmarkArrayLookup(b *testing.B) {
	array := []float32{0, 1, 2, 3}

	for i := 0; i < b.N; i++ {
		_ = array[3]
	}
}

func BenchmarkMapLookup(b *testing.B) {
	array := map[string]float32{"one": 1, "two": 2, "three": 3, "four": 4}

	for i := 0; i < b.N; i++ {
		_ = array["three"]
	}
}

func BenchmarkSingleCall(b *testing.B) {
	gain := processor.Gain{}

	for i := 0; i < b.N; i++ {
		gain.Process(1, 1)
	}
}

func BenchmarkSingleCallWithArray(b *testing.B) {
	gain := processor.Gain{}
	params := []float32{0, 0}

	callGain := func(gain *processor.Gain, in []float32) []float32 {
		in[0] = gain.Process(in[0], in[1])
		return in
	}

	for i := 0; i < b.N; i++ {
		params[0] = 1
		params[1] = 1
		callGain(&gain, params)
	}
}

func BenchmarkSingleCallWithMethodValues(b *testing.B) {
	gain := processor.Gain{}
	callGain := gain.Process

	for i := 0; i < b.N; i++ {
		callGain(1, 1)
	}
}

func BenchmarkSingleCallWithReflection(b *testing.B) {
	gain := processor.Gain{}
	processMethod := reflect.ValueOf(&gain).MethodByName("Process")
	values := []reflect.Value{reflect.ValueOf(float32(1.0)), reflect.ValueOf(float32(1.0))}

	for i := 0; i < b.N; i++ {
		processMethod.Call(values)
	}
}

type golangEngine struct {
	osc  processor.Oscillator
	env  processor.Envelope
	gain processor.Gain
}

func (g *golangEngine) Start(sampleRate int) {
	g.osc.Start(sampleRate)
	g.env.Start(sampleRate)
	g.gain.Start(sampleRate)
}

func (g *golangEngine) Process(samplesIn [][]float32, midiIn []midi.Event) (samplesOut [][]float32, midiOut []midi.Event) {
	var len = len(samplesIn[0])

	for i := 0; i < len; i++ {
		osc := g.osc.Process(440)
		env := g.env.Process(0, 0)
		out := g.gain.Process(osc, env)
		samplesIn[0][i] = out
		samplesIn[1][i] = out
	}
	return samplesIn, midiIn
}

func (g *golangEngine) Stop() {
	g.osc.Stop()
}

func BenchmarkCompileGolang(b *testing.B) {
	blockSize := 1024
	samples := [][]float32{}
	samples = append(samples, make([]float32, blockSize, blockSize))
	samples = append(samples, make([]float32, blockSize, blockSize))

	process := golangEngine{}
	// process := compileProcessorGraph(Graph{}, CompileGolang)
	process.Start(44100)

	for i := 0; i < b.N; i++ {
		process.Process(samples, nil)
	}

	process.Stop()
}

func BenchmarkCompileIntepreted(b *testing.B) {
	blockSize := 1024
	samples := [][]float32{}
	samples = append(samples, make([]float32, blockSize, blockSize))
	samples = append(samples, make([]float32, blockSize, blockSize))

	graph := loadProcessorGraph("")
	process := compileProcessorGraph(graph, CompileInterpreted)

	process.Start(44100)

	for i := 0; i < b.N; i++ {
		process.Process(samples, nil)
	}

	process.Stop()
}

func BenchmarkCompileIntepretedSingleSample(b *testing.B) {
	blockSize := 1024
	samples := [][]float32{}
	samples = append(samples, make([]float32, blockSize, blockSize))
	samples = append(samples, make([]float32, blockSize, blockSize))

	graph := loadProcessorGraph("")
	process := compileProcessorGraph(graph, CompileInterpretedSingleSample)

	process.Start(44100)

	for i := 0; i < b.N; i++ {
		process.Process(samples, nil)
	}

	process.Stop()
}
