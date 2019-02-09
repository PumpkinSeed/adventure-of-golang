package main

import (
	"fmt"
	"strings"
)

var data = `2 75 49 46 90 82 64 64 64 64 48 48 48 48 48 48 48 48 50 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 48 82 66 48 48 49 50 51 85 48 48 48 51 90 68 80 64 64 64 53 52 68 52 49 49 49 49 49 49 49 49 49 49 3 100`

func main() {
	fmt.Println(strings.Join(strings.Split(data, " "), ", "))
}
