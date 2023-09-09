package main

import "fmt"

type User struct {
	FName    string
	LName    string
	Email    string
	IsActive bool
	Age      int
}

func main() {
	fmt.Println("Methods in Golang:-")

	shahid := User{"Shahid", "Akhtar", "shahid@gmail.com", true, 21}

	fmt.Println(shahid)
	fmt.Printf("Shahid details are: %+v\n", shahid)
	fmt.Printf("Shahid mail is: %+v\n and age is: %+v\n", shahid.Email, shahid.Age)

	// Methods
	fmt.Printf("Status:")
	shahid.GetStatus()

	// Functions
	sum := getSum(3, 4)
	fmt.Println("Sum of two numbers: ", sum)

	proRes := proAdder(2, 3, 4, 6, 7)
	fmt.Println("Addition of all numbers is: ", proRes)
}

func (u User) GetStatus() {
	fmt.Printf("%+v Is Active? -- %+v\n", u.FName, u.IsActive)
}

// variadic parameters
func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}
