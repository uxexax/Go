/* This program reads a file containing names into a slice of Name structures, and prints the
content to the standard output. The name of the file is obtained from the user.

Each line in the file should contain only one first name and one last name separated by a
space character. If the number of words in a line is more than two, only the first two words
are considered, the others are ignored. If there is less than two words in a line,
it is ignored. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type Name struct {
		fname string
		lname string
	}

	// ----- Obtain file name from the user
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\nEnter input file name: ")
	scanner.Scan()
	fmt.Println()

	if scanner.Err() != nil {
		fmt.Println("Error:", scanner.Err(), "\n")
		os.Exit(-1)
	}

	// ----- Open the file
	file, openErr := os.Open(scanner.Text())

	if openErr != nil {
		fmt.Println("File open error:", openErr, "\n")
		os.Exit(-2)
	}

	// ----- Read the file line by line
	fileReader := bufio.NewScanner(file)
	names := make([]Name, 0, 10)
	continueReading := true
	lineContent := make([]string, 2)
	warnsAlreadyGiven := map[string]bool{
		"lessThan2": false,
		"moreThan2": false}

readingLoop:
	for {
		continueReading = fileReader.Scan()

		if !continueReading {
			break
		}

		lineContent = strings.Fields(fileReader.Text())

		// If a line in the file contains not exactly two words a warning is given,
		// and the processing continues.
		switch {
		case len(lineContent) < 2:
			if !warnsAlreadyGiven["lessThan2"] {
				fmt.Println("(!) A line with less than two words is ignored.")
				warnsAlreadyGiven["lessThan2"] = true
			}
			continue readingLoop
		case len(lineContent) > 2:
			if !warnsAlreadyGiven["moreThan2"] {
				fmt.Println("(!) Only the first two words are read from a line.")
				warnsAlreadyGiven["moreThan2"] = true
			}
			continue readingLoop
		}

		names = append(names, Name{lineContent[0], lineContent[1]})
	}

	if fileReader.Err() != nil {
		fmt.Println("File read error:", fileReader.Err(), "\n")
		os.Exit(-3)
	}

	// ----- Output file content
	fmt.Println("File content:")
	fmt.Println("-------------")
	for _, n := range names {
		fmt.Println(n.fname, n.lname)
	}

	fmt.Println()
}
