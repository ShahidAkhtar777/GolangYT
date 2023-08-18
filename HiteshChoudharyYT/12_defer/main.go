package main

import "fmt"

func main() {

	defer fmt.Println("Two")
	fmt.Println("Welcome")
	defer fmt.Println("One")
	defer fmt.Println("World")
	fmt.Println("Hello")
	mydefer()
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

// Some reference docs for Golang
// https://www.velotio.com/engineering-blog/build-a-containerized-microservice-in-golang
