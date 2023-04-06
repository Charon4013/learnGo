package main

import (
	"fmt"
	"sync"
	"time"
)

func say(id string) {
	time.Sleep(time.Second)
	fmt.Println("I am done! Id: ", id)
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	wg.Add(3)

	go say("Hello")
	go say("World")
	go func(id string) {
		fmt.Println("HERE ", id)
		wg.Done()
	}("idInput")
	wg.Wait()
	fmt.Println("exit...")
}
