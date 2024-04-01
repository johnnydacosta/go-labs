package main

import "testing"

func BenchmarkNaive(b *testing.B) { // 1 134 693 584 ns/op
	for i := 0; i < b.N; i++ {
		naiveApproach()
	}
}

func BenchmarkNaiveWithSleep(b *testing.B) { // 1 058 887 959 ns/op
	for i := 0; i < b.N; i++ {
		betterApproachButSillNotOptimize()
	}
}

func BenchmarkWithCondSleep(b *testing.B) { // 577 685 021 ns/op
	for i := 0; i < b.N; i++ {
		withSyncCond()
	}
}
