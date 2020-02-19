package main

import (
	"fmt"
	"time"
)

func main() {
	mes := time.Now()
	var ch = make(chan bool)
	for i := 0; i< 4; i++ {
		go worker(ch)
	}

	for i:= 0;i<1000; i++ {
		ch <- true
	}
	close(ch)
	released := time.Since(mes)
	time.Sleep(1*time.Millisecond)

	fmt.Println(released)
}

func worker(ch chan bool) {
	var inc int

	for t := range ch {
		if t {
			inc++
		}
	}

	println(inc)
}
