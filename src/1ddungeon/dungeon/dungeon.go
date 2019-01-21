package dungeon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Dungeon is the aggregator object for the map and the travel path.
type Dungeon struct {
	chart      []int
	travelPath []int
	travelRes  int
	source     string
}

// NewDungeon instantiates a new dungeon.
func NewDungeon() *Dungeon {
	dungeon := new(Dungeon)
	dungeon.Init(100)
	return dungeon
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

func (dungeon *Dungeon) GetSource() string {
	return dungeon.source
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

func (dungeon *Dungeon) FancyChart() (string, string) {
	separatorMark := " "

	fancyChart := ""
	shortFancyChart := ""
	for _, chartValue := range dungeon.chart {
		fancyChart += strconv.Itoa(chartValue) + separatorMark
	}
	fancyChart = fancyChart[:len(fancyChart)-1]

	shortFancyChart = fancyChart // TODO

	return fancyChart, shortFancyChart
}

func (dungeon *Dungeon) FancyTravelPath() (string, string) {
	// inAndOutMark := "==>"
	transitMark := "="
	stopMark := "O"
	separatorMark := " "
	travelPathIndex := 0

	fancyPath := ""
	fancyShortPath := ""
	for chartIndex, chartValue := range dungeon.chart {
		markSize := len(strconv.Itoa(chartValue))
		mark := transitMark
		if chartIndex == dungeon.travelPath[travelPathIndex] {
			mark = stopMark
			travelPathIndex++
		}

		for m := 0; m < markSize; m++ {
			fancyPath += mark
		}
		fancyPath += separatorMark
	}

	fancyShortPath = fancyPath // TODO

	return fancyPath, fancyShortPath
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

// ObtainMap takes dungeon map points either from a given file, or from the
// standard input. The map should be a single line of space separated integers.
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
	source := fileName

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

		fileHandle.Close()

		if fileReader.Err() != nil {
			return fileReader.Err()
		}
	}

	dungeon.StoreSource(source)

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
