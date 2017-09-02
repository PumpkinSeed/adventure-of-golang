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
	default:
		fmt.Println("Unknown example")
	}
}

/*
	START - Random Calculator
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
	END - Random Calculator
*/

/*
	Utils
*/

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}
