package main

import "fmt"

type User struct {
	Name  string
	Email string
	Phone string
	Age   int
}

func main() {
	fmt.Println("~~ Welcome to Structs ~~")

	usr1 := User{"Shahid Akhtar", "shahidakhtar016@gmail.com", "9102380505", 24}

	fmt.Println("User is", usr1)
	fmt.Printf("Detailed Val: %+v\n", usr1)

	// While Loop implementation in Golang
	n := 5

	for n < 7 {
		fmt.Println("C++")
		n++
	}
}
