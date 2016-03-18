  package main

  import "fmt"

func main() {
  slicA := make([]float64, 5)
  fmt.Println(slicA, len(slicA), cap(slicA))

  slicB := make([]float64, 5,7)
  fmt.Println(slicB, len(slicB), cap(slicB))

  for i:=1; i<7; i++ {
    slicA = append(slicA,1)
    fmt.Println(slicA, len(slicA), cap(slicA))
  }

  for i:=1; i<7; i++ {
    slicB = append(slicB,1)
    fmt.Println(slicB, len(slicB), cap(slicB))
  }

}