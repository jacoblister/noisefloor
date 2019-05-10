package synth

import (
	"reflect"
	"testing"

	"github.com/jacoblister/noisefloor/app/audiomodule/synth/processor"
)

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

func BenchmarkSingleCompileGolang(b *testing.B) {
	blockSize := 1024
	samples := [][]float32{}
	samples = append(samples, make([]float32, blockSize, blockSize))
	samples = append(samples, make([]float32, blockSize, blockSize))

	processFunc := compileProcessorGraph(Graph{}, CompileGolang)

	for i := 0; i < b.N; i++ {
		processFunc(samples, nil)
	}
}
