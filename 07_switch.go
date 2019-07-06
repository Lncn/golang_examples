package main

import (
    "fmt"
    "time"
)

func main() {
    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    fmt.Print("\nSwitching on value 'time.Now().Weekday()': ", time.Now().Weekday(), "\n")
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend!")
    default:
        fmt.Println("It's a weekday...")
    }

    // Switch without an expression is an alternate way to express if/else logic
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    // You can also use a switch statement to compare a type instead of value
    // This is called a 'type switch'.
    // This is also cool, you can define a function inline another function
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }

    whatAmI(true)
    whatAmI(1)
    whatAmI("hello")
}
