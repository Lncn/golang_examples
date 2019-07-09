package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func addtwice(x, y, z int) int {
	return x + y + z
}

func main() {
	res := add(1, 2)
	fmt.Println("add(1, 2) returned", res)

	res = addtwice(1, 2, 3)
	fmt.Println("addtwice(1, 2, 3) returned", res)
}
