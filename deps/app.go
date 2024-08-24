package main

import (
	"fmt"

	"github.com/go-zoox/fetch"
)

func main() {
	fmt.Println("Starting........")
	res, err := fetch.Get("https://jsonplaceholder.typicode.com/posts/3")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.JSON())
}
