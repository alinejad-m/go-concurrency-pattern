# Using Channel
### Synchronization
- When the main function executes <–c, it will wait for a value to be sent.
- Similarly, when the boring function executes c <– value, it waits for a receiver to be ready.
- A sender and receiver must both be ready to play their part in the communication. Otherwise, we wait until they are.
- Thus channels both communicate and synchronize. ``