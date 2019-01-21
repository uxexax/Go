/* -----
TODO:
 * suppress console printout
 * create a few bigger dungeons
 * add description
 * output file: if there's no way out, show the deepest the traveler can go
 * short paths on console, long in output file
 * output lookout cleanup
----- */

package main

import (
	dungeon "1ddungeon/dungeon"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	D := dungeon.NewDungeon()

	err := D.ObtainMap()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	D.Travel(0, D.FirstSteps(), 0)

	PrintWayOut(D)
	SaveWayOut(D)
}

// PrintWayOut shows a formatted route out of the dungeon (if exists) on the console.
func PrintWayOut(D *dungeon.Dungeon) {
	theresWayOut := (D.TravelResult() == 0)

	if theresWayOut {
		_, shortChart := D.FancyChart()
		_, shortWayOut := D.FancyTravelPath()
		fmt.Println()
		fmt.Println("The way out:")
		fmt.Println(shortChart)
		fmt.Printf(shortWayOut)
		fmt.Println()
	} else {
		fmt.Println()
		fmt.Println("There's no way out!")
		fmt.Println()
	}
}

// SaveWayOut stores the route out of the dungeon (if exists) in a file named similarly to
// the map file, but with a .wayout extension, or named as "#manual.wayout" if the map was
// provided manually.
func SaveWayOut(D *dungeon.Dungeon) {
	theresWayOut := (D.TravelResult() == 0)

	longChart, _ := D.FancyChart()
	longWayOut, _ := D.FancyTravelPath()

	fileName := strings.Split(D.GetSource(), ".")[0]
	fileName += ".wayout"
	fileHandle, _ := os.Create(fileName)
	fileWriter := bufio.NewWriter(fileHandle)

	fileWriter.WriteString(fmt.Sprintln(longChart))

	if theresWayOut {
		fileWriter.WriteString(fmt.Sprintln(longWayOut))
	} else {
		fileWriter.WriteString(fmt.Sprintln("There's no way out!"))
	}

	fileWriter.Flush()
	fileHandle.Close()

	fmt.Println("Way out saved in", fileName)
}
