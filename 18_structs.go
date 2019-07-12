package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alicia", age: 20})
	fmt.Println(person{name: "Ol' Greg"})
	fmt.Println(&person{name: "Lincoln", age: 32})

	s := person{name: "Sean", age: 5000}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)
}
