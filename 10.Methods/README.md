Methods
====

Better lets understand Methods along with a syntax

#### Syntax

```
func (variable_name variable_data_type) function_name() [return_type]{
   /* function body*/
}

```
Here

```
variable_name is the variable which you'll passing as variables
variable_data_type is what we call the Receiver type of the method
return_type is the type which the method will return

```

#### Sample Code

```
/* define a method for circle */
func(circle Circle) area() float64 {
   return math.Pi * circle.radius * circle.radius
}

In this code
'circle' is a variable,
whereas 'Circle' represents the receiver type
of the function 'area'.
And has to be defined previously as a struct type

```

#### Examples to try yourself

```
package main

import (
   "fmt"
   "math"
)

/* define a circle */
type Circle strut {
   x,y,radius float64
}

/* define a method for circle */
func(circle Circle) area() float64 {
   return math.Pi * circle.radius * circle.radius
}

func main(){
   circle := Circle(x:0, y:0, radius:5)
   fmt.Printf("Circle area: %f", circle.area())
}

```