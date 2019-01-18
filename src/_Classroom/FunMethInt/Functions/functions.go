package main

import (
	"fmt"
	"math"
)

func main() {
	i := 1
	defer fmt.Printf("Bye! (%d)\n", i)
	i += 1

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

	// ----- Functions as variables.
	var F func(int) int

	F = func(x int) int {
		return x * 10
	}

	fmt.Println("Function definition:", F(2))

	var F2 func(int) int = func(x int) int {
		return x * 100
	}

	fmt.Println("Function definition 2:", F2(25))

	F3 := func(x int) int {
		return x * 33
	}

	fmt.Println("Function definition 3:", F3(33))

	// ----- Functions as parameters.
	dbl := func(x float32) float32 { return x * 2 }
	hlv := func(x float32) float32 { return x / 2 }
	apply := func(f func(float32) float32, x float32) float32 {
		return f(x)
	}

	fmt.Println("Function as parameter 1:", apply(dbl, 200))
	fmt.Println("Function as parameter 2:", apply(hlv, 200))

	add := func(x float64, y float64) float64 {
		return (x + y)
	}

	sqFun := func(p1, p2 float64, f func(float64, float64) float64) float64 {
		result := math.Pow(f(p1, p2), 2)
		return result
	}

	fmt.Println("Function as parameter 3:", sqFun(3, 4, add))

	// ----- Function in function.
	fiver := func(x int) int {
		halve := func(x int) int {
			return x / 2
		}
		x *= 10
		x = halve(x)
		return x
	}

	fmt.Println("Function in function:", fiver(10))

	// ----- Anonymous function.
	fmt.Println("Anonymous function 1:", func() float64 { return (3) }())

	y := 123.11
	fmt.Println("Anonymous function 2:", func() float64 { return (y / 3) }())
	fmt.Println("Anonymous function 3:", func(p float64) float64 { return (p / 9) }(y))

	// ----- Function returned. Below works because of lexical scoping?
	multiplier := func(m float64) func(float64) float64 {
		return func(x float64) float64 {
			return m * x
		}
	}

	triple := multiplier(3)
	sixer := multiplier(6)
	fiftier := multiplier(50)

	fmt.Println("Function returned:", triple(55), sixer(55), fiftier(55))

	// ----- Variadic call.
	fmt.Println("Variadic:", sum(101.44, 205.1, 123))

	V := []float64{229.53, 441.2, 110, 3000.55}
	fmt.Println("Variadic: the sum of", V, "is", sum(V...))
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
	for index := range *arr {
		(*arr)[index] = (*arr)[index] * 2
	}
}

func doubleFirstWithSlice(sli []int) {
	sli[0] *= 2
}

func firstWithSlice(sli []int) int {
	return sli[0]
}

// ----- Variadic function.
func sum(numbers ...float64) float64 {
	// sumOfNumbers := numbers[0]
	// if len(numbers) > 1 {
	// 	sumOfNumbers += sum(numbers[1:]...)
	// }
	// return sumOfNumbers

	if len(numbers) <= 1 {
		return numbers[0]
	}
	return (numbers[0] + sum(numbers[1:]...))
}
