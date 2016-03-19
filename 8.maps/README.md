Maps
====

We all are familar with the concept of hash tables and hashes.
Hashes are name value pair allocated in memory block. They are a very effective way looking, adding and deleting data.

In Go hash tables are known as maps.

### Declaration and Initialization

```
map[KeyType]ValueType

```

where both KeyType and ValueType may be any type

For example :-

```
var m map[string]int

```

Where,
```
	m      => variable assigned
	map    => key words to create a map
	string => this describes the datatype of the keys in the maps
	int    => this describes the datatype of the keys in the maps

```

Technically speaking, maps are references just like slices.
so when you declare m as a map it's value will be nil. It does not have any values in it. It's an empty map.

### Working with maps

1)Setting a value

```
m["route"] = 66

```

2)Assigning a new variable

```
i := m["route"]

```

3)Finding length of map

```
n := len(m)

```

4)Remove an entry from map

```
delete(m, "route")

```

5)Tricky one. Testing for an element

```
i, ok := m["route"]

```

Here 
'i' is assigned the value stored under the key 'routes'. If there is no value under 'routes', then the value is assigned to 0.

'ok' is actually a boolean variable, which will hold true if the value is stored under 'routes' else it will hold false.

Also,

```
_,ok := m['routes']

```
even this tests for an element but without assigning the value.

Now for some example

##### Examples

1) Lets try initializing a map

```
package main

import "fmt"

func main() {
	commits := map[string]int{
		"abc": 123,
		"def": 456,
		"ghi": 789,
	}
	fmt.Println(commits)
}

```
##### Try online here
([Go Playground](https://play.golang.org/p/rReyVb16F2))

2) Looping a map

```
package main

import "fmt"

func main() {
	commits := map[string]int{
		"abc": 123,
		"def": 456,
		"ghi": 789,
	}
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}
}

```
##### Try online here
([Go Playground](https://play.golang.org/p/P37BTIT3fC))