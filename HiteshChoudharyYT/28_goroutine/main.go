package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition: LCO")

	var score = []int{0}
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		mut.Lock() //Lock the resources/memory
		score = append(score, 1)
		mut.Unlock() //Release the resource/memory
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Four R")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
