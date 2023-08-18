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

	fmt.Printf("Status:")
	shahid.GetStatus()
}

func (u User) GetStatus() {
	fmt.Printf("%+v Is Active? -- %+v\n", u.FName, u.IsActive)
}
