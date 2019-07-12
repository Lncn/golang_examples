package main

import "fmt"

type rect struct {
	width, height int
}

// This is a horrible example for a "pointer receiver type"
// Typically, you want to define a method as a pointer receiver type because
// it is most efficient since Go will by default pass everything by value.
//
// In this case, both of these methods do NOT modify the fields of the struct
// used, so there is no difference.
//
// In my beginner brain, in general, just create methods with a "pointer
// receiver type" by default (like area() here) unless otherwise needed.
// The perim() method here is less efficient since it copies the entire rect
// struct each time the perim() method is called

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
