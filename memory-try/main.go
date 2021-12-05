package main

import (
	"github.com/PumpkinSeed/memory-try/pointer"
)

//go:noinline
func main() {
	for {
		//stack.Test()
		//pointer.Test()
		s := pointer.TestWithAlloc()
		s.Field1 = "ad"
	}
}