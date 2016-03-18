## Control Statements and Loops

A control statement is a statement that determines whether other statements will be executed.
A loop decides how many times to execute another statement.

=====================================================================================
### For

for loop has three components separated by semicolons:

* the init statement: executed before the first iteration
* the condition expression: evaluated before every iteration
* the post statement: executed at the end of every iteration

The loop will stop iterating once the boolean condition evaluates to false.


    for init; condition; post {
        // run commands till condition is true
    }

#### Example

    package main

    import "fmt"

    func main() {
      sum := 0
      for i := 0; i < 10; i++ {
        sum += i
      }
      fmt.Println(sum)
    }

([Playground](http://play.golang.org/p/BRdFL9QtU8))


========================================================================================

### While


C's while is spelled for in Go.

    initialization

    for condition {
          statements;
          adjustment;
    }



#### Example

      package main

      import "fmt"

      func main() {
        sum := 1
        for sum < 10 {
          sum += sum
        }
        fmt.Println(sum)
      }


([Playground](http://play.golang.org/p/gckruLHj8O))

============================================================================================================
### Forever

If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.


#### Example

      package main

      import "fmt"

      func main() {
        for {
          fmt.Println("To infinity and beyond")
        }
      }
===================================================================================================

### If

#### Example

      package main

      import (
        "fmt"
        "math"
      )

      func sqrt(x float64) string {
        if x < 0 {
          return sqrt(-x) + "i"
        }
        return fmt.Sprint(math.Sqrt(x))
      }

      func main() {
        fmt.Println(sqrt(2), sqrt(-4))
      }


([Playground](http://play.golang.org/p/oTup-9cN8l))
===================================================================================================
### If and else

#### Example 1

      package main

      import  "fmt"

      func main() {
          i := 12
          if i<=10 {
            fmt.Println("<10")
          }else{
            fmt.Println(">10")
          }
      }
([Playground](http://play.golang.org/p/NV8zQ3KNL9))

#### Example 2

      func main() {

      if 7%2 == 0 {
          fmt.Println("7 is even")
      } else {
          fmt.Println("7 is odd")
      }

      // A statement can precede conditionals; any variables  declared in this statement are available in all branches.
      if num := 9; num < 0 {
          fmt.Println(num, "is negative")
      } else if num < 10 {
          fmt.Println(num, "has 1 digit")
      } else {
          fmt.Println(num, "has multiple digits")
      }
  }

([Playground](http://play.golang.org/p/2_PF8lI8ai))

============================================================================================================
### Switch

Avoid doing complex series of if else statements.
Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

#### Example

    package main

    import "fmt"

    func main() {
      num := 3
      v := num % 2
      switch v {
      case 0:
        fmt.Println("even")
      case 3 - 2:
        fmt.Println("odd")
      }
    }
([Playground](http://play.golang.org/p/zhsjeP8HNG))

===============================================================================
### Switch with no condition

#### Example

    package main

    import (
      "fmt"
      "time"
    )

    func main() {
      t := time.Now()
      switch {
      case t.Hour() < 12:
        fmt.Println(time.Now())
        fmt.Println("Good morning!")
      case t.Hour() < 17:
        fmt.Println(time.Now())
        fmt.Println("Good afternoon.")
      default:
        fmt.Println(time.Now())
        fmt.Println("Good evening.")
      }
    }



================================================================================
### Defer

#### Example

A defer statement defers the execution of a function until the surrounding function returns.

    package main

    import "fmt"

    func main() {
      defer fmt.Println("world")

      fmt.Println("hello")
    }

([Playground](http://play.golang.org/p/KsaVkG5PUe))

=====================================================================================

### break keyword

#### Example

    package main

    import "fmt"

    func main() {
        i := 0
        for { //since there are no checks, this is an infinite loop
            if i >= 3 { break } //break out of this for loop when this condition is met
            fmt.Println("Value of i is:", i)
            i++;
        }
        fmt.Println("A statement just after for loop.")
    }

([Playground](http://play.golang.org/p/dgo5EepQiG))

=====================================================================================
### continue keyword
      You can go to the beginning of the for loop using continue

#### Example

    package main

    import "fmt"

    func main() {


        for i := 0; i<7 ; i++ {
            if i%2 == 0 {
                continue
            }
            fmt.Println("Odd:", i)
        }
    }

([Playground](http://play.golang.org/p/apY6Zh2TY9))

========================================================================================

### range keyword

#### Example

    package main

    import "fmt"

    func main() {
        //on an array, range returns the index
        a := [...]string{"a", "b", "c", "d"}
        for i := range a {
            fmt.Println("Array item", i, "is", a[i])
        }

        //on a map, range returns the key
        capitals := map[string] string {"France":"Paris", "Italy":"Rome", "Japan":"Tokyo" }
        for key := range capitals {
            fmt.Println("Map item: Capital of", key, "is", capitals[key])
        }

        //range can also return two items, the index/key and the corresponding value
        for key2, val := range capitals {
            fmt.Println("Map item: Capital of", key2, "is", val)
        }
    }


([Playground](http://play.golang.org/p/o6sRlWejQD))