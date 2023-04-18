package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	var ptr *int
	fmt.Println("Value of ptr is", ptr)

	myNum := 23
	ptr = &myNum

	fmt.Println("Value of ptr: ", ptr)
	fmt.Println("Actual value at address: ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New Value: ", *ptr)
}
