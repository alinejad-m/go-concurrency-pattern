package main

import (
	"fmt"
	"runtime"
)

func main() {
	const n = 5
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(i, left, right)
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Printf("go:%d final value is:%d\n", getGoroutineID(), <-leftmost)
}

func f(i int, left, right chan int) {
	fmt.Printf("go:%d\tloop:%d\n", getGoroutineID(), i)
	rightNumber := <-right
	fmt.Printf("go:%d\tI get right chan number:%d\n", getGoroutineID(), rightNumber)
	left <- 1 + rightNumber
}

func getGoroutineID() int {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	var id int
	fmt.Sscanf(string(b), "goroutine %d ", &id)
	return id
}
