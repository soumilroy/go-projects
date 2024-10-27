package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

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
