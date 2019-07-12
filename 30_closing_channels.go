package main

import "fmt"

func main() {
	jobs := make(chan int, 1)
	godone := make(chan bool)

	go func() {
		for {
			// The two value version of returns `false` in
			// `more` if the `jobs` channel is closed
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				godone <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		fmt.Println("Sending job", j)
		jobs <- j
	}
	close(jobs)
	fmt.Println("Sent all jobs!")

	<-godone
}
