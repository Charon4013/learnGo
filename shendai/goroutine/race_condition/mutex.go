package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int32
var mtx sync.Mutex
var ch = make(chan int, 1)

func MutexIncCounter() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		mtx.Lock()
		counter++
		mtx.Unlock()
	}
}

func AtomicIncCounter() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		atomic.AddInt32(&counter, 1)
	}
}

func ChannelIncCounter() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		count := <-ch
		count++
		ch <- count
	}
}

func main() {
	wg.Add(2)

	//go MutexIncCounter()
	//go MutexIncCounter()

	//go AtomicIncCounter()
	//go AtomicIncCounter()

	go ChannelIncCounter()
	go ChannelIncCounter()
	ch <- 0

	wg.Wait()
	//fmt.Println(counter)
	fmt.Println(<-ch)

}
