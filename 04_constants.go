package main

import (
	"fmt"
	"math"
)

const s string = "conststr"

func main() {
	fmt.Println("Constant string:", s)

	const n = 500000000
	const d = 3e20 / n

	// Constants do not have types and are inferred in their
	// usage context unless explicitly converted
	fmt.Println(d)        // "6e+11"
	fmt.Println(int64(d)) // "600000000000"

	// math.Sin() expects a float64, so 'n' becomes float64 here:
	fmt.Println(math.Sin(n))
}
