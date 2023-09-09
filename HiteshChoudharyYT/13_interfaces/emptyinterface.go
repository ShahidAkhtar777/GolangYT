package main

import "fmt"

type i interface{}

func found(i interface{}) {
	fmt.Printf("Type=%T, Value=%v\n", i, i)
}

func typeSwitch(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I learn at %s\n", i.(string))
	case int:
		fmt.Printf("This is %d (number)\n", i.(int))
	default:
		fmt.Println("Unknown Type!!")
	}

}

func main() {
	// When a interface has zero methods its called empty interface
	str := "Golang"
	found(str)

	//	Type Switch Example
	typeSwitch(str)
	typeSwitch(56)
	typeSwitch(45.5)
}
