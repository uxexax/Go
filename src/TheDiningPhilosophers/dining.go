/* -----
This is an alternative implementation of the dining philosophers problem, where
eating is controlled by a host. A philosopher is allowed to eat when the following
criteria are fulfilled:
// * The maximum number of parallelly eating philosophers has not yet been reached.
// * The philosopher is not already eating, do not exceed the number of allowed dinners and
//   there's enough time elapsed since the previous dinner.
// * Both chopsticks corresponding to the philosopher's seat are free.
// Note: seats are randomized.
----- */

package main

import (
	dinnerhost "TheDiningPhilosophers/dinnerhost"
	philosopher "TheDiningPhilosophers/philosopher"
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGrp sync.WaitGroup

	names := []string{"John", "Robert", "Lilla", "Charles", "Stella"}
	tableCount := len(names)

	host := dinnerhost.NewDinnerHostPtr(tableCount, 2, 3)
	requestChannel, finishChannel := host.AskChannels()
	go host.Listen()

	for _, name := range names {
		phi := philosopher.NewPhilosopherPtr(name)
		accepted := host.Add(*phi)
		if accepted {
			waitGrp.Add(1)
			go phi.GoToDinner(&waitGrp, requestChannel, finishChannel)
		}
	}

	waitGrp.Wait()

	fmt.Println("===== EVERYBODY LEFT THE TABLE. =====")
	fmt.Println()

	time.Sleep(time.Duration(2) * time.Second)
}
