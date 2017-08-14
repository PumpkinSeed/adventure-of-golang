package main

import "fmt"

func val1(section int) bool {
	fmt.Println("val1 gets called in section: ", section)
	return true
}

func val2(section int) bool {
	fmt.Println("val2 gets called in section: ", section)
	return false
}

/*
	The short circuit evaluation (wikipedia):
	is the semantics of some Boolean operators in some programming languages
	in which the second argument is executed or evaluated only if the first
	argument does not suffice to determine the value of the expression

	The section 2 and 3 is evaulate the condition by the first argument, so
	the second argument is not necessary to evaulate in this case
*/
func main() {
	if val1(1) && val2(1) {
		fmt.Println("Shouldn't print")
	}
	if val2(2) && val1(2) {
		fmt.Println("Shouldn't print")
	}
	if val1(3) || val2(3) {
		fmt.Println("The boolean expression is true")
	}
	if val2(4) || val1(4) {
		fmt.Println("The boolean expression is true")
	}
}
