package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var signals = []string{"text"}

var wg sync.WaitGroup // pointers

func main() {
	// Concurrency: We pick task and work on them but not all on the same time. Pick one do it for sometime
	// 				then change task and do that this is concurrency. Context Switching is there.

	// Parallelism: We pick tasks and do all of them at the same time

	// Thread: Managed by OS =>fixed stack - 1 mb
	// Goroutine: Managed by Go runtime =>flexible stack - 2k

	// Do not communicate by sharing memory instead share memory by communicating
	// fire a thread going to call this func with hello but we never waited for their thread to come back
	// go greeter("Hello")
	// greeter("world")

	websitelist := []string{
		"http://lco.dev",
		"http://go.dev",
		"http://google.com",
		"http://github.com",
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

	signals = append(signals, endpoint)

	fmt.Printf("%d status code for website %s\n", res.StatusCode, endpoint)
}
