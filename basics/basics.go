package main

import "fmt"

// Basic datatypes in Go
func DataTypes() {
	// full initlisation syntax
	var num int = 20
	// type inference
	var numInf = 443
	// shorthand syntax
	port := 5467

	fmt.Printf("%d %d %d", num, numInf, port)

	// dataypes
	var name string = "Mohit"
	var age uint8 = 28
	var score float32 = 24.3443
	var id byte = 254   // byte = 8 bit
	var uid rune = '\a' // rune = 32 bit (int32) , represet unicode characters
	var isActive bool = true

	fmt.Printf(`
        name = %s
        age = %d
        score = %f
        id = %d
        uid = %d
        isActive = %t`,
		name,
		age,
		score,
		id,
		uid,
		isActive)
	fmt.Println()
}

// Complex datatypes in Go
func ComplexDataTypes() {
	// arrays
	var arr [3]int = [3]int{1, 23, -9}
	fmt.Println("Array = ", arr) // cannot do '+'
	// inbuilt println, printf always runs after fmt functions
	// println, printf writes to stderr but fmt writes to stdout
	// println("Array length = ", len(arr))
	fmt.Println("Array length = ", len(arr))
	fmt.Println("arr[1] = ", arr[1])

	// append(arr, 23)   // gives compile error as its a fixed size array

	// dynamic arrays = slice
	slice := []string{"Mohit", "Malhotra"}
	slice2 := make([]string, 4, 4) // 2nd arg is length and 3rd is capacity
	fmt.Println("Slice = ", slice)
	fmt.Println("Slice2 = ", slice2)
	fmt.Println("slice2[1] = ", slice2[1])
	fmt.Println("Slice length = ", len(slice))
	fmt.Println("Slice2 length = ", len(slice2))

	updatedSlice := append(slice, "Coder")
	fmt.Println("Slice length = ", len(slice))
	fmt.Println("Updated slice length = ", len(updatedSlice))
	// fmt.Println("slice[2] = ", slice[2])  // runtime error : index out of range
	fmt.Println("updatedSlice[2] = ", updatedSlice[2])

}

// Loops
func Iterations() {

	arr := [5]int16{132, 43, 0, -23, 5}
	// traverse - basic loop
	// var i int = 0
	// for ; i < len(arr); i++ {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// traverse - range loop
	for i, e := range arr {
		fmt.Printf("arr[%d] = %d\n", i, e)
	}

	names := []string{"Liam", "Elon"}
	for index, name := range names {
		fmt.Printf("name[%d] = %s\n", index, name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}

// Conditions
func isOdd(num int) bool {
	if num%2 == 0 {
		return false
	} else {
		return true
	}
}

func isSupported(osType string) bool {
	switch osType {
	case "macos":
		fallthrough
	case "darwin":
		return true
	case "linux":
		return true
	case "windows":
		return false
	default:
		fmt.Println(osType + " not supported")
		return false
	}
}

func Conditions() {
	fmt.Printf("%d is odd = %t\n", 5, isOdd(5))
	fmt.Println("Linux is supported = ", isSupported("darwin"))
}
