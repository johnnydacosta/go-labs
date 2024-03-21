package main

import "testing"

func BenchmarkProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process()
	}
}

func BenchmarkProcessConcurrently(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processConcurrently()
	}
}
