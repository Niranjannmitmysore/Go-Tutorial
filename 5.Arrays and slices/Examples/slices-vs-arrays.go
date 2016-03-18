package main
import "fmt"

func main(){
  arr := [5]int{1,2,3,4,5}
  slicA := arr[:]
  fmt.Println(slicA, len(slicA), cap(slicA))
  fmt.Printf("\n%p\n%p\n", &arr,&slicA[0])
  slicA = append(slicA,99)
  fmt.Println(slicA, len(slicA), cap(slicA))
  fmt.Printf("\n%p\n%p\n", &arr,&slicA[0])
  for i:=1; i<10; i++ {
    slicA = append(slicA,99)
    fmt.Println(slicA, len(slicA), cap(slicA))
  }






}