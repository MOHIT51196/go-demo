package main

import (
	"fmt"
	"reflect"
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

// has-a raltionship
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
 where you can get properties of one struct in another.
*/

type Item struct {
	Title       string
	Description string
	Tags        map[string]string
}

type Phone struct {
	Item
	Brand string
	Color string
	Model string
}

func CreateItem() {
	item1 := Phone{
		Brand: "Samsung",
		Model: "S24",
		Color: "Onyx Black",
		// Title: "Samsung S24 5G",  // does not work for embeddings

		// Item: Item{ // valid syntax for embedding
		// 	Title:       "Samsung S24 5G",
		// 	Description: "Flagship Samsung S series phone with Sanpdragon 888",
		// 	Tags:        map[string]string{"manufacturer": "Samsung", "launchYear": "2023"},
		// },
	}

	item1.Title = "Samsung S24 5G"
	item1.Description = "Flagship Samsung S series phone with Sanpdragon 888"
	item1.Tags = map[string]string{"manufacturer": "Samsung", "launchYear": "2023"}

	fmt.Println(item1)

	item2 := Phone{
		Item{
			"iPhone 15 Pro",
			"Latest iPhone 15 256 GB with A16 bionic chip to support AI",
			map[string]string{"manufacturer": "Apple", "launchYear": "2024"},
		},
		"Apple",
		"15 Pro",
		"Perl White",
	}

	fmt.Println(item2)
}

type Product interface{}

func TypeChecking() {
	var prod Product = Phone{
		Item: Item{
			"iPhone 15 Pro",
			"Latest iPhone 15 256 GB with A16 bionic chip to support AI",
			map[string]string{"manufacturer": "Apple", "launchYear": "2024"},
		},
	}

	_, isPhone := prod.(Phone) // this kind operation is only supported on interfaces
	fmt.Println("Product is instance of Phone = ", isPhone)

	// check embedding effect
	_, isItem := prod.(Item)
	fmt.Println("Product is instance of Item = ", isItem) // object is not the type of ite embedding

	// Type Switches
	// .(type) is only supported with switch
	switch prod.(type) {
	// case Product:				// its also true
	// 	fmt.Println("Product type")
	case Phone:
		fmt.Println("Phone type")
		// fallthrough		// cannot use fallthrough in type switches
	case Item:
		fmt.Println("Item type")
	}

	// Use refect package
	typ := reflect.TypeOf(prod)
	fmt.Printf("Type : %v , Type name = %v\n", typ, typ.Name())
	fmt.Println("Type check isInterface : ", typ.Kind() == reflect.Interface) // false
	fmt.Println("Type check isPhone : ", typ == reflect.TypeOf(Phone{}))
	fmt.Println("Type check isItem : ", typ == reflect.TypeOf(Item{}))
	_, ok := prod.(Phone)
	fmt.Println("Type check isPhone (type assertion) : ", ok)
}
