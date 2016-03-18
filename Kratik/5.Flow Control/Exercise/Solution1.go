package main

import "fmt"

func main() {
    i := 2

    switch  {
    case i<5:
        fmt.Println("<5")
    case i>5:
        fmt.Println(">5")
    case i==5:
        fmt.Println("==5")
    }

}