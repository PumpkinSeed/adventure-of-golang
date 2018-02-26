package main

import (
	"net/http"

	"github.com/PumpkinSeed/adventure-of-golang/raft/consensus"
)

func main() {
	consensus.Config(3)

	http.HandleFunc("/", consensus.Handler)
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
