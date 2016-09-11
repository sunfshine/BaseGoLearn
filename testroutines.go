package main

import "fmt"
import "sync"

var count int

func Testroutines() {
	count = 0
	var wg sync.WaitGroup
	for i := 0; i < 1200; i++ {
		wg.Add(1)
		go addandprint(&wg)
		// if pass wg as value, error will happen: all goroutines are asleep
		// go addandprint(&wg)
	}
	wg.Wait()
}

func addandprint(wg *sync.WaitGroup) {
	for i := 0; i < 20; i++ {
		count = count + 1
		fmt.Printf("count %d \n", count)
	}
	wg.Done()
}
