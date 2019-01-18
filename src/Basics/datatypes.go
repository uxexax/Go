package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 100
	var pt *int

	pt = &x

	fmt.Printf("pt = %d\n", pt)
	fmt.Printf("*pt = %d\n", *pt)
	fmt.Printf("&pt = %d\n", &pt)

	fmt.Printf("x = %d\n", x)

	*pt = 101

	fmt.Printf("x after '*pt = 101' = %d\n", x)

	var z int16 = 1
	var y int32 = 2
	//y = z fails

	y = int32(z)

	fmt.Printf("y = %d and z = %d\n", y, z)

	var a int16 = 10000
	fmt.Printf("int8(%d) = %d\n", a, int8(a))

	var s string = "Hello string!"
	s = s + "!!\n"
	fmt.Printf(s)
	//fmt.Printf("%s3.5\n", s)
	fmt.Printf("%d\n", s[1])
	fmt.Printf(string(s[1]) + "\n")

	fmt.Printf(strconv.Itoa(100) + "\n")

	type Weekdays uint8
	const (
		Mon Weekdays = iota
		Tue
		Wed
		Thu
		Fri
		Sat
		Sun
	)

	w1 := Mon
	var w2 Weekdays = Wed

	fmt.Printf("Mon = %d\n", w1)
	fmt.Printf("Wed = %d\n", w2)

	var str *string
	fmt.Scanln(&str)
	fmt.Println(str)

	var f float64 = 968

	if f > 0 {
		fmt.Println("f is greater than zero")
	} else if f == 0 {
		fmt.Println("f is exactly zero")
	} else {
		fmt.Println("f is smaller than zero")
	}

	switch f {
	case 0:
		fmt.Println("ZERO")
	case 5:
		fmt.Println("FIVE")
	case 7:
		fmt.Println("SEVEN")
	case 8:
		fmt.Println("EIGHT")
	default:
		fmt.Println("unnamed")
	}

	switch {
	case f > 1000:
		fmt.Println("f is really big")
	case f > 100:
		fmt.Println("f is big")
	case f > 10:
		fmt.Println("f is average")
	case f > 0:
		fmt.Println("f is small")
	}

	i := 0
	for {
		fmt.Printf("%d", i)
		if i == 5 {
			fmt.Println()
			break
		}
		i++
	}

	for i = 0; i != 70; i = i + 10 {
		fmt.Printf("%d", i)
	}
	fmt.Println()

	var scn string

	n, err := fmt.Scanln(&scn)
	fmt.Printf("n = %d, err = %s\n", n, err)
	fmt.Println("ECHO: " + scn)
}
