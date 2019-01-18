package main

import (
	"fmt"
)

func main() {
	go fmt.Println("Additional goroutine executed.")
	fmt.Println("Main goroutine executed.")
	// time.Sleep(time.Duration(100) * time.Millisecond)
}
