## Intro to Golang

 Go is a concurrent programming language created by Robert Griesemer, Rob Pike, and Ken Thompson . It was Sponsored and introduced by Google in 2009. It’s been mentioned by the creators that Go was originally developed as a language for servers; but has grown, alongside its epic community, into something much more general purpose.

### Hello World!

```
      package main

      import "fmt"

      func main() {

          fmt.Println("Hello world!")

      }

```
The program defines a new package (main is always the package that contains the main function) then imports the input/output formatting package (fmt), defines its main function (which is the program’s entry point), and uses the Println function from the fmt package to print the string Hello world!.


### Why Go?

* It is fast. And not only fast in the sense that programs written in it run fast when compared to other common languages; but also fast in the sense that its compiler can compile projects in the blink of an eye. So much so that it makes Go feels like an interpreted language instead of a compiled one. You can even edit and run Go programs directly on the Web.

*  Although it is compiled, Go is also a garbage-collected language. This puts less pressure on the developer to do memory
  management, as the language itself takes care of most of the grunt work needed.

*  Go has built-in concurrency, which allows parallelism in an easier way than is possible in other languages. Go has the concept
  of goroutines to start concurrent work and the concept of channels to permit both communication

