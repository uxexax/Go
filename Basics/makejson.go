package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var addrBook map[string]string
	addrBook = make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)

	readValues := [2]string{"name", "address"}

	fmt.Println()
	for _, v := range readValues {
		fmt.Printf("Enter the %s: ", v)
		scanner.Scan()

		if scanner.Err() != nil {
			fmt.Println("Leaving: could not scan the %s.", v)
			os.Exit(-1)
		}

		addrBook[v] = scanner.Text()
	}

	jsonBytes, encError := json.Marshal(addrBook)

	if encError == nil {
		fmt.Println()
		fmt.Println("----- The JSON object as byte array:")
		fmt.Println(jsonBytes)
		fmt.Println("----- The JSON object as human readable string:")
		fmt.Println(string(jsonBytes))
		fmt.Println()
	} else {
		fmt.Println("Marshalling failed:", encError)
		os.Exit(-2)
	}
}
