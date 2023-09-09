package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to Channels in Golang: LCO")
	// Channels are a way in which out multiple goroutines can actually talk to each other
	// They would not be knowing what's going inside them, but maybe we are waiting for just a signal
	// ,and we can just do something on the go without waiting for thread to end execution

	// Buffered channel
	// We don't want to pick all channels data one by one we want to loop through them all and print them
	// The buffered channel comes for rescue here add nums at last
	myChan := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// Below creates a deadlock as channels says add in me when someone is listening
	// It also says listen me only when someone adds in me [Makes deadlock]
	// myChan <- 5
	// fmt.Println(<-myChan)

	wg.Add(2)

	//read only channel <-chan
	// we are not allowed to close channel in this case
	go func(ch <-chan int, wg *sync.WaitGroup) {
		//If we close channel before consuming it makes infinite loop
		val, isChanelOpen := <-myChan

		fmt.Println(isChanelOpen)
		fmt.Println(val)

		//fmt.Println(<-myChan)
		//fmt.Println(<-myChan)
		wg.Done()
	}(myChan, wg)

	// right only channel ch<-
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChan <- 5
		close(myChan)

		//myChan <- 5
		//myChan <- 6
		//close(myChan)
		wg.Done()
	}(myChan, wg)

	wg.Wait()
}
