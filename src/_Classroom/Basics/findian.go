package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a string: ")
	scanner.Scan()

	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
		return
	}

	s := strings.ToLower(scanner.Text())

	if strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") && strings.Contains(s, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
