package main

import "fmt"

func main() {
	// Errors
	// Login("", "dbjhbd27272")
	// Login("mohit.malhotra@barco.co", "")
	// Login("mohit.malhotra@barco.com", "dwjhbwjhdb236823jbjh")

	// Panic & Recover
	_, status := InitDB("user:password@tcp(127.0.0.1:3306)/products")
	fmt.Println("DB(3306) state = ", status)

	_, otherStatus := InitDB("user:password@tcp(127.0.0.1:5432)/products")
	fmt.Println("DB(5432) state = ", otherStatus)
}
