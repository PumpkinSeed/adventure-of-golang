package main

type entity struct {
	field1 string
	field2 string
}

type complex struct {
	field1 string
	entity *entity
}

func main() {
	complexStructure := newComplex()

	println("e1", &complexStructure)
}

//go:noinline
func newComplex() complex {
	c := complex{
		field1: "test",
		entity: &entity{
			field1: "test1",
			field2: "test2",
		},
	}

	//println("Pointer", c.entity)
	return c
}