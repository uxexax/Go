package main

import "fmt"

func main() {
	var f float64

	fmt.Print("Enter a floating point number: ")
	n, err := fmt.Scan(&f)

	if err == nil {
		fmt.Printf("The truncated number is: %d", int(f))
	} else {
		fmt.Printf("Scanned %d items due to the following error: %s\n", n, err)
	}
}
