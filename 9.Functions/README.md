Function
====

The idea of fucntion in Go is similar to that of C.
A function is a set of code enclosed within flower braces which (usually) performs one set of operarion.

Example

```
func main() {
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