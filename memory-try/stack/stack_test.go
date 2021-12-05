package stack

import "testing"

func BenchmarkTest(b *testing.B) {
	for i:=0;i<b.N;i++ {
		b.ReportAllocs()
		Test()
	}
}
