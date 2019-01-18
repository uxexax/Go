/* ----------
This program asks the user to provide a sequence of integers on the standard input,
which is then sorted using bubble sort algorithm. The result is printed to the
standard output.
---------- */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputStr, inputErr := AskForInput()
	if inputErr != nil {
		fmt.Println("Error:", inputErr, "\n")
		os.Exit(1)
	}

	sequenceOfInts, parseError := ParseInput(inputStr)
	if parseError != nil {
		fmt.Println("Error:", parseError, "\n")
		os.Exit(2)
	}

	BubbleSort(sequenceOfInts)

	fmt.Println("The sorted sequence is:")
	fmt.Println(sequenceOfInts, "\n")
}

// BubbleSort implements the bubble sort algorithm for integers.
func BubbleSort(sequence []int) {
	for upperIndex := len(sequence) - 1; upperIndex > 0; upperIndex-- {
		for runningIndex := 0; runningIndex < upperIndex; runningIndex++ {
			if sequence[runningIndex] > sequence[runningIndex+1] {
				Swap(sequence, runningIndex)
			}
		}
	}
}

// Swap takes two adjacent elements in a slice of integers and swaps them.
func Swap(sequence []int, swapAt int) {
	temp := sequence[swapAt]
	sequence[swapAt] = sequence[swapAt+1]
	sequence[swapAt+1] = temp
}

// AskForInput retrieves input data from the user via standard input.
func AskForInput() (string, error) {
	fmt.Println("\nEnter a sequence of integers, divided by spaces.")
	fmt.Print(">> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Println()

	return scanner.Text(), scanner.Err()
}

// ParseInput takes a string of integers separated by spaces,
// and parses it to a slice of integers.
func ParseInput(inputLine string) ([]int, error) {
	sizeLimit := 10
	parsedSequence := make([]int, 0, sizeLimit)

	splitLine := strings.Fields(inputLine)

	for i, v := range splitLine {
		if i == sizeLimit {
			fmt.Printf("(!) Sequence size is limited to %d, last %d element(s) are ignored.\n\n",
				sizeLimit, len(splitLine)-sizeLimit)
			break
		}

		intValue, convError := strconv.Atoi(v)
		if convError != nil {
			return nil, convError
		}

		parsedSequence = append(parsedSequence, intValue)
	}

	return parsedSequence, nil
}
