package main

import (
	"fmt"
	"runtime"
	"time"
)

type Message struct {
	str   string
	block chan int
}

func main() {
	ch := fanIn(generator("Hello"), generator("Bye"))
	for i := 0; i < 10; i++ {
		msg1 := <-ch
		fmt.Printf("g%d: %s\n", getGoroutineID(), msg1.str)

		msg2 := <-ch
		fmt.Printf("g%d: %s\n", getGoroutineID(), msg2.str)

		<-msg1.block // reset channel, stop blocking
		<-msg2.block
	}
}

// fanIn is itself a generator
func fanIn(ch1, ch2 <-chan Message) <-chan Message { // receives two read-only channels
	newCh := make(chan Message)
	go func() {
		for {
			msg := <-ch1
			msg.str = fmt.Sprintf("g%d: %s", getGoroutineID(), msg.str)
			newCh <- msg
		}
	}() // launch two goroutine while loops to continuously pipe to new channel
	go func() {
		for {
			msg := <-ch2
			msg.str = fmt.Sprintf("g%d: %s", getGoroutineID(), msg.str)
			newCh <- msg
		}
	}()
	time.After(time.Second)
	return newCh
}

func generator(msg string) <-chan Message { // returns receive-only channel
	ch := make(chan Message)
	blockingStep := make(chan int) // channel within channel to control exec, set false default
	go func() {                    // anonymous goroutine
		for i := 0; ; i++ {
			ch <- Message{fmt.Sprintf("g%d: %s %d", getGoroutineID(), msg, i), blockingStep}
			time.Sleep(time.Second)
			blockingStep <- 1 // block by waiting for input
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
