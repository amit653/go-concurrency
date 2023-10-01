package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func main() {
	msg = "hello"
	wg.Add(1)
	go updateMessage("hello A")
	wg.Wait()

	wg.Add(1)
	go updateMessage("hello B")
	wg.Wait()

	wg.Add(1)
	go updateMessage("hello C")
	wg.Wait()
}

func updateMessage(s string) {
	defer wg.Done()
	msg = s
	printMessage()

}
func printMessage() {
	fmt.Println(msg)
}

