
import "fmt"

func main() {

  cars := []string{}
  fmt.Println(len(cars))


  cars = append(cars, "bmw")
  fmt.Println(len(cars))
}