package main
// Simulate race condition using waitgroup
import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

// test race using go run -race .
func main() {
	//var mx sync.Mutex
	msg = "hello"
	wg.Add(3)

	//go updateMessage("hello A", &mx)
	//go updateMessage("hello B", &mx)
	//go updateMessage("hello C", &mx)
	go updateMessage("hello A")
	go updateMessage("hello B")
	go updateMessage("hello C")
	wg.Wait()
	//fmt.Println(msg)

}

func updateMessage(s string) {
	defer wg.Done()
	//m.Lock()
	msg = s
	printMessage()
	//m.Unlock()

}
func printMessage() {
	fmt.Println(msg)
}

//Output
/*
go run  -race . 
hello C
==================
WARNING: DATA RACE
Write at 0x000102e1aa00 by goroutine 6:
  main.updateMessage()
      /Users/amitg/go/src/first-example/main.go:31 +0x68
  main.main.func1()
      /Users/amitg/go/src/first-example/main.go:20 +0x34
*/
