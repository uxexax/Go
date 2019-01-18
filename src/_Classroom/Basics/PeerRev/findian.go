package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Println("Input a string: ")
		reader := bufio.NewReader(os.Stdin)
		inputValue, _ := reader.ReadString('\n')
		inputValue = strings.Replace(inputValue, "\n", "", -1)
		inputValue = strings.Replace(inputValue, "\r", "", -1)
		upperValue := strings.ToUpper(inputValue)
		if strings.HasPrefix(upperValue, "I") && strings.HasSuffix(upperValue, "N") && strings.Contains(upperValue, "A") {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}
}