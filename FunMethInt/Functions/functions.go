package main

import "fmt"

func main() {
	var x int = 10
	var m int = 31
	z := float32(x)

	fmt.Println("Double:", x, double(x))
	fmt.Println("Multiply:", x, multiply(x, m))

	// multiParam := []int{100, 3}
	// fmt.Println("Multiply:", multiParam[0], multiply(multiParam...))

	v, f := split(1000, 5)
	fmt.Println("Slice:", v, f)

	halve(&z)
	fmt.Println("Half:", z)

	a := [3]int{100, 200, 300}
	s := []int{100, 200, 300}

	fmt.Println("First by name:", firstByName(a))

	doubleByReference(&a)
	fmt.Println("First by reference:", a)

	doubleFirstWithSlice(s)
	fmt.Println("Double first with slice:", s)

	fmt.Println("First with slice:", firstWithSlice(s))

	doubleFirstWithSlice(a[:])
	fmt.Println("Double first with slice:", a, "Arr2Sli")
}

// ----- By value.
func double(val int) int {
	return val * 2
}

func multiply(val int, mul int) int {
	return val * mul
}

func split(val float32, factor float32) (float32, float32) {
	return val / factor, val - val/factor
}

// ----- By reference.
func halve(val *float32) {
	*val = *val / 2
}

// ----- Passing arrays.
func firstByName(arr [3]int) int {
	return arr[0]
}

func doubleByReference(arr *[3]int) {
	for index, _ := range *arr {
		(*arr)[index] = (*arr)[index] * 2
	}
}

func doubleFirstWithSlice(sli []int) {
	sli[0] *= 2
}

func firstWithSlice(sli []int) int {
	return sli[0]
}
