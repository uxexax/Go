/* ----------
This program sorts a sequence of integers provided by the user on the standard input
in the following way:
 * the input is parsed into an integer slice
 * this slice is split into _numOfPieces_ number of equally or almost equally same
   length pieces
 * each piece is sorted in an individual goroutine using the bubble sorting algorithm
 * sync.WaitGroup is used for suspending the main thread until all goroutines finish their task
 * when all sorting is finished, the pieces are merged into one sorted sequence

The implementation is an extension of an earlier bubble sort algorithm project.
 ---------- */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var sortWait sync.WaitGroup
	numOfPieces := 4

	inputStr, inputErr := AskForInput()
	if inputErr != nil {
		fmt.Println("Error:", inputErr)
		fmt.Println()
		os.Exit(1)
	}

	sequenceOfInts, parseError := ParseInput(inputStr)
	if parseError != nil {
		fmt.Println("Error:", parseError)
		fmt.Println()
		os.Exit(2)
	}

	sequencePieces := SplitSlice(sequenceOfInts, numOfPieces)

	for index, piece := range sequencePieces {
		sortWait.Add(1)
		go BubbleSort(piece, &sortWait, index)
	}

	sortWait.Wait()
	sortedSequence := MergePieces(sequencePieces)

	fmt.Println()
	fmt.Println("The sorted sequence is:")
	fmt.Println(sortedSequence)
	fmt.Println()
}

// MergePieces takes a slice of ordered slice pieces and returns a single merged and
// ordered slice. _pieces_ can contain any number of slice pieces, although performance
// has not been verified for this algorithm. The idea is the following. Imagine the pieces
// are stacked on top of each other, aligned at index 0. In each round the function obtains
// a vertical slice of the stacked pieces, takes the minimum value, appends it to the
// merged slice and increases the corresponding piece's index (_currentIndices_ is used
// to track the indexes of pieces).
func MergePieces(pieces [][]int) []int {
	totalLength := 0
	for _, piece := range pieces {
		totalLength += len(piece)
	}
	slice := make([]int, 0, totalLength)

	currentIndices := make([]int, len(pieces))
	vertical := GetVertical(pieces, currentIndices)
	for len(vertical) > 0 {
		minInd, minVal := GetMin(vertical)
		slice = append(slice, minVal)
		currentIndices[minInd]++
		vertical = GetVertical(pieces, currentIndices)
	}

	return slice
}

// GetVertical returns a vertical slice of slice _pieces_ based on _currentIndices_.
// Indices out of range are ignored, and thus a vertical does not contain elements
// from pieces whose all elements hava already been put into the merged slice.
func GetVertical(pieces [][]int, currentIndices []int) map[int]int {
	vertical := make(map[int]int)
	for i, v := range currentIndices {
		if v < len(pieces[i]) {
			vertical[i] = pieces[i][v]
		}
	}
	return vertical
}

// GetMin returns the index and the value of the first hash element holding a minimum value.
func GetMin(vertical map[int]int) (int, int) {
	if len(vertical) == 0 {
		return -1, -1
	}

	minIndex := 0
	minValue := 0
	firstIter := true
	for i, v := range vertical {
		if firstIter || (v < minValue) {
			minIndex = i
			minValue = v
			firstIter = false
		}
	}

	return minIndex, minValue
}

// SplitSlice splits a slice into _numOfPieces_ number of equal or almost equal length
// smaller slices. (If the original slice cannot be split into equal length pieces, then
// the first pieces get one extra item until the number of remainders becomes zero.)
func SplitSlice(slice []int, numOfPieces int) [][]int {
	pieces := make([][]int, numOfPieces)

	if len(slice) < numOfPieces {
		return pieces
	}

	itemsPerPiece := len(slice) / numOfPieces
	remainder := len(slice) % numOfPieces
	pieceBegin := 0
	pieceEnd := itemsPerPiece

	for i := range pieces {
		// Remainder elements are distributed one per piece, until none is left.
		if remainder > 0 {
			pieceEnd++
			remainder--
		}

		// Copies are made to avoid unintentional change of original slice.
		pieces[i] = make([]int, len(slice[pieceBegin:pieceEnd]))
		copy(pieces[i], slice[pieceBegin:pieceEnd])
		pieceBegin = pieceEnd
		pieceEnd += itemsPerPiece
	}

	return pieces
}

// BubbleSort implements the bubble sort algorithm for integers.
func BubbleSort(sequence []int, wg *sync.WaitGroup, id int) {
	fmt.Println("Goroutine no.", id, "sorts", sequence)

	for upperIndex := len(sequence) - 1; upperIndex > 0; upperIndex-- {
		for runningIndex := 0; runningIndex < upperIndex; runningIndex++ {
			if sequence[runningIndex] > sequence[runningIndex+1] {
				Swap(sequence, runningIndex)
			}
		}
	}
	wg.Done()
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

	for _, v := range splitLine {
		intValue, convError := strconv.Atoi(v)
		if convError != nil {
			return nil, convError
		}

		parsedSequence = append(parsedSequence, intValue)
	}

	return parsedSequence, nil
}
