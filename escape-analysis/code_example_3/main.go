package main

type entity struct {
	field1 string
	field2 string
}

func main() {
	e1 := createEntity()
	e2 := createEntityPtr()

	println("e1", &e1, "e2", &e2)
}

func createEntity() entity {
	e := entity{
		field1: "test",
		field2: "test2",
	}

	println("Non-pointer", &e)
	return e
}

func createEntityPtr() *entity {
	e := entity{
		field1: "test",
		field2: "test2",
	}

	println("Pointer", &e)
	return &e
}