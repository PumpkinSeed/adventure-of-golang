package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

var example string

func init() {
	flag.StringVar(&example, "example", "fruits", "Number of example")
	flag.Parse()
}

func main() {
	switch example {
	case "fruits":
		var f = Fruit{"Fruit", 2}
		f.countForSeeds()
		j, _ := f.getJSON()
		fmt.Println(string(j))
		var apple = Apple{Fruit{"Apple", 1}, "England"}
		apple.countForSeeds()
		j, _ = apple.getJSON()
		fmt.Println(string(j))
	case "vegetables":
		var v = Vegetable{
			Name:  "Vegetable",
			Seeds: 2,
		}
		v.setSelf(v)
		v.countForSeeds()
		j, _ := v.getJSON()
		fmt.Println(string(j))
		var carrot = Carrot{
			Vegetable: Vegetable{
				Name:  "Carrot",
				Seeds: 0,
			},
			Country: "England",
		}
		carrot.setSelf(carrot)
		carrot.countForSeeds()
		j, _ = carrot.getJSON()
		fmt.Println(string(j))
	}
}

/* Fruits
The Apple have it's own getJSON method because the
Go not support method override, but can shadow it
*/

// Fruit struct
type Fruit struct {
	Name  string `json:"name"`
	Seeds int    `json:"seeds"`
}

func (f Fruit) countForSeeds() {
	fmt.Printf("%s have %d seeds\n", f.Name, f.Seeds)
}

func (f Fruit) getJSON() ([]byte, error) {
	return json.Marshal(&f)
}

// Apple struct
type Apple struct {
	Fruit
	Country string `json:"country"`
}

func (a Apple) getJSON() ([]byte, error) {
	return json.Marshal(&a)
}

// End of fruits

/* Vegetables
The carrot don't have the own getJSON function, but the
Vegetable have the self and the setSelf, and if we want to
us the getJSON just have to setSelf and use it
It is solved me if we have a BaseObject with 300+ embedded
struct, it's enough write the base functions only once
*/

// Vegetable struct
type Vegetable struct {
	Name  string `json:"name"`
	Seeds int    `json:"seeds"`
	self  interface{}
}

func (v Vegetable) countForSeeds() {
	fmt.Printf("%s have %d seeds\n", v.Name, v.Seeds)
}

func (v Vegetable) getJSON() ([]byte, error) {
	return json.Marshal(&v.self)
}

func (v *Vegetable) setSelf(p interface{}) {
	v.self = p
}

// Carrot struct
type Carrot struct {
	Vegetable
	Country string `json:"country"`
}

// End of vegetables
