# Generator: function that returns a channel
### How it works
- Send a channel on a channel, making goroutine wait its turn.
- Receive all messages, then enable them again by sending on a private channel.
- I use block channel to restore sequencing. In main function, main goroutine waits until message sends to block channels.
- I've indicated the number of goroutines for clarity.