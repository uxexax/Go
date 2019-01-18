package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"
import "sort"

func main() {
	mySlice := make([]int, 3)

	var flagx int = 1

	for flagx != 0 {

		// Creating a empty integer slice of size (lenght) 3 before entering the loop.
		// fmt.Println(mySlice)

		fmt.Printf("Enter a Integer or X no exit: ")
		y := bufio.NewScanner(os.Stdin)
		y.Scan()
		x := y.Text()

		if strings.Compare(x, "X") != 0 {
			if z, err := strconv.Atoi(x); err == nil {
				mySlice = append(mySlice, z)
				sort.Sort(sort.IntSlice(mySlice))
				fmt.Println(mySlice)
			}
		} else {
			fmt.Println("EXIT!")
			os.Exit(0)
		}
	}
}
