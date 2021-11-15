package cppprocessor

import "testing"

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
