package main

import "fmt"

func main() {
  x := make([]float64, 5,10)
  fmt.Println(x);
  x[0]= 1;
  x[1]= 2;
  fmt.Println(x);
}