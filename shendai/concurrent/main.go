package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"shendai/concurrent/pool"
	"sync"
	"sync/atomic"
	"time"
)

func createTask() func(int) {
	return func(id int) {
		time.Sleep(time.Second)
		fmt.Printf("Task complete #%d\n", id)
	}
}

// 定义一个资源
type DBConnection struct {
	id int32
}

func (D DBConnection) Close() error {
	fmt.Println("database closed, #", D.id)
	return nil
}

var counter int32

func Factory() (io.Closer, error) {
	atomic.AddInt32(&counter, 1)
	return &DBConnection{id: counter}, nil
}

func performQuery(query int, pool *pool.Pool) {
	defer wg.Done()

	resource, err := pool.AcquireResource()

	if err != nil {
		fmt.Println(err)
	}

	defer pool.ReleaseResource(resource)

	t := rand.Int()%10 + 1
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Println("finish query: ", query)
}

var wg sync.WaitGroup

func main() {

	// Runner
	//r := runner.New(4 * time.Second)
	//
	//r.AddTasks(createTask(), createTask(), createTask())
	//
	//err := r.Start()
	//switch err {
	//case runner.ErrInterrupt:
	//	fmt.Println("tasks interrupted")
	//case runner.ErrTimeout:
	//	fmt.Println("tasks timeout")
	//default:
	//	fmt.Println("all tasks finished")
	//}

	// pool
	p, err := pool.New(Factory, 5)
	if err != nil {
		log.Fatalln(err)
	}

	num := 10
	wg.Add(num)
	for id := 0; id < num; id++ {
		go performQuery(id, p)
	}
	wg.Wait()

	p.Close()

}
