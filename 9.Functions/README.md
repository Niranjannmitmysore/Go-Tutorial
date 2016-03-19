Function
====

The idea of function in Go is similar to that of C.
A function is a set of code enclosed within flower braces which (usually) performs only one set of operarion.

Example

```
func method_name(arguements_1, arguements_2,.....) [return_type]{
	// Your code
}

```

##### Example
```
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

```
##### Try online here
 ([Go Playground](https://play.golang.org/p/7UPdlfdzJ9))

##### Try another example

```
package main

import (
   "fmt"
)

func main() {
   /* local variable definition */
   var a int = 100
   var b int = 200
   var ret int

   /* calling a function to get max value */
   ret = max(a, b)

   fmt.Printf( "Max value is : %d\n", ret )
}

/* function returning the max between two numbers */
func max(num1, num2 int) int {
   /* local variable declaration */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result 
}

```
##### Try online here
 ([Go Playground](https://play.golang.org/p/iUviT9NJR1))


##### Returning multiple values from a function

```
package main

import "fmt"

func swap(x, y string) (string, string) {
   return y, x
}

func main() {
   a, b := swap("Mahesh", "Kumar")
   fmt.Println(a, b)
}

```
##### Try online here
 ([Go Playground](https://play.golang.org/p/RsT_KBU0sr))

##### Say ola to pointers

```
package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)
    zeroval(i)
    fmt.Println("zeroval:", i)
    zeroptr(&i)
    fmt.Println("zeroptr:", i)
    fmt.Println("pointer:", &i)
}

```
##### Try online here
([Go Playground](https://play.golang.org/p/F6sVEtXxMI))
