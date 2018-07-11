# String concatenation

The string concatenation without fmt much faster.

```
goos: darwin
goarch: amd64
pkg: github.com/PumpkinSeed/adventure-of-golang/string-concatenation
BenchmarkStringConcatWithFmt-4   	 2000000	       744 ns/op
BenchmarkStringConcat-4          	10000000	       164 ns/op
BenchmarkStringConcatBuffer-4    	10000000	       136 ns/op
```