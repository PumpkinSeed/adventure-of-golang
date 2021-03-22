package main

import (
	"fmt"
	"github.com/hashicorp/vault/shamir"
	"log"
)

func main() {
	var secret = "test"

	parts, err := shamir.Split([]byte(secret), 3, 2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(parts)
	var newParts = [][]byte{}
	newParts = append(newParts, parts[1])
	newParts = append(newParts, parts[2])
	data, err := shamir.Combine(newParts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
