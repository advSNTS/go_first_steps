package main

import "fmt"

func main() {
	var n int = 8
	fmt.Printf("The %d value of fibonacci using switch is: %d \n", n, fibo(n))
	fmt.Printf("The %d value of fibonacci using if is: %d \n", n, fiboif(n))
}

func fibo(n int) int {

	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibo(n-1) + fibo(n-2)
	}
}

func fiboif(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fiboif(n-1) + fiboif(n-2)
}
