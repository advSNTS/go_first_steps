package main

import "fmt"

type car struct {
	model string
	year  int
}

type semi struct {
	car
	size int
}

func main() {
	s := semi{
		car: car{
			model: "jac",
			year:  2020,
		},
		size: 4378,
	}

	fmt.Print(s, "\n")
}
