package main

import (
	"fmt"
	"time"
)

func pingf(pings chan<- string, msg string) {
	time.Sleep(time.Second)
	fmt.Println("pings <-", msg)
	pings <- msg
}

func pongf(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	fmt.Println(msg, "<- pings")
	time.Sleep(time.Second)
	fmt.Println("pongs <-", msg)
	pongs <- msg
}

func main() {
	ping := make(chan string, 1)
	pong := make(chan string, 1)

	pingf(ping, "Here's a passed message!")
	pongf(ping, pong)
	fmt.Println(<-pong)
}
