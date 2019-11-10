package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Print(rand.Intn(100), ".", rand.Intn(100))
	fmt.Println()
	fmt.Println(rand.Float32())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64() * 1000000)

	/* rand is deterministic by default, so you need to seed (or provide "NewSource")
	   if you want different behavior each run of the program */
	rand.Seed(time.Now().UnixNano())
	fmt.Print(rand.Intn(100), ".", rand.Intn(100))
	fmt.Println()
	fmt.Println(rand.Float32())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64() * 1000000)

	/* You can create multiple random number generators as well.
	   As you might expect, seeding with the same value will produce the same rand sequence */
	source := rand.NewSource(time.Now().UnixNano())
	rand1 := rand.New(source)
	fmt.Print(rand1.Intn(1000), ".", rand1.Intn(10000))
	fmt.Println()

	rand2 := rand.New(rand.NewSource(42))
	rand3 := rand.New(rand.NewSource(42))
	fmt.Print(rand2.Intn(1000), ".", rand2.Intn(10000))
	fmt.Println()
	fmt.Print(rand3.Intn(1000), ".", rand3.Intn(10000))
	fmt.Println()
}
