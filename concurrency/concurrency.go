package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var example string

func main() {
	flag.StringVar(&example, "example", "random-calculator", "Set the example")
	flag.Parse()

	switch example {
	case "random-calculator":
		randomCalculator()
	case "count-down":
		countDown()
	default:
		fmt.Println("Unknown example")
	}
}

/*
	START - Generator
	example for Generator pattern
*/

// countDown call the generator and handle every message
// what came on the channel
func countDown() {
	for counter := range countDownGenerator(6) {
		fmt.Printf("> %d\n", counter)
	}
}

// countDownGenerator is generating the own communication channel
// and return it for the caller function for usage
// then start to count back from the parameter
func countDownGenerator(from int) <-chan int {
	c := make(chan int)
	go func() {
		for i := from; i > 0; i-- {
			c <- i
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

/*
	END - Generator
*/

/*
	START - Random Calculator example
	example for synchronize channels
*/

// randomCalculator creating the communication channels
// and passed them to the worker
// in random times creating some random int and put them into a and b
// then waiting for the result which is not necessary,
// just to avoid the ddos on ourself
func randomCalculator() {
	c := make(chan int)
	a := make(chan int)
	b := make(chan int)
	go worker(a, b, c)
	for {
		va := random(1, 200)
		vb := random(1, 200)
		a <- va
		b <- vb
		fmt.Printf("A: %d B: %d\n", va, vb)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		result := <-c
		fmt.Printf("Result: %d\n", result)
	}
}

// Worker waiting for a and b through channels
// when get them, make the sum operation
// and put the result into c which is the result channel
func worker(a, b, c chan int) {
	var va, vb int
	for {
		va = <-a
		vb = <-b
		c <- va + vb
	}
}

/*
	END - Random Calculator example
*/

/*
	Utils
*/

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}
