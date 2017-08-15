package main

import "testing"

func BenchmarkByteSliceToString(b *testing.B) {
	test := []byte("Test it!")

	for n := 0; n < b.N; n++ {
		byteSliceToString(test)
	}
}

var new string

func BenchmarkByteSliceToStringCast(b *testing.B) {
	test := []byte("Test it!")

	for n := 0; n < b.N; n++ {
		new = string(test)
	}
}
