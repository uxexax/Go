package main

import (
	"fmt"
)

func main() {
	var number float64
	var truncated int

	fmt.Println("Enter floating point number:")
	fmt.Scanf("%f", &number)

	truncated = int(number)
	fmt.Println(truncated)
}
