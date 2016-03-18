    package main
    import "fmt"
    func main() {
      fmt.Println("Hello World")
    }


// Every Go program starts with a package declaration which provides a way for
// use to reuse code

    package main

// import allows use to use code from other packages
// The format package provides formatting for input and output

    import "fmt"


// Functions start with func and surround the code with { }
// main is the function that is executed when you execute your program

    func main() {

        // Println is a function in the fmt package that outputs a string, which
        // is surrounded by double quotes and a newline to the screen
        fmt.Println("Hello World")

        // Arithmetic Operators : +, -, *, /, %
        fmt.Println("6 + 4 =", 6 + 4)
        fmt.Println("6 - 4 =", 6 - 4)
        fmt.Println("6 * 4 =", 6 * 4)
        fmt.Println("6 / 4 =", 6 / 4)
        fmt.Println("6 % 4 =", 6 % 4)
      }


([Playground](http://play.golang.org/p/YxGaFNIA4l))


