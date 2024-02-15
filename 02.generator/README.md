# Generator
## function that returns a channel
### How it works
- In the main function, there's the main goroutine. Within the boring function, there's another goroutine. This one sends a message to the channel and waits until the main goroutine receives it. If not, the main goroutine loops, waiting until the boring goroutine sends a message through the channel. Please review the output of go run main.go.
- I've indicated the number of goroutines for clarity.