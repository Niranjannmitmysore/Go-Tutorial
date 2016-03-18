/* A golang-struct is a class with fields where all the methods are non-virtual. */

package main

import (
	"fmt"
)

type Rectangle struct {
  Name          string
  Height, Width float64
}

func main() {

  var a = Rectangle{"Kedarnag", 10, 20}
  var b = Rectangle{Height: 12, Width: 14}

  fmt.Println(a)
  fmt.Println(b)
}
