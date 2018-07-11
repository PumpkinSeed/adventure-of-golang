package main

import "testing"

var stringArray = []string{"test ", "testing", " ", "test awesome"}

func BenchmarkStringConcatWithFmt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stringConcatWithFmt(stringArray...)
	}
}

func BenchmarkStringConcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stringConcat(stringArray...)
	}
}

func BenchmarkStringConcatBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stringConcatBuffer(stringArray...)
	}
}
