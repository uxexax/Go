package main

import (
	"fmt"
	"strconv"
)

func main() {

	// ----- ARRAYS
	fmt.Println("----- ARRAYS")

	var a [3]int
	a = [...]int{5, 6, 7}
	fmt.Printf("a = %d\n", a)

	var b [3]int = [3]int{1, 5, 7}
	fmt.Printf("b = %d\n", b)

	var c = [3]int{3, 2, 1}
	fmt.Printf("c = %d\n", c)

	d := [...]int{9, 2, 9}
	fmt.Printf("d = %d\n", d)

	for index, item := range d {
		fmt.Printf("index: %d, item: %d\n", index, item)
	}

	for index := 0; index < len(d); index++ {
		fmt.Printf("index: %d, item: %d\n", index, d[index])
	}

	str := "Hejjejejjj"
	for index, item := range str {
		fmt.Print(strconv.Itoa(index))
		fmt.Print(item)
	}
	fmt.Println()

	for index := 0; index < len(str); index++ {
		fmt.Printf("%d%d", index, str[index])
	}
	fmt.Println()

	// ----- SLICES
	fmt.Println("----- SLICES")

	y := []int{1, 4, 9, 18, 100, 200}

	fmt.Println(y)

	slice := make([]int, 3, 10)
	fmt.Println(slice)
	slice[0] = 1
	fmt.Println(slice)
	copy(slice[1:3], []int{2, 2})
	fmt.Println(slice)

	app := []int{3, 4, 5}
	slice = append(slice, 8, 8, 8, 8)
	slice = append(slice, app...)
	fmt.Println(slice)
	fmt.Println(slice[:])
	fmt.Println(slice[2:])
	fmt.Println(slice[:6])

	// ----- HASH TABLES
	fmt.Println("----- HASH TABLES")

	var hash map[string]int
	hash = make(map[string]int)

	// hash := map[string]int{"alpha": 10, "beta": 20}

	fmt.Println(hash)

	hash["alpha"] = 10
	hash["beta"] = 20

	fmt.Println(hash)

	hash["gamma"] = 35
	delete(hash, "alpha")

	fmt.Println(hash)

	fmt.Println(hash["alpha"])

	value, exists := hash["alpha"]
	fmt.Println(value, exists)

	value, exists = hash["gamma"]
	fmt.Println(value, exists)

	for k, v := range hash {
		fmt.Println(k, v)
	}

	// ----- STRUCTS
	fmt.Println("----- STRUCTS")

	type location struct {
		x           int
		y           int
		description string
	}

	L := new(location)
	L.x = 122
	L.y = 200
	L.description = "First location"
	fmt.Println(L)
	fmt.Println(*L)
	(*L).x = 123
	fmt.Println(L)
	fmt.Println(*L)

	L2 := location{x: 100, y: 200, description: "Second location"}
	fmt.Println(L2)

	var L3 location
	L3 = location{x: 211, y: 140, description: "Third location"}
	fmt.Println(L3)

	L3.y = 142
	fmt.Println(L3)
}
