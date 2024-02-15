package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// A channel connects the main and boring goroutines, so they can communicate.
func main() {
	c := boring("boring!") // Function returning a channel.
	for i := 0; i < 5; i++ {
		fmt.Printf("%d: You say: %q\n", getGoroutineID(), <-c)
	}
	fmt.Printf("%d: You're boring; I'm leaving.\n", getGoroutineID())
}

func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			fmt.Printf("%d: sleep done in loop %d\n", getGoroutineID(), i)
		}
	}()
	fmt.Printf("%d: I'm out of loop\n", getGoroutineID())
	return c // Return the channel to the caller.
}

func getGoroutineID() int {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	var id int
	fmt.Sscanf(string(b), "goroutine %d ", &id)
	return id
}

/* result: go run ./main.go
1: I'm out of loop
1: You say: "boring! 0"
6: sleep done in loop 0
1: You say: "boring! 1"
6: sleep done in loop 1
1: You say: "boring! 2"
6: sleep done in loop 2
1: You say: "boring! 3"
6: sleep done in loop 3
1: You say: "boring! 4"
1: You're boring; I'm leaving */
