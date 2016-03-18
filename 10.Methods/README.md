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

#### For example

```
/* define a method for circle */
func(circle Circle) area() float64 {
   return math.Pi * circle.radius * circle.radius
}

In this code
'circle' is a variable whereas Circle represents the receiver type of the function area. So basically you can set your own receiver types based on the context.

```