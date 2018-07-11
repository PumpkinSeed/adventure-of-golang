package main

import (
	"bytes"
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

func stringConcatBuffer(str ...string) string {
	var buffer bytes.Buffer

	for _, e := range str {
		buffer.WriteString(e)
	}

	return buffer.String()
}
