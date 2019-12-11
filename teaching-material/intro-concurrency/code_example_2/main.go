package main

import "sync"

func main() {}

func bubbleSort(numbers []int) {
	n := len(numbers)
	for i := 0; i<n; i++ {
		if !sweep(numbers, i) {
			return
		}
	}
}

func sweep(numbers []int, currentPass int) bool {
	var idx int
	idxNext := idx + 1
	n := len(numbers)
	var swap bool

	for idxNext < (n - currentPass) {
		a := numbers[idx]
		b := numbers[idxNext]
		if a > b {
			numbers[idx] = b
			numbers[idxNext] = a
			swap = true
		}
		idx++
		idxNext = idx+1
	}
	return swap
}

func bubbleSortConcurrent(goroutines int, numbers []int) {
	totalNumbers := len(numbers)
	lastGoroutine := goroutines - 1
	stride := totalNumbers / goroutines

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func(g int) {
			start := g * stride
			end := start + stride
			if g == lastGoroutine {
				end = totalNumbers
			}

			bubbleSort(numbers[start:end])
			wg.Done()
		}(g)
	}

	wg.Wait()

	bubbleSort(numbers)
}