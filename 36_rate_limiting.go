package main

import (
	"fmt"
	"time"
)

func main() {
	// Lets serve some requests!
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("Request", req, time.Now())
	}

	// This is funny looking...
	// Lets say we want to guarantee an overall rate limit, but allow
	// for "bursty" periods of events. We can do this with a Time channel
	// with a depth of >1. For example, this `burstyLimiter` will allow
	// up to 3 events to be buffered for handling
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("Req", req, time.Now())
	}
}
