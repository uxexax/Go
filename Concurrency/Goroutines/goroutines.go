/* -----
A variable is changed with the same value at regular intervals while
it is also sampled at regular intervals. This is done in multiple rounds,
with the same changing and sampling intervals in every round. As a result
of the race condition caused by simultaneously writing and reading the
same variable from different goroutines, the sampled values cannot be
predicted and the samples of subsequent rounds will be different.
----- */

package main

import (
	"fmt"
	"time"
)

func main() {
	initV := 0
	numOfSamples := 4
	changingInterval := 10
	samplingInterval := 10

	for round := 0; round < 21; round++ {
		v := initV
		samples := make([]int, numOfSamples)

		go ChangeLoop(&v, 50, changingInterval)
		go TakeSamples(&v, &samples, numOfSamples, samplingInterval)

		time.Sleep(time.Duration(100) * time.Millisecond)

		fmt.Printf("Round %2d: |Init: %d |Final: %d |Samples:", round+1, initV, v)
		for _, v := range samples {
			fmt.Printf("%4d", v)
		}
		fmt.Println()
	}
}

// ChangeLoop changes the _value_ by _changeRate_ at regular intervals defined by _delayMs_.
func ChangeLoop(value *int, changeRate int, delayMs int) {
	for i := 0; i < 4; i++ {
		time.Sleep(time.Duration(delayMs) * time.Millisecond)
		*value += changeRate
	}
}

// TakeSamples captures the value of _subject_ periodically at _intervalMs_ milliseconds.
// The function takes a total of _numOfSamples_ samples and returns them in _samples_.
func TakeSamples(subject *int, samples *[]int, numOfSamples, intervalMs int) {
	for i := range *samples {
		time.Sleep(time.Duration(intervalMs) * time.Millisecond)
		(*samples)[i] = *subject
	}
}
