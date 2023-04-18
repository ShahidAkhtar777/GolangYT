package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("~~ Welcome to slices ~~")

	var fruitList = []string{}
	fruitList = append(fruitList, "Mango", "Apple", "Orange")

	fmt.Println("FruitList: ", fruitList)

	vegList := make([]string, 3)
	vegList[0] = "Onion"
	vegList[1] = "Potato"
	vegList[2] = "Tomato"

	vegList = append(vegList, "Garlic", "Mint")
	fmt.Printf("Type of DS: %T\n", vegList)
	fmt.Println("List of Vegetables: ", vegList)

	sort.Sort(sort.StringSlice(vegList))
	fmt.Println("Sorted VegList: ", vegList)

	// Sorted Ascending
	isSorted := sort.SliceIsSorted(vegList, func(i, j int) bool {
		return vegList[i] < vegList[j]
	})
	fmt.Println("Is VegList Sorted Ascending: ", isSorted)

	sort.Sort(sort.Reverse(sort.StringSlice(vegList)))

	isSorted = sort.SliceIsSorted(vegList, func(i, j int) bool {
		return vegList[i] > vegList[j]
	})
	fmt.Println("Is VegList Sorted Descending: ", isSorted)

	// Remove values from slice based on index

	var courses = []string{"React.Js", "Javascript", "Node.js", "Python", "Golang", "Ruby"}
	fmt.Println("All Courses: ", courses)

	// Index to be removed
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Printf("Removed index %v form the slice new slice: %v\n", index, courses)
}
