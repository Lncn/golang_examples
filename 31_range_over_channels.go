package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// We can also use the `range` keyword to iterate over channel values
	for elem := range queue {
		fmt.Println(elem)
	}
	// Also note, despite the `queue` being `close`d,
	// the elements are still able to be received.
}
