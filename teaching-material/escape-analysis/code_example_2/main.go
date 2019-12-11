package main

func main() {
	testdata := "test"
	println("main:5\tvalue:", testdata, "\taddress:", &testdata)

	f(&testdata)
	println("main:8\tvalue:", testdata, "\taddress:", &testdata)
}

//go:noinline
func f(str *string) {
	*str += "_f"
	println("f:14\tvalue:", str, "\taddress:", &str, "\tvalue to points to:", *str)
}
