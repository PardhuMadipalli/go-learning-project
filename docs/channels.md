# Channels

## Unbuffered channels
- These channels don't have any buffer memory to keep the received messages.
- So the receiver will be blocked until the sender sends a message.
- This makes the program synchronous. Until the goroutine sends a message, the caller function won't proceed.

Example: [unbufferedChannelDemo.go](../main/unbufferedChannelDemo.go)