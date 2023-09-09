package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to our Pizza App!!")
	fmt.Println("Please rate our pizza between 1 to 5")

	// var num string
	// fmt.Scan(&num)
	// fmt.Println("Thanks for rating:", num)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating our pizza: ", input)

	// Increment rating by one
	numRating, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

}
