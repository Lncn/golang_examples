package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "starting job", j)
		time.Sleep(time.Second)
		res := j * 2
		fmt.Println("Worker", id, "finished job", j)
		results <- res
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 8; w++ {
		go worker(w, jobs, results)
	}

	for j := 0; j < 20; j++ {
		jobs <- j
	}
	close(jobs)
	fmt.Println("Main finished producing jobs!")

	for a := 0; a < 20; a++ {
		r := <-results
		fmt.Println("Result out:", r)
	}
	close(results)
}
