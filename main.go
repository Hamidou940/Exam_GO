package main

import (
	"fmt"
)

func main() {
	var p Person
	modifyPerson(p, 53, "Marc")

	fmt.Println(p)
}

type Person struct {
	age  int
	name string
}

func modifyPerson(p Person, age int, name string) {
	p.name = name
	p.age = age
}
