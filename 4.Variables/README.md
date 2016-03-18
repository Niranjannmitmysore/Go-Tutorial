## Variables

Variables are at the heart of the language and provide the ability to read from and write to memory. In Go, access to memory is type safe. This means the compiler takes type seriously and will not allow us to use variables outside the scope of how they are declared.


## Variable Definition in Go:
A variable definition means to tell the compiler where and how much to create the storage for the variable. A variable definition specifies a data type and contains a list of one or more variables of that type as follows:

    var variable_list optional_data_type;

Here, optional_data_type is a valid Go data type including byte, int, float32, complex64, boolean or any user-defined object, etc., and variable_list may consist of one or more identifier names separated by commas


## Links

[Built-In Types](http://golang.org/ref/spec#Boolean_types)




## Code Review
([Go Playground](http://play.golang.org/p/Zv45CSMaiD))

[Declare and initialize variables](Examples/Example1.go) ([Go Playground](http://play.golang.org/p/6w6hBNE75a))
([Go Playground](http://play.golang.org/p/Zv45CSMaiD))

## Exercises

### Exercise 1

**Part A:** Declare three variables that are initialized to their zero value and three declared with a literal value. Declare variables of type string, int and bool. Display the values of those variables.

**Part B:** Declare a new variable of type float32 and initialize the variable by converting the literal value of Pi (3.14).

([Go Playground](https://play.golang.org/p/1xUWjHMB3I)) |
[Answer](Exercises/Exercise1.go) ([Go Playground](https://play.golang.org/p/d2M0Q3mRnd))

