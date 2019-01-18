package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	theSlice := make([]int, 0, 3)

	var inputValue string
	var inputError error

	var parsedNumber int
	var parseError error

	scanner := bufio.NewScanner(os.Stdin)

	i := 1

	fmt.Println("Enter one integer per line. Type 'X' to leave.")

infinity:
	for {
		fmt.Printf("Enter an integer (#%d): ", i)
		// _, inputError = fmt.Scan(&inputValue)

		scanner.Scan()

		inputValue = scanner.Text()
		inputError = scanner.Err()

		switch {
		case inputError != nil:
			fmt.Println(inputError)
			continue infinity
		case inputValue == "X":
			break infinity
		}

		parsedNumber, parseError = strconv.Atoi(inputValue)

		if parseError == nil {
			theSlice = append(theSlice, parsedNumber)
			sort.Sort(sort.IntSlice(theSlice))
			fmt.Println(theSlice)
		} else {
			fmt.Println(parseError)
			continue
		}

		i++
	}
}
