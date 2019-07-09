package main

import "fmt"

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(nextInt())
	fmt.Println(newInts())
	fmt.Println(nextInt())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
