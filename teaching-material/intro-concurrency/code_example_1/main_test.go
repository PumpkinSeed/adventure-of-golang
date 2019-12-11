package main

import (
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func BenchmarkSequential(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	n := numbers()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		add(n)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	n := numbers()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		addConcurrent(runtime.NumCPU(), n)
	}
}

func numbers() []int {
	var nums []int
	for i := 0; i<10000000; i++ {
		nums = append(nums, randomInt(1, 30000))
	}

	return nums
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}