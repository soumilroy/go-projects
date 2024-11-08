package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader // type

// Build our own data type using structs

type User struct {
	UserName string
	Age      int
	FavNum   float64
	OwnsADog bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readString("Enter your name: ")
	user.Age = readInt("Enter your age: ")
	user.FavNum = readFloat("Enter your favorite number: ")
	user.OwnsADog = readBool("Do you own a dog? (y/n)")

	fmt.Printf("Hello %s! Your age is %d and your Fav number is %.2f. You own a dog? %t\n",
		user.UserName,
		user.Age,
		user.FavNum,
		user.OwnsADog,
	)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	fmt.Println(s)
	prompt()

	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\r\n", "", -1)
	input = strings.Replace(input, "\n", "", -1)

	return input
}

func readInt(s string) int {
	input := readString(s)
	num, err := strconv.Atoi(input)

	if err != nil {
		panic(err)
	}

	return num
}

func readFloat(s string) float64 {
	input := readString(s)
	fNum, err := strconv.ParseFloat(input, 64)

	if err != nil {
		panic(err)
	}

	return fNum
}

func readBool(s string) bool {
	fmt.Println(s)
	prompt()

	err := keyboard.Open()

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			panic(err)
		}

		if strings.ToLower(string(char)) == "y" {
			return true
		} else if strings.ToLower(string(char)) == "n" {
			return false
		} else {
			fmt.Println("Invalid input. Please enter 'y' or 'n'")
		}
	}
}
