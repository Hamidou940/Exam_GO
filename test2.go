package main

import "fmt"

func main() {
	sl := []string{"one", "two", "three"}

	for _, elt := range sl {
		fmt.Print(elt)
		fmt.Print(" ")
	}
}
