package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
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
// reference types, pointers, slices, maps, functions, channels
// interface types, error type, interface{}, custom interfaces

func main() {
	err := keyboard.Open()
	hasEscPressed := false

	if err != nil {
		log.Fatal(err)
	}

	// as soon as main ends, defer will close the keyboard
	defer func() {
		_ = keyboard.Close()
	}()

	foods := make(map[int]string)

	foods[1] = "Chicken Biryani"
	foods[2] = "Mutton Biryani"
	foods[3] = "Vegetable Biryani"
	foods[4] = "Tea"

	fmt.Println("Menu")
	fmt.Println("----")
	fmt.Println("1 - Chicken Biryani")
	fmt.Println("2 - Mutton Biryani")
	fmt.Println("3 - Vegetable Biryani")
	fmt.Println("4 - Tea")
	fmt.Println("Press ESC or 'q' to quit")

	for {
		char, key, err := keyboard.GetKey()

		if err != nil {
			log.Fatal(err)
		}

		opt, _ := strconv.Atoi(string(char))

		if char == 'q' || char == 'Q' || key == keyboard.KeyEsc {
			if hasEscPressed {
				break
			} else {
				hasEscPressed = true
				fmt.Println("Are you sure you want to exit? Press ESC or 'q' again")
				continue
			}
		}

		if opt < 1 || opt > 4 {
			fmt.Println("Invalid option")
			continue
		}

		fmt.Println(fmt.Sprintf("You chose %s", foods[opt]))
	}

	fmt.Println("Program exiting..")
}
