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
type User struct {
	FirstName, LastName string
}

func (u User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

In this code
'u' is a variable,
whereas 'User' represents the receiver type
of the function 'Greeting'.
And 'User' has to be defined previously as a struct type

```

#### For example

```
package main

import (
	"fmt"
)

type User struct {
	FirstName, LastName string
}

func (u User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

func main() {
	u := User{"Bhagirath", "Purushotham"}
	fmt.Println(u.Greeting())
}

```
##### Try it online

```
([Go Playground](http://play.golang.org/p/ITVfJkCiwk))

##### Working with pointers

```
package main

import (
	"fmt"
)

type User struct {
	FirstName, LastName string
}

func (u *User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

func main() {
	u := &User{"Bhagirath", "Purushotham"}
	fmt.Println(u.Greeting())
}

```
##### Try it online

```
([Go Playground](http://play.golang.org/p/tEVN-vhyAi))

