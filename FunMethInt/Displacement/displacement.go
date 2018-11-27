/* -----
This program calculates physical displacement based on acceleration, initial velocity,
initial physical displacement and elapsed time provided by the user via standard input.
----- */

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	requiredParams := []string{
		"acceleration",
		"initial velocity",
		"initial displacement",
		"elapsed time"}

	params := AskForInput(requiredParams...)

	distanceTravelled := GenDisplaceFn(
		params["acceleration"],
		params["initial velocity"],
		params["initial displacement"])

	fmt.Println("Distance travelled:", distanceTravelled(params["elapsed time"]))
}

// GenDisplaceFn is a generic displacement function, which returns a displacement function
// specific to acceleration (acc), initial velocity (iVelo) and iDispl (initial displacement)
// The returned function takes one parameter (elapsed time), and returns the total displacement.
func GenDisplaceFn(acc float64, iVelo float64, iDispl float64) func(float64) float64 {
	return func(time float64) float64 {
		displacement := 0.5*acc*math.Pow(time, 2) + iVelo*time + iDispl
		return displacement
	}
}

// AskForInput prompts the user to give the input(s) necessary for the program.
func AskForInput(neededInputs ...string) map[string]float64 {
	scanner := bufio.NewScanner(os.Stdin)
	gatheredValues := make(map[string]float64)

	for _, currentInput := range neededInputs {
		fmt.Printf("Enter %s: ", currentInput)
		scanner.Scan()
		if scanner.Err() != nil {
			fmt.Println("Scan error:", scanner.Err())
			continue
		}
		gatheredValues[currentInput] = ParseInput(scanner.Text())
	}
	return gatheredValues
}

// ParseInput converts an input string to float64, or exits the program if parsing is not
// successful.
func ParseInput(inputString string) float64 {
	parsedValue, err := strconv.ParseFloat(inputString, 5)
	if err != nil {
		fmt.Println("Parse error:", err)
		os.Exit(1)
	}
	return parsedValue
}
