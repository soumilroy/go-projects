package main

import (
	"fmt"
	"sort"
)

// basic types, numbers, strings, booleans
var myInt int
var myInt16 int16
var myInt32 int32
var myInt64 int64

var myUint uint
var myUint16 uint16
var myUint32 uint32
var myUint64 uint64

var myFloat float32
var myFloat64 float64

var myStr string // strings are immutable
var myBool bool

// composite types, arrays, slices, maps, structs

var myArr [5]int // arr of length 5 of int type
var myStrArr [5]string

// assign
// myArr[0] = 1
// myArr[1] = 2
// myArr[2] = 3
// myArr[3] = 4
// myArr[4] = 5

// myStrArr[0] = "one"
// myStrArr[1] = "two"
//
// Structs

type Person struct {
	Name    string
	Age     int
	Address string
}

// var p1 Person
// p1.Name = "John Doe"

// p1 := Person{
// 	Name:    "John Doe",
// 	Age:     30,
// 	Address: "123 Main St",
// }
// reference types, pointers, slices, maps, functions, channels

// interface types, error type, interface{}, custom interfaces

func main() {
	var animals []string

	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "fish")

	fmt.Println(animals)

	newSlice := append(animals, "bird")

	fmt.Println(newSlice)

	for i, animal := range animals {
		fmt.Println(i, animal)
	}

	fmt.Println("Length of animals slice: ", len(animals))
	fmt.Println("First 2 animals", newSlice[0:2])

	// sorting is easier with slices than arrays

	fmt.Println("Sorting animals slice", sort.StringsAreSorted(newSlice))
	sort.Strings(newSlice)

	fmt.Println("Sorted animals slice", newSlice)
	// capacity of slice
	fmt.Println("Capacity of animals slice: ", cap(animals))
}

func deleteFromSlice() {}

// pointers
func pointersFn() {
	var myInteger int
	myInteger = 42

	fmt.Println(myInteger)

	// address of myInteger
	myFirstPtr := &myInteger
	fmt.Println(myFirstPtr)
	// 0xc0000120f0

	*myFirstPtr = 21
	fmt.Println(myInteger)   // 21
	fmt.Println(*myFirstPtr) // 21

	changeValueOfPtrNum(&myInteger)

	fmt.Println(myInteger)   // 100
	fmt.Println(*myFirstPtr) // 100
}

func changeValueOfPtrNum(num *int) {
	*num = 100
}
