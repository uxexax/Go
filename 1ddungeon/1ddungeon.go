/* -----
TODO:
 * put things into a package
 * create out put file
 * suppress console printout
 * create a few bigger dungeons
 * add description
 * src, bin etc. directories
----- */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dungeon := NewDungeon()

	err := dungeon.ObtainMap()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(dungeon.chart)

	dungeon.Travel(0, dungeon.FirstSteps(), 0)

	if dungeon.TravelResult() == 0 {
		dungeon.PrintPath()
	} else {
		fmt.Println("There's no way out!")
	}
}

// Dungeon is the aggregator object for the map and the travel path.
type Dungeon struct {
	chart      []int
	travelPath []int
	travelRes  int
	source     string
}

// NewDungeon instantiates a new dungeon.
func NewDungeon() Dungeon {
	var d Dungeon
	d.Init(100)
	return d
}

// Init initializes the dugeon.
func (dungeon *Dungeon) Init(mapSize int) {
	dungeon.chart = make([]int, 0, mapSize)
	dungeon.travelPath = make([]int, 0, mapSize/10)
	dungeon.travelRes = -1
	dungeon.source = "#undef"
}

// StoreMap is used to add a chart to the Dungeon object. It ereases an existing travel path!
func (dungeon *Dungeon) StoreMap(chart []int) {
	dungeon.chart = chart
	dungeon.travelPath = make([]int, 0, len(chart))
}

// StoreSource adds the source of the map to the Dungeon object.
func (dungeon *Dungeon) StoreSource(source string) {
	dungeon.source = source
}

// FirstSteps returns the number of maximum possible steps indicated in the first place
// of the dungeon.
func (dungeon *Dungeon) FirstSteps() int {
	return dungeon.chart[0]
}

// TravelResult returns the result of the travel:
// ~  0:  There's a way out.
// ~ -1:  There's no way out.
func (dungeon *Dungeon) TravelResult() int {
	return dungeon.travelRes
}

// TravelPath returns the way out of the dungeon.
func (dungeon *Dungeon) TravelPath() []int {
	return dungeon.travelPath
}

// Travel recursively advances on the dungeon map. It tries the new position indicated in the
// current position. If there's a dragon there, it tries the new-1 position, and so on. Return
// value 0 is propagated back if the exit can be reached.
func (dungeon *Dungeon) Travel(currentPos, stepsToTake, level int) {
	newPos := currentPos + stepsToTake

	if level >= len(dungeon.travelPath) {
		dungeon.travelPath = append(dungeon.travelPath, 0)
	}

	switch {
	case newPos == currentPos:
		dungeon.travelRes = -1
		return
	case newPos >= len(dungeon.chart):
		dungeon.travelRes = 0
		return
	}

	dungeon.travelRes = -1
	for i := 0; i < stepsToTake; i++ {
		dungeon.Travel(newPos, dungeon.chart[newPos], level+1)
		if dungeon.travelRes == 0 {
			dungeon.travelPath[level+1] = newPos
			break
		}
		newPos = newPos + dungeon.travelRes
	}

	return
}

// PrintPath prints the path to the exit to the standard output. Elements of the path are
// the indices of the related points of the dungeon map.
func (dungeon *Dungeon) PrintPath() {
	travelPathStr := make([]string, len(dungeon.travelPath))
	for i, v := range dungeon.travelPath {
		travelPathStr[i] = strconv.Itoa(v)
	}

	fmt.Printf("%s, out", strings.Join(travelPathStr, ", "))
	fmt.Println()
}

// ObtainMap takes dungeon map points one by one and one per line from the standard
// input. Empty line indicates end of input.
func (dungeon *Dungeon) ObtainMap() error {
	consoleLoop := func() (string, error) {
		promptReader := bufio.NewScanner(os.Stdin)
		endLoop := false
		for !endLoop {
			fmt.Print(">> ")
			promptReader.Scan()
			endLoop = (promptReader.Text() != "")
		}
		return promptReader.Text(), promptReader.Err()
	}

	fmt.Println()
	fmt.Println("Enter the file name of a map, or #manual to provide the map on the console.")
	userInput, userError := consoleLoop()

	if userError != nil {
		return userError
	}
	fileName := userInput

	rawMap := ""
	if strings.ToLower(fileName) == "#manual" {
		userInput, userError = consoleLoop()

		if userError != nil {
			return userError
		}

		rawMap = userInput
	} else {
		fileHandle, openErr := os.Open(fileName)
		if openErr != nil {
			return openErr
		}

		fileReader := bufio.NewScanner(fileHandle)

		for fileReader.Scan() {
			rawMap += " " + fileReader.Text()
		}
		fmt.Println(rawMap)

		if fileReader.Err() != nil {
			return fileReader.Err()
		}
	}

	return dungeon.parseMap(rawMap)
}

// parseMap takes the raw dungeon map string obtained from the input,
// parses it to a slice of integers, and stores it in the Dungeon object.
func (dungeon *Dungeon) parseMap(rawMap string) error {
	splitString := strings.Fields(rawMap)

	for _, v := range splitString {
		intValue, convError := strconv.Atoi(v)
		if convError != nil {
			return convError
		}

		dungeon.chart = append(dungeon.chart, intValue)
	}

	return nil
}
