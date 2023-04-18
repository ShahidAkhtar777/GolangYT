package main

import (
	"fmt"
)

func main() {
	fmt.Println("~~ Welcome to Maps in Golang ~~")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["Py"] = "Python"
	languages["Rb"] = "Ruby"

	fmt.Println("Stored Map: ", languages)
	fmt.Println("JS stands for: ", languages["JS"])

	delete(languages, "Rb")
	fmt.Println("New Map: ", languages)

	// Display all maps
	for key, val := range languages {
		fmt.Printf("Key <%v> has Value <%v>\n", key, val)
	}
}
