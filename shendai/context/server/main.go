package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler start")
	ctx := r.Context()

	complete := make(chan struct{})

	go func() {
		// do something
		time.Sleep(5 * time.Second)
		complete <- struct{}{}
	}()

	select {
	case <-complete: // finish doing something
		fmt.Println("finish doing something")
	case <-ctx.Done(): // ctx is cancelled
		err := ctx.Err()
		fmt.Println("err: ", err.Error())
	case <-time.After(5 * time.Second): // 5 second pass
		fmt.Println("finish doing something...")
	}
	fmt.Println("handler end")

	//time.Sleep(5 * time.Second)
	//_, err := fmt.Fprintln(w, "Hello world")
	//if err != nil {
	//	log.Fatalln(err)
	//}
}
