package synth

import (
	"testing"

	"github.com/jacoblister/noisefloor/app/audiomodule/synth/processor"
)

func BenchmarkSingleCall(b *testing.B) {
	gain := processor.Gain{}

	for i := 0; i < b.N; i++ {
		gain.Process(1, 1)
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
