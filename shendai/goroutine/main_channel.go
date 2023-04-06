package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg2 sync.WaitGroup

func player(name string, ch chan int) {
	defer wg2.Done()

	for {
		ball, ok := <-ch // 从通道取值
		if !ok {
			fmt.Printf("channel closed! %s wins\n", name)
			return
		}

		n := rand.Intn(1000)

		if n%10 == 0 {
			close(ch)
			fmt.Printf("%s missed the ball %s loses\n", name, name)
			return
		}

		ball++
		fmt.Printf("%s received ball %d\n", name, ball)
		ch <- ball

	}
}

func main() {
	wg2.Add(2)

	ch := make(chan int, 0)

	go player("sun", ch)
	go player("wang", ch)

	ch <- 0

	wg2.Wait()
}
