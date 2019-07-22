package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	fmt.Printf("%v\n", p)  // "verb" prints Go value: {1 2}
	fmt.Printf("%+v\n", p) // {x:1 y:2}
	fmt.Printf("%#v\n", p) // main.point{x:1, y:2}

	fmt.Printf("%T\n", p) // "type": main.point

	fmt.Printf("%t\n", true)        // boolean value
	fmt.Printf("%d\n", 123)         // decimal
	fmt.Printf("%b\n", 14)          // binary representation
	fmt.Printf("%c\n", 33)          // char
	fmt.Printf("%x\n", 456)         // hex
	fmt.Printf("%f\n", 78.9)        // float
	fmt.Printf("%e\n", 123400000.0) // exponential
	fmt.Printf("%E\n", 123400000.0) // exponential

	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"") // This includes double quotes and appears "unformatted"
	fmt.Printf("%x\n", "hex this")

	fmt.Printf("%p\n", &p)

	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
