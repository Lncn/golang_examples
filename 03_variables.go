package main

import "fmt"

func main() {
	// Create a variable using the keyword 'var'
	var a = "initial"
	fmt.Println(a)

	// You can declare and assign multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go also supports type inference!
	var d = true
	fmt.Println(d)

	// Variable declarations without initialization are 'zero-valued'
	// by default. (Question... to what extend does this extend to Go's
	// user defined non-builtin struct/class types?)
	var e int
	fmt.Println(e)

	// Go also supports a ':=' shorthand for declaring and initializing
	f := "apple"
	fmt.Println(f)

	// Of course I can also reassign values, right?
	f = "f, reassigned!"
	fmt.Println(f)
}
