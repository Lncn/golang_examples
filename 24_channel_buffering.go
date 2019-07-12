package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 3)

	messages <- "buffered"
	// This next messages put would cause a deadlock if our `messages` channel
	// was not large enough (e.g. if we set the depth to 1)
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
