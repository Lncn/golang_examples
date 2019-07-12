package main

import "fmt"
import "time"

func main() {
	messages := make(chan string)

	go func() {
		time.Sleep(time.Second)
		messages <- "ping!"
	}()

	// This is a blocking call by default which waits for something to be put
	// into the `messages` buffer for the assignment to be complete
	msg := <-messages
	fmt.Println(msg)
}
