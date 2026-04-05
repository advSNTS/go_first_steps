package main

import "fmt"

type book struct {
	name string
	isbn string
}

type getName interface {
	getName() string
}

func (b book) getName() string {
	return b.name
}

func main() {
	b := book{
		name: "Hail Mary",
		isbn: "92745-3947",
	}

	fmt.Printf("Name: %s  \n", b.getName())
}
