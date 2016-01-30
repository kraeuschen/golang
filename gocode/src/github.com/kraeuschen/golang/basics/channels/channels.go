package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and
// receive those values into another goroutine.
func main() {
	// new channel with make
	messages := make(chan string)

	// send a value into a channel using the channel <- syntax
	go func() { messages <- "ping" }()

	// <-channel syntax receives a value from an channel
	msg := <-messages
	fmt.Println(msg)
}
