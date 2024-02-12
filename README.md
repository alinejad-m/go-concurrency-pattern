# Golang Concurrency + Patterns

*I've just documented my journey of learning Golang concurrency, outlining the key concepts and insights gained along the way.*
***
# Definition
## Main Concepts
### Concurrency
**Concurrency** in Go refers to the ability to perform multiple tasks simultaneously, using goroutines and tools like WaitGroups and channels to synchronize and communicate between them.
### Goroutine
**Goroutine** is an independently executing function, launched by go statement.
### Channel
A **channel** in Go provides a connection between two goroutines, allowing them to communicate.
### Synchronization
- When the main function executes <–c, it will wait for a value to be sent.
- Similarly, when the boring function executes c <– value, it waits for a receiver to be ready.
- A sender and receiver must both be ready to play their part in the communication. Otherwise we wait until they are.
- Thus channels both communicate and synchronize. 