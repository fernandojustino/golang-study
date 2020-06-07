package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Fernando", 20})

	fmt.Println(newPerson("Jose"))

	p1 := newPerson("Marcelo")

	fmt.Println(p1)

	s := person{name: "Luca", age: 21}

	fmt.Println(&s)
	fmt.Println(s)
}
