package cppprocessor

import "testing"

func BenchmarkAddInt(b *testing.B) {
	var x int32 = 1
	var y int32 = 1
	var res int32
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			res += x + y
		}
	}
}

func BenchmarkAddFloat32(b *testing.B) {
	var x float32 = 1.2
	var y float32 = 1.5
	var res float32
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			res += x + y
		}
	}
}

func BenchmarkNothing(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

func BenchmarkCCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CCall()
	}
}

func BenchmarkCCallCallback(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CCallCallback()
	}
}

func BenchmarkCProcess(b *testing.B) {
	CStart()
	for i := 0; i < b.N; i++ {
		CProcess()
	}
}
