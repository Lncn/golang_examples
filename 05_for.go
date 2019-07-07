package main

import "fmt"

func main() {
	// "for" is Go's ONLY looping construct?!  (I think I like that :])

	fmt.Println("Using for like a 'while' loop:")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("Using for like a 'for' loop:")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Println("Using for like a 'while (true)' loop:")
	for {
		fmt.Println("\tinside 'for {}' loop and breaking immediately")
		break
	}

	fmt.Println("Inside for's, 'continue' works as expected:")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
