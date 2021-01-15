package main

import (
	"fmt"
)

func main() {
	p := &Person{}
	p.modifyPerson(53, "Marc")

	fmt.Println(p)
}

type Person struct {
	age  int
	name string
}

func (p *Person) modifyPerson(age int, name string) {
	p.name = name
	p.age = age
}
