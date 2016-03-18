package main
import "sort"
import "fmt"

func main() {
  var x [5]float64
   x[0] = 98
  fmt.Println(x)

  s := []int{5, 2, 6, 3, 1, 4} // unsorted
  fmt.Println(s)
  sort.Ints(s)
  fmt.Println(s)

  strs := []string{"c", "a", "b"}
  fmt.Println(strs)
  sort.Strings(strs)
  fmt.Println(strs)

  check1 := sort.IntsAreSorted(s)
  fmt.Println("Sorted: ", check1)

  check2 := sort.StringsAreSorted(strs)
  fmt.Println("Sorted: ", check2)
}
