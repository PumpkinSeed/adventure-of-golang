package main

import (
	"fmt"

	worker "github.com/PumpkinSeed/adventure-of-golang/worker-pool-pattern"
)

func main() {
	cancel, in, out := worker.Dispatch(10)
	defer cancel()

	for i := 0; i < 10; i++ {
		in <- worker.WorkRequest{
			Op:   worker.Hash,
			Text: []byte(fmt.Sprintf("messages %d", i)),
		}
	}

	for i := 0; i < 10; i++ {
		res := <-out
		if res.Err != nil {
			panic(res.Err)
		}
		in <- worker.WorkRequest{
			Op:      worker.Compare,
			Text:    res.Wr.Text,
			Compare: res.Result,
		}
	}

	for i := 0; i < 10; i++ {
		res := <-out
		if res.Err != nil {
			panic(res.Err)
		}
		fmt.Printf("string: '%s'; matched: %v\n", string(res.Wr.Text), res.Matched)
	}
}
