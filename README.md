![CI Workflow](https://github.com/felixenescu/golang-map-set/actions/workflows/ci.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/felixenescu/golang-map-set)](https://goreportcard.com/report/github.com/felixenescu/golang-map-set)
[![GoDoc](https://godoc.org/github.com/felixenescu/golang-map-set?status.svg)](http://godoc.org/github.com/felixenescu/golang-map-set)


# golang-map-set

A simple, [generic](https://go.dev/doc/tutorial/generics), non-threadsafe set type for the Go language. This package provides an easy-to-use and efficient implementation of sets, tailored for developers looking for an essential set functionality in Go.


## Install

Use `go get` to install this package.

```shell
go get github.com/felixenescu/golang-map-set
```

Import it with:

```go
import set "githug.com/felixenescu/golang-map-set"
```

and use `set` as the package name inside the code.

## Usage

Set works with any [comparable](https://go.dev/ref/spec#Comparison_operators) type. In Go, comparable types are those that can be compared using equality operators `==` and `!=`. These types include:

 - Boolean types
 - Numeric types (integer, float, complex)
 - String types
 - Pointer types
 - Channel types
 - Interface types
 - An array type is comparable if the elements' type is comparable.
 - A struct type is comparable if all its fields are comparable.

Note that slices, maps, and functions are not comparable.


### Creating a set

 - **Empty Set:** Create a new empty set.
```go
// create a set of integers
s1 := set.New[int]()

// create a set of strings
s2 := set.New[string]()
```

 - **From Slice:** Create a new set from a slice.

```go
// create a set from a slice of integers
slice := []int{1, 2, 3}
s := set.NewFromSlice(slice)

// create a set from a slice of strings
slice := []string{"a", "b", "c"}
s := set.NewFromSlice(slice)
```

 - **From Map Keys**: Create a new set from the keys of a map.

```go
// create a set from the keys of a map of integers
m := map[int]string{1: "a", 2: "b"}
s := set.NewFromMapKeys(m)

// create a set from the keys of a map of strings
m := map[string]int{"a": 1, "b": 2}
s := set.NewFromMapKeys(m)
```

### Basic Operations

 - Add Elements: Add one or more elements to the set.

```go
s1 := set.New[int]()
s.Add(4) // set is now {4}
s.AddAll([]int{5, 6}) // set is now {4, 5, 6}
```

 - **Remove Elements:** Remove one or more elements from the set.

```go
s := set.NewFromSlice([]int{1, 2, 3, 4, 5})
s.Remove(2) // set is now {1, 3, 4, 5}
s.RemoveAll([]int{3, 4}) // set is now {1, 5}
```


 - **Contains:** Check if an element is in the set.
```go
s := set.New[int]()
s.Add(1)
if s.Contains(1) {
    // element 1 is in the set
    fmt.Println("1 is in the set")
}
```

 - **To Slice:** Convert the set to a slice.

```go
s := set.New[int]()
s.AddAll([]int{1, 2, 3})
slice := s.ToSlice()
fmt.Println(slice) // [1 2 3]
```


### Set Operations

 - **Union:** Combine two sets.

```go
s1 := set.NesFromSlice([]int{1, 2, 3})
s2 := set.NewFromSlice([]int{4, 5, 6})
u := s1.Union(s2) // u is now {1, 2, 3, 4, 5, 6}
```

 - **Intersection:** Get the common elements of two sets.

```go
s1 := set.NewFromSlice([]int{1, 2, 3})
s2 := set.NewFromSlice([]int{2, 3, 4})
i := s1.Intersection(s2) // i is now {2, 3}
```

 - **Difference:** Get the elements that are in one set but not the other.

```go
s1 := set.NewFromSlice([]int{1, 2, 3})
s2 := set.NewFromSlice([]int{2, 3, 4})
d := s1.Difference(s2) // d is now {1}
```

Remember, golang-map-set is **not threadsafe**, so appropriate precautions should be taken when using it in a concurrent environment.