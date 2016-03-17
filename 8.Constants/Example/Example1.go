
package main

import "fmt"

// `const` declares a constant value.
const s string = "constant"
const Pi = 3.14
func main() {

    const World = "世界"
    fmt.Println("Hello", World)

    fmt.Println("Happy", Pi, "Day")

    const Truth = true
    fmt.Println("Go rules?", Truth)

    fmt.Printf("%T: %v\n", "Hello, 世界", "Hello, 世界")

    fmt.Println(s)
    const n = 500000000

    const d = 3e20 / n
    fmt.Println(d)
    fmt.Println(int64(d))

}