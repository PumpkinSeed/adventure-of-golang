package main

type Robi struct {
	Field1 string
}

//go:noinline
func (r Robi) Test() Robi {
	println(&r.Field1)
	//println(r.Field1)
	return r
}

func main() {
	var r = Robi{"ads"}
	r2 := r.Test()
	for i:=0;i<10000;i++{
		r2.Field1 += "lofasz"
	}
	r2.Test()
	r.Test()
	println(&r2)
}
