package main

import "testing"

func BenchmarkNewComplex(b *testing.B) {
	b.ReportAllocs()
	for i:=0;i<b.N;i++ {
		c := newComplex()
		c.field1 = "2"
	}
}
