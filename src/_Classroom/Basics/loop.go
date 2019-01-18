package main

import "fmt"

func main() {
	var xtemp int
	x1 := 0
	x2 := 1
	for x := 0; x < 5; x++ {
		xtemp = x2
		x2 = x2 + x1
		x1 = xtemp
		fmt.Printf("%d: x1 = %d, x2 = %d\n", x, x1, x2)
	}
	fmt.Printf("Final x2 is %d.\n", x2)
}
