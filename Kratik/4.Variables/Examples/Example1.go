package main

import "fmt"

func main() {
   var a, b, c = 3, true, "foo"

   fmt.Println(a)
   fmt.Println(b)
   fmt.Println(c)
   fmt.Printf("a is of type %T\n", a)
   fmt.Printf("b is of type %T\n", b)
   fmt.Printf("c is of type %T\n", c)

  var d, e, f = 1, true, "HO HO HO"
  fmt.Println("d=", d, "e=", e, "f=", f)


}
