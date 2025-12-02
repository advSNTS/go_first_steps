package main

import "fmt"

func main() {
	var n int = 8
	fmt.Printf("The %d value of fibonacci is: %d \n", n, fibo(n))
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
