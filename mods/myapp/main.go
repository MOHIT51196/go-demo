package main

import (
	"fmt"

	users "example.com/db"
)

func main() {
	userService := users.UserService{}

	userService.Add(users.User{
		Name:   "Mohit",
		Email:  "mohit.malhotra@barco.com",
		Region: "IN",
	})

	userService.Add(users.User{
		Name:   "Ron Shazam",
		Email:  "shazam.ron@abc.com",
		Region: "KOR",
	})

	fmt.Printf("Users = %v\n", userService.FindAll())
	isFound, foundUser := userService.FindByEmail("mohit.malhotra@barco.com")
	fmt.Printf("Find user for id %s [%t] = %v\n", "mohit.malhotra@barco.com", isFound, foundUser)

	userService.Delete(users.User{
		Name:   "Ron Shazam",
		Email:  "shazam.ron@abc.com",
		Region: "KOR",
	})

	fmt.Printf("Users = %v\n", userService.FindAll())
}
