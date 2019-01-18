package philosopher

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Philosopher represents an individual philosopher.
type Philosopher struct {
	name        string
	respChannel chan string
}

// NewPhilosopherPtr creates and initializes a Philosopher object and returns a pointer to it.
func NewPhilosopherPtr(name string) *Philosopher {
	phi := new(Philosopher)
	phi.Init(name)
	return phi
}

// Init can be used to initialize a Philosopher.
func (phi *Philosopher) Init(name string) {
	phi.name = name
	phi.respChannel = make(chan string)
}

// Name returns the name of the philosopher.
func (phi *Philosopher) Name() string {
	return phi.name
}

// RespChannel returns the philosopher's dedicated response channel.
func (phi *Philosopher) RespChannel() chan string {
	return phi.respChannel
}

// GoToDinner is the philosopher's main task. They periodically issue eat requests to the host, unless
// not already eating. When asked so by the host, the philosopher leaves.
func (phi *Philosopher) GoToDinner(waitGrp *sync.WaitGroup, requestChannel, finishChannel chan string) {
	defer waitGrp.Done()

	retryInterval := time.Duration(2000) * time.Millisecond
	eatingDuration := time.Duration(5000) * time.Millisecond

	for {
		requestChannel <- phi.name
		switch <-phi.respChannel {
		case "OK":
			time.Sleep(eatingDuration)
			finishChannel <- phi.name
		case "E:LIMIT":
			fmt.Println(strings.ToUpper("----- " + phi.name + " LEFT THE TABLE. -----"))
			fmt.Println()
			return
		default:
			time.Sleep(retryInterval)
		}
	}
}
