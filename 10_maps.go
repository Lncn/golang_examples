package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len(m):", len(m))

	delete(m, "k2")
	fmt.Println("map (post 'k2' delete):", m)

	// The second optional return value from a map access is a presence flag
	// Use this to disambiguate between missing keys and keys with zero value
	_, prs := m["k2"]
	fmt.Println("presence:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Declared map:", n)
}
