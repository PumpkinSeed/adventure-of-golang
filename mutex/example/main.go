package main

import (
	"fmt"
	"sync"

	"github.com/PumpkinSeed/adventure-of-golang/mutex"
	"github.com/PumpkinSeed/adventure-of-golang/mutex/atomic-counter"
)

func main() {
	o := ordinal.NewOrdinal()
	m := mutex.NewSafeMap()
	o.Init(1123)
	fmt.Println("initial ordinal is:", o.GetOrdinal())

	wg := sync.WaitGroup{}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Set(fmt.Sprint(i), "success")
			o.Increment()
		}(i)
	}

	wg.Wait()
	for i := 0; i < 10000; i++ {
		v, err := m.Get(fmt.Sprint(i))
		if err != nil || v != "success" {
			panic(err)
		}
	}
	fmt.Println("final ordinal is:", o.GetOrdinal())
	fmt.Println("all keys found and marked as: 'success'")
}
