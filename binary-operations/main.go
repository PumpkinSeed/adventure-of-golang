package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const inNum uint32 = 512

func main() {
	// in-memory buffer
	buf := &bytes.Buffer{}

	// creates the input value, uint32 (32/8 = 4 bytes)
	val := uint32(inNum)
	fmt.Printf("val: %032b\n", val)

	fmt.Printf("buffer size before: %d, should be 0\n", buf.Len())

	// writes the uint32 to the bytes.Buffer
	binary.Write(buf, binary.BigEndian, val)

	// show what's in the buffer, should have a length of 4
	fmt.Printf("buffer size after: %d, should be 4\n", buf.Len())
	fmt.Println(buf.Bytes())

	// uint32 to read from the bytes.Buffer
	var readBinary uint32
	fmt.Printf("readBinary before: %032b\n", readBinary)

	// read in to the uint32
	binary.Read(buf, binary.BigEndian, &readBinary)

	// print the binary value of readBinary
	fmt.Printf("readBinary after: %032b\n", readBinary)

	// print details
	fmt.Println(val, readBinary, val == readBinary)
}
