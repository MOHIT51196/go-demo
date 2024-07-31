package main

import (
	"fmt"
	"strings"
)

/*
struct is enclapsulation provided by Go
struct supports inheritance via embeddings and compositions
*/

// type alias
// type HelperFunc = func(string)

// Removing '=' will create a new type
type HelperFunc func(string, string) string

type User struct {
	Name             string
	Email            string
	Region           string
	GenerateUsername HelperFunc
}

func Structs() {
	// unnamed fields initialisation needs all the fields to have value
	generateUsername := func(name, email string) string {
		username := name[:3] + email[3:strings.Index(email, "@")]
		return strings.Replace(username, ".", "", -1)
	}
	var user1 User = User{
		"Mohit",
		"mohit.malhotra@barco.com",
		"IN",
		func(x, y string) string { return "" },
	}
	user2 := User{
		Name:             "Aron",
		Region:           "FR",
		Email:            "aron.flisky@dosomework.com",
		GenerateUsername: generateUsername,
	}

	fmt.Println("Username = ", user2.GenerateUsername(user2.Name, user2.Email))

	// named fields initialisation allows you to skip some fields
	user3 := User{
		Name:   "Miloni",
		Region: "CN",
	}

	users := []User{user1, user2, user3}
	for _, u := range users {
		fmt.Println(u)
	}
}

type EngineType struct {
	Capacity int16
	Type     string
}
type Car struct {
	Name   string
	Brand  string
	Engine EngineType
}

func StructComposition() {
	lambo := Car{
		Name:  "Urus",
		Brand: "Lamborgini",
		Engine: EngineType{
			Type:     "v12",
			Capacity: 4,
		},
	}

	fmt.Printf("Car details = %v\n", lambo)
	fmt.Printf("Car engine = %v\n", lambo.Engine)
	fmt.Printf("Car capacity = %v\n", lambo.Engine.Capacity)
	fmt.Printf("Car brand = %v\n", lambo.Brand)
}

/*
 Go does not support inheritance but Embedding is a kind of inheritance
 where you can get properties of one struct in another
*/
