package main

import "fmt"

func main() {
    var a [5]int
    fmt.Println("empty array:", a)

    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))

    b := [5]int{1,2,3,4,5}
    fmt.Println("b initialized:", b)

    const rows = 5
    const cols = 5
    var mult_tab [rows][cols]int
    for row := 0; row < rows; row++ {
        for col := 0; col < cols; col++ {
            mult_tab[row][col] = (row + 1) * (col + 1)
        }
    }
    fmt.Println("mult_tab:", mult_tab)
}
