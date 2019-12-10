package main

import (
	"runtime"
	"testing"
)

func getDocs() []string {
	var docs []string
	for i := 0; i< 1000; i++ {
		docs = append(docs, file)
	}
	return docs
}

func BenchmarkSequential(b *testing.B) {
	docs := getDocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		find("amazing", docs)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	docs := getDocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findConcurrent(runtime.NumCPU(), "amazing", docs)
	}
}