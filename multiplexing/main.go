package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// A channel connects the main and boring goroutines so they can communicate.
func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: You say: %q\n", getGoroutineID(), <-c)
	}
	fmt.Printf("%d: You're boring; I'm leaving.\n", getGoroutineID())
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			input := fmt.Sprintf("%d: %s", getGoroutineID(), <-input1)
			c <- input
		}
	}()
	go func() {
		for {
			input := fmt.Sprintf("%d: %s", getGoroutineID(), <-input2)
			c <- input
		}
	}()
	return c
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
1: I'm out of loop
1: You say: "20: Joe 0"
1: You say: "21: Ann 0"
18: sleep done in loop 0
1: You say: "20: Joe 1"
19: sleep done in loop 0
1: You say: "21: Ann 1"
19: sleep done in loop 1
1: You say: "21: Ann 2"
18: sleep done in loop 1
1: You say: "20: Joe 2"
19: sleep done in loop 2
1: You say: "21: Ann 3"
19: sleep done in loop 3
1: You say: "21: Ann 4"
18: sleep done in loop 2
1: You say: "20: Joe 3"
19: sleep done in loop 4
1: You say: "21: Ann 5"
1: You're boring; I'm leaving.
*/
