# Generator: function that returns a channel
### How it works
- In the main function, there's the main goroutine. We call the 'boring' function twice within the main function, resulting in two goroutines for each call. Additionally, within the 'fanIn' function, two goroutines are responsible for fetching input values from the input channel and forwarding them to the 'c' channel. Subsequently, the main goroutine receives each message individually.
- I've indicated the number of goroutines for clarity.