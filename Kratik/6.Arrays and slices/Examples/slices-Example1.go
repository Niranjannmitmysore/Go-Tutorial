  package main

  import "fmt"

func main() {
  x := make([]float64, 5,10)
  fmt.Println(x);
  x[0]= 1;
  x[1]= 2;
  x[2],x[3],x[4]= 3,4,5;
  fmt.Println(x);

  for v := range x {
    fmt.Println(x[v])
  }
  fmt.Println("\n")
  for v := 0; v < len(x); v++ {
    fmt.Println(x[v])
  }

}