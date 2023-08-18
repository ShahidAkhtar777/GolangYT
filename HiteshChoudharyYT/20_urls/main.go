package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=jfs2dfn435"

func main() {
	fmt.Println("Welcome to handling URLs in Golang:")
	fmt.Println(myurl)

	//Parsing
	result, _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("Query params are: %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Params are: ", val)
	}

	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsofUrl.String()
	fmt.Println(anotherUrl)
}
