package main

import "fmt"

func main() {
	// Arrays in Go are fixed lengthed and defined by their type declaration:
	//  [n]T = Array of T's of length n
	// Slices are declared similarly, but a length is not specified
	//  []T = Slice of T's
	//
	// You can create a slice, with initial values like this:
	//  s := []string{"a", "b", "c"}
	// but more commonly, you'd use the builtin 'make'
	s := make([]string, 3)
	fmt.Println("empty slice:", s)

	s[0] = "a"
	s[1] = "bcd"
	s[2] = "e"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "f")
	s = append(s, "g", "h")

	fmt.Println("appended to:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy of s, c:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)
	l = s[:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"i", "j", "k"}
	fmt.Println("Declared t:", t)

	const rows = 4
	twoD := make([][]int, rows)
	for row := 0; row < rows; row++ {
		innerLen := row + 1
		twoD[row] = make([]int, innerLen)
		for col := 0; col < innerLen; col++ {
			twoD[row][col] = (row + 1) * (col + 1)
		}
	}
	fmt.Println("twoD:", twoD)
}
