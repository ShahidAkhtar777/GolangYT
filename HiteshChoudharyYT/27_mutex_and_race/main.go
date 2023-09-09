package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var signals = []string{"text"}

var wg sync.WaitGroup // pointers

var mut sync.Mutex //pointer

func main() {
	// Mutex provides us lock over the memory
	// Use case: Say we have fired multiple goroutine writing on a single memory space
	// But yes we can use go multiple goroutine to read a single memory space

	// What it does is: It says hey I am going to lock this memory this this one goroutine is working
	// and till the time its using that memory space i will not allow other goroutines to access it
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 4; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		log.Fatal(err)
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()

	fmt.Printf("%d status code for website %s\n", res.StatusCode, endpoint)
}
