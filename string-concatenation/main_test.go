package main

import "testing"

var stringArray = []string{"test ", "testing", " ", "test awesome"}

func BenchmarkStringConcatWithFmt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stringConcatWithFmt(stringArray...)
	}
}

var new string

func BenchmarkStringConcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stringConcat(stringArray...)
	}
}
