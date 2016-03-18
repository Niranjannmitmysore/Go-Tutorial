package main

import "fmt"

func main() {
  cars := []string{"bmw", "jaquar", "range_rover", "mercedes"}

  for c := range cars {
    fmt.Println(cars[c])
  }
}