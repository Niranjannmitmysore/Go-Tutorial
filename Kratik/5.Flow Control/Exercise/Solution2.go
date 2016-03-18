package main

import "fmt"

func main() {
  i := 10;

  for j:=1; j<20;j++{
    if j%5 ==0 {
     continue
   }
   fmt.Print(i*j)

   fmt.Print(",")
 }

}