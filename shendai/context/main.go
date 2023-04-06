package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建空context的方法
	ctx := context.Background() // 返回一个空context, 不能被cancel, kv为空

	//todoCtx := context.TODO() // 类似

	ctx, cancel := context.WithCancel(ctx)
	//ctx, cancel := context.WithTimeout(ctx, 2 * time.Second)
	//ctx, cancel := context.WithDeadline(ctx, 3 * time.Second)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	doSomething(ctx)

}

func doSomething(ctx context.Context) {
	select {
	case <-ctx.Done(): // ctx is cancelled
		err := ctx.Err()
		fmt.Println("err: ", err.Error())
	case <-time.After(5 * time.Second): // 5 second pass
		fmt.Println("finish doing something...")
	}
	ctx.Done()

}
