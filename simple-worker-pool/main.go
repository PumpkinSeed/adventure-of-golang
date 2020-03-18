package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mes := time.Now()
	var ch = make(chan bool)
	var wg = sync.WaitGroup{}

	for i := 0; i< 4; i++ {
		wg.Add(1)
		go worker(ch, &wg)
	}

	for i:= 0;i<10000; i++ {
		func(ch chan bool) {
			ch <- true
		}(ch)
	}
	close(ch)

	wg.Wait()
	released := time.Since(mes)
	fmt.Println(released)
}

func worker(ch chan bool, wg *sync.WaitGroup) {
	var inc int

	for t := range ch {
		if t {
			inc++
		}
	}

	println(inc)
	wg.Done()
}
