package pointer

import "testing"

func BenchmarkTest(b *testing.B) {
	for i:=0;i<b.N;i++ {
		b.ReportAllocs()
		Test()
	}
}

func BenchmarkTestWithAlloc(b *testing.B) {
	for i:=0;i<b.N;i++ {
		b.ReportAllocs()
		TestWithAlloc()
	}
}
