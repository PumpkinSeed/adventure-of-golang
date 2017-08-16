package main

import (
	"fmt"
)

func main() {

}

func stringConcatWithFmt(str ...string) string {
	return fmt.Sprint(str)
}

func stringConcat(str ...string) string {
	var result string
	for _, e := range str {
		result += e
	}
	return result
}
