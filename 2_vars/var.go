package main

import "fmt"

func main() {

	const place, locale string = "usa", "when"
	var x, y int = 1, 2
	fmt.Println(place, locale)

	fmt.Printf("x: %d, y: %d\n", x, y)
	temp := x
	x = y
	y = temp
	fmt.Printf("x2: %d, y2: %d\n", x, y)

}
