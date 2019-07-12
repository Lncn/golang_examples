package main

import (
	"fmt"
	"time"
)

// This is a worker channel that takes a "synchronizatino channel" as input
// to notify a waiting task ("<-done") when the worker is done.
func worker(done chan bool) {
	fmt.Print("Working... ")
	time.Sleep(time.Second)
	fmt.Println("DONE")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}
