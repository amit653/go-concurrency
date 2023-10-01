package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

// test race using go run -race .
func main() {
	var mx sync.Mutex
	msg = "hello"
	wg.Add(3)

	go updateMessage("hello A", &mx)
	go updateMessage("hello B", &mx)
	go updateMessage("hello C", &mx)

	wg.Wait()
	//fmt.Println(msg)

}

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	printMessage()
	m.Unlock()

}
func printMessage() {
	fmt.Println(msg)
}

/*
// Fix data race using sync.mutex
go run  -race . 
hello B
hello C
hello A
*/
