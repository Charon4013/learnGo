package main

import "fmt"

func main() {

	//var intVal int64 = 10
	//stringVal := "Hello"

	type book struct {
		name  string
		pages int
	}
	a := "fuxx the world"
	b := transStringToList(a)
	fmt.Println(b)

}

func transStringToList(a string) string {
	nameCharAr := []byte(a)
	b := ""
	for i := len(nameCharAr) - 1; i >= 0; i-- {
		b += string(nameCharAr[i])
	}
	return b
}
