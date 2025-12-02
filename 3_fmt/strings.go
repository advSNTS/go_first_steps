package main

import "fmt"

func main() {

	x := "Hello world"
	y := 1834
	fmt.Println(x)

	fmt.Printf("Number: %d. This is a hello world message: %s\n", y, x)

	var z string = fmt.Sprintf("Hey: %s", x)

	fmt.Println(z)
}
