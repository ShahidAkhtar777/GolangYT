package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO with web request:")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of Type: %T\n", response)
	defer response.Body.Close()

	databyte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databyte)
	fmt.Println("Content:\n", content)
}
