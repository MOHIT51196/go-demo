package main

import (
	"fmt"
	"strings"
)

/*
There is no 'throw' in Go but you return error from function.
Which can then handled depending on whether it is nill or not
just like JS callbacks. Error object represents exceptions as there is nothing
called exception in Go.
*/
func Validate(param string, key string) (bool, error) {
	if len(strings.TrimSpace(param)) == 0 {
		// return false, errors.New(fmt.Sprintf("%s cannot be empty", key))
		return false, fmt.Errorf("%s cannot be empty", key)
	}
	return true, nil
}

type Tuple struct {
	First  string
	Second string
}

func Login(email string, pass string) {
	// isEmailValid, emailError := validate(email, "email")
	// isPassValid, passError := validate(pass, "password")

	// if isEmailValid && isPassValid {
	// 	fmt.Printf("%s logged in succesfully\n", email)
	// } else {
	// 	// fmt.Println("Failed to login") // missing details : "WHY" failure happened

	// 	if emailError != nil {
	// 		fmt.Println("Failed to login due to :", emailError)
	// 	}

	// 	if passError != nil {
	// 		fmt.Println("Failed to login due to :", passError)
	// 	}
	// }

	// Better way to code
	isValid := true
	for _, tup := range []Tuple{
		{email, "email"},
		{pass, "pass"},
	} {
		isParamValid, err := Validate(tup.First, tup.Second)
		if err != nil {
			fmt.Println("Failed to login due to :", err)
			return
		}
		isValid = isValid && isParamValid
	}
	if isValid {
		fmt.Printf("%s logged in succesfully\n", email)
	} else {
		fmt.Printf("%s logged in failed for some reason\n", email)
	}
}
