# Quit channel: 
## We can turn this around and tell Joe to stop when we're tired of listening to him.
### How it works
- In the boring function, an infinite loop containing a select statement is utilized. Whenever the ch channel is prepared to transmit a message, the first action is executed. Conversely, if the quit channel is prepared, the function retrieves the message from the quit channel and subsequently exits.
- I've indicated the number of goroutines for clarity.