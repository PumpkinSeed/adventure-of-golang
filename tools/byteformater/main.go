package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	f string
)

func init() {
	flag.StringVar(&f, "file", "", "")
	flag.Parse()
}

func main() {
	f1, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	data, err := ioutil.ReadAll(f1)
	if err != nil {
		panic(err)
	}

	str := string(data)
	elems := strings.Split(str, " ")
	str = strings.Join(elems, ", ")
	str = strings.Replace(str, "[", "", -1)
	str = strings.Replace(str, "]", "", -1)
	result := fmt.Sprintf("[]byte{%s}", str)
	ioutil.WriteFile("output.txt", []byte(result), 0644)
}

