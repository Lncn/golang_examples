package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		// Adding sleep here because I never saw interleaved output without!
		time.Sleep(1)
	}
}

func main() {
	f("direct call")

	go f("goroutine call")

	go func(msg string) {
		fmt.Println(msg)
	}("going!")

	fmt.Scanln()
	fmt.Println("DONE")
}
