package main

import "testing"

func BenchmarkGoCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeGoCall()
	}
}

func BenchmarkCCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeCCall()
	}
}

func BenchmarkCCallBack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeCCallBack()
	}
}
