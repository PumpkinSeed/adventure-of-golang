package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pkg/profile"

	"github.com/PumpkinSeed/adventure-of-golang/string-builder/str"
)

var iterations = 100000

func main() {
	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	t := time.Now()

	withStringBuilder()

	fmt.Println(time.Since(t))
}

func withStringBuilder() {
	var builder strings.Builder

	for i := 0; i < iterations; i++ {
		builder.WriteString("cool")
	}

	//fmt.Println(builder.String())
}

func withStr() {
	var builder str.Builder

	for i := 0; i < iterations; i++ {
		builder.Write("cool")
	}

	//fmt.Println(builder.String())
}
