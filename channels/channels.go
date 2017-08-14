package main

import (
	"flag"
	"fmt"
	"time"
)

var example string

func init() {
	flag.StringVar(&example, "example", "1", "Number of example")
	flag.Parse()
}

func main() {
	switch example {
	case "1":
		example1()
	case "2":
		example2()
	}
}

/* Start Example 1
The main function get block for ten seconds from sender before
print out because the msg := <-messages section obligate the runtime
for waiting a response on the messages channel
*/
func example1() {
	fmt.Println("Example 1")
	messages := make(chan string)
	go func() {
		fmt.Println("Func goroutine sleeps 10 seconds")
		time.Sleep(10 * time.Second)
		fmt.Println("Func goroutine begins sending data")
		messages <- "ping"
		fmt.Println("Func goroutine ends sending data")
	}()
	fmt.Println("Main goroutine begins receiving data")
	msg := <-messages
	fmt.Println("Main goroutine received data:", msg)
}

// End Example 1

/* Start Example 2
The goroutine get block because of the receiver sleeping in the
main funciton
*/
func example2() {
	fmt.Println("Example 2")
	messages := make(chan string)

	go func() {
		fmt.Println("Func goroutine begins sending data")
		messages <- "ping"
		fmt.Println("Func goroutine ends sending data")
	}()

	fmt.Println("Main goroutine sleeps 10 seconds")
	time.Sleep(time.Second * 10)

	fmt.Println("Main goroutine begins receiving data")
	msq := <-messages
	fmt.Println("Main goroutine received data:", msq)

	time.Sleep(time.Second)
}

// End Example 2
