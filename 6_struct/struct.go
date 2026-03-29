package main

import "fmt"

type book struct {
	title     string
	editorial string
	isbn      string
	Author    author
}

type author struct {
	name string
	age  int
}

func main() {
	myBook := book{}
	myBook.Author.name = "Lois"
	fmt.Print(myBook.Author.name, "\n")
}
