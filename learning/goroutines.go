package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var msg = "hello"
	wg.Add(1) //waiting for one go routine
	go func(msg string) {
		fmt.Println(msg)
		wg.Done() //indicate the waitgroup's wait method that go routine is done
	}(msg)
	msg = "goodbye"
	// time.Sleep(100 * time.Millisecond)
	wg.Wait()

}
