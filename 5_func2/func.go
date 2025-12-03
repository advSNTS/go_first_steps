package main

import "fmt"

func main() {

	j := 4
	i := 6
	fmt.Println(sumAndMul(j, i))

	first, _ := sumAndMul(j, i)
	_, second := sumAndMul(j, i)
	fmt.Println("First: ", first)
	fmt.Println("Second: ", second)

	var first1, _ int = sumAndMul(j, i)
	fmt.Println(first1)
}

func sumAndMul(x, y int) (int, int) {
	return x + y, x * y
}
