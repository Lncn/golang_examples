package main

import "fmt"

func main() {
	fmt.Println("Looks like if/else statements work as expected in Go...")
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	fmt.Println("Caveats:\n(1) No parenthesis around condition!\n(2) Curly braces are REQUIRED!")
	fmt.Println("(3) Also, statements can precede conditional and variables are accessible in all branches")
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	fmt.Println("Also of note, there is NO ternary if '?' operator in Go! (This appears to be one complaint I've seen)")
}
