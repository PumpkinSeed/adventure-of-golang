package main

import (
	"log"
	"runtime"
	"time"
)

var counter int

/*
Example of the garbage collection, the Alloc sections says the currently
used and allocated memory, TotalAlloc says the whole of that, and NumGC says
the number of the garbage collections
*/

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(8)
	counter = 1
	go memoryUsage()
	for {
		go fib(counter)
		counter++
		time.Sleep(1 * time.Millisecond)
		if counter == 100000 {
			break
		}
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	since := time.Since(start)
	log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v Fibs = %d Spent: %dms\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC, counter, int64(since/time.Millisecond))
}

func memoryUsage() {
	for {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v Fibs = %d\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC, counter)
		time.Sleep(200 * time.Millisecond)
	}
}

func fib(counter int) {
	first := 0
	second := 1
	for {
		swap := first
		first = second
		second = swap + first
		if second == 99194853094755497 {
			break
		}
		time.Sleep(1 * time.Millisecond)
	}
}
