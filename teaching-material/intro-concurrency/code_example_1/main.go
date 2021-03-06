package main

import (
	"sync"
	"sync/atomic"
)

func main() {

}

func add(numbers []int) int {
	var v int
	for _, n := range numbers {
		v += n
	}
	return v
}

func addConcurrent(goroutines int, numbers []int) int {
	var v int64
	totalNumbers := len(numbers)
	lastGoroutine := goroutines -1
	stride := totalNumbers / goroutines // determine the length of the smaller list handled by goroutines

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ { // loop over the goroutine pool, each iteration starts a new concurrent function
		go func(g int) {
			start := g * stride
			end := start + stride
			if g == lastGoroutine {
				end = totalNumbers // make sure the last goroutine won't run out from the numbers
			}

			var lv int
			for _, n := range numbers[start:end] {
				lv += n
			}
			atomic.AddInt64(&v, int64(lv)) // atomic operation provide the thread safety
			wg.Done()
		} (g)
	}

	wg.Wait()

	return int(v)
}