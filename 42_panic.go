package main

import "os"

func main() {
	//panic("A Problamo!")

	_, err := os.Create("/this/directory/doesnt/exist")
	if err != nil {
		panic(err)
	}
}
