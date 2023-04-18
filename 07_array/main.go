package main

import "fmt"

func main() {
	fmt.Println("Welcome to array")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"

	fmt.Printf("My Fruit List: %v has length: %v\n", fruitList, len(fruitList))

	// Method 2
	var vegList = [3]string{"tomato", "brinjal", "beans"}
	fmt.Println("My VegList:", vegList)
}
