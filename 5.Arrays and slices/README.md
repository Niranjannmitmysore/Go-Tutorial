## Arrays
An array is a numbered sequence of elements of a single type with a fixed length. In Go they look like this:

    var x [5]int

x is an example of an array which is composed of 5 ints. Try running the following program:

    package main

    import "fmt"

    func main() {
      var x [5]int
      x[4] = 100
      fmt.Println(x)
      fmt.Println("len:", len(x))
    }

You should see this:

    [0 0 0 0 100]
    len: 5


#### Example
[Array Example](Examples/arrays-Example1.go)([Go Playground](http://play.golang.org/p/9eRi-7q3YZ))

### Array with Parameters

[Array with params Example](Examples/Array-with-params.go)([Go Playground](http://play.golang.org/p/ug9dOvqE6N))



==================================================================================================================


## Slices

A slice is a segment of an array. Like arrays slices are indexable and have a length. Unlike arrays this length is allowed to change. Here's an example of a slice:

    var x []float64

The only difference between this and an array is the missing length between the brackets. In this case x has been created with a length of 0.

If you want to create a slice you should use the built-in make function:

   x := make([]float64, 5)

This creates a slice that is associated with an underlying float64 array of length 5.

   x := make([]float64, 5, 8)

This creates a slice that is associated with an underlying float64 array of length 5 and 8 is the capacity.

[Example](Example/Slices-Capacity.go)

#### Example
[Slice Example](Example/slices-Example1.go) ([Go Playground](http://play.golang.org/p/LfhqQPvEak))


Another way to create slices is to use the [low : high] expression:

     slice := [5]float64{1,2,3,4,5}

low is the index of where to start the slice and high is the index where to end it (but not including the index itself). For example while slice[0:5] returns [1,2,3,4,5], slice[1:4] returns [2,3,4].

#### Example

[Slice Example] (Examples/slices-Example2.go)      ([Go Playground](http://play.golang.org/p/T68jK-rlGf))


### Sort

[Sort Example] (Examples/sort.go)      ([Go Playground](http://play.golang.org/p/fp7sXZJZaz))


