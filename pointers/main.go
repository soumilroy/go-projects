package main

import (
	"fmt"
)

type User struct {
	Name     string
	UserName string
	Password string
}

func updateNewPasswordByValue(u User) {
	u.Password = "newPassword1234"

	fmt.Println("Copies user modify user struct", u)
}

func updateNewPasswordByPointer(u *User) {
	u.Password = "newPassword1234ByPointer"

	fmt.Println("Pointers modify user struct", u)
}

func main() {

	freshUser := User{
		Name:     "John Doe",
		UserName: "johndoe",
		Password: "password",
	}

	fmt.Println("Learning pointers and shit")

	fmt.Println("Original user struct", freshUser)

	// lets update passwowrd

	freshUser.Password = "Abcd#1234"

	fmt.Println("Direct modify user struct", freshUser)

	updateNewPasswordByValue(freshUser)

	fmt.Println("Unmodified user struct", freshUser)

	updateNewPasswordByPointer(&freshUser) // passing address

	fmt.Println("Original user struct", freshUser)
}
