//	Kevin Chen (2017)
//	Patterns from Pike's Google I/O talk, "Go Concurrency Patterns"

//  Deterministically quit goroutine with quit channel option in select

package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func main() {
	quit := make(chan bool)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Printf("%d: %s %d\n", getGoroutineID(), <-c, i)
	}
	quit <- true
}

func boring(msg string, quit chan bool) <-chan string {
	ch := make(chan string)
	go func() { // anonymous goroutine
		for {
			select {
			case ch <- fmt.Sprintf("%s", msg):
				// nothing
			case <-quit:
				fmt.Printf("%d: Goroutine done\n", getGoroutineID())
				return
			}
		}
	}()
	return ch
}

func getGoroutineID() int {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	var id int
	fmt.Sscanf(string(b), "goroutine %d ", &id)
	return id
}
