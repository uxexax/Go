package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	initV := 0
	v := initV

	wg.Add(2)
	go ChangeLoop(&v, 50, 3, 100, &wg)
	go ChangeLoop(&v, -30, 1, 400, &wg)
	wg.Wait()

	fmt.Printf("Initial value: %d |Final value: %d", initV, v)
	fmt.Println()
}

// ChangeLoop changes the _value_ by _changeRate_ at regular intervals defined by _delayMs_.
func ChangeLoop(value *int, changeRate, delayMs, rounds int, wg *sync.WaitGroup) {
	for i := 0; i < rounds; i++ {
		time.Sleep(time.Duration(delayMs) * time.Millisecond)
		*value += changeRate
	}
	wg.Done()
}
