package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web server class")

	// getUrl := "http://localhost:8000/get"
	// performGetRequest(getUrl)

	// postUrl := "http://localhost:8000/post"
	// performPostRequest(postUrl)

	performPostFormRequest()
}

func performGetRequest(url string) {

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Response Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println("Content is: ", responseString.String())

	// content, _ := io.ReadAll(response.Body)
	// fmt.Println(string(content))
}

func performPostRequest(url string) {

	// Fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename": "Lets go with Golang",
			"price": "2000",
			"platform": "LearnCodeOnline.in"
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func performPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	// FormData
	data := url.Values{}
	data.Add("firstname", "Shahid")
	data.Add("lastname", "Akhtar")
	data.Add("rollno", "01")

	// url encoded data
	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
