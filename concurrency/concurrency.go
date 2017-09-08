package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
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
	case "multiplexing":
		multiplexing()
	default:
		fmt.Println("Unknown example")
	}
}

/*
	START - Multiplexing
	example for Multiplexing pattern
*/

// signal represent a communication signal layout
type signal struct {
	data string
	done bool
}

func newSignal(data string, done bool) signal {
	return signal{
		data: data,
		done: done,
	}
}

// the fanIn function get n input signal, and create go routine for each of
// when there is a message on the channel write into the response channel
func fanIn(inputs ...<-chan signal) <-chan string {
	c := make(chan string)
	for counter, input := range inputs {
		go func(ch <-chan signal, co int) {
			for {
				s := <-ch
				if s.done {
					break
				}
				c <- getTag(co) + s.data
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			}
		}(input, counter)
	}
	return c
}

// format the output
func getTag(c int) string {
	return strconv.Itoa(c) + " > "
}

// multiplexing call the fanIn function, the parameters countDown function
// listening on the channel what return by the fanIn function
// there is a select inside the listener function to avoid the deadlock
func multiplexing() {
	c := fanIn(countDownGeneratorString(12), countDownGeneratorString(4), countDownGeneratorString(7))
	for {
		select {
		case value := <-c:
			fmt.Println(value)
		default:
		}
	}
}

// the function get an initial variable and start to count down
// each number take into channel, the signal type of channel store the data and
// the done cancellation option
func countDownGeneratorString(from int) <-chan signal {
	c := make(chan signal)
	go func() {
		for i := from; i > 0; i-- {
			c <- newSignal(strconv.Itoa(i), false)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		c <- newSignal("", true)
		close(c)
	}()
	return c
}

/*
	END - Multiplexing
*/

/*
	START - Generator
	example for Generator pattern
*/

// countDown call the generator and handle every message
// what came on the channel
// the generator return the communication channel, and we can have more instance
func countDown() {
	go func() {
		for counter := range countDownGenerator(5) {
			fmt.Printf(">>> %d\n", counter)
		}
	}()
	for counter := range countDownGenerator(11) {
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
