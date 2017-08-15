package main

import "unsafe"
import "fmt"

func main() {
	unsafePointer([]byte("test"))
}

func byteSliceToString(bs []byte) string {
	// This is copied from runtime. It relies on the string
	// header being a prefix of the slice header!
	return *(*string)(unsafe.Pointer(&bs))
}

/*
	Understand the code above
*/

func unsafePointer(bs []byte) {
	fmt.Println(unsafe.Pointer(&bs))
}
