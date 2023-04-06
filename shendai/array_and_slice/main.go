package main

import "fmt"

func modify(array []int) {
	array[0] = 100
}

func main() {

	// array
	//array := [5]string{"zhangsan", "lisi", "wangwu", "chenliu", "sunqi"}
	////array2 := [...]string{"123", "345"}
	//array2 := array // 拷贝所有字符串
	//
	//fmt.Println(array, len(array), cap(array))
	//fmt.Println(array2)

	// slice (起始指针, 长度, 容量)
	//slice := []string{"123", "345", "3456", "114514"}
	//fmt.Println(slice, len(slice), cap(slice))
	//
	//newSilce := make([]int, 2, 2)
	//fmt.Println(newSilce, len(newSilce), cap(newSilce))
	//newSilce = append(newSilce, 114514)
	//fmt.Println(newSilce, len(newSilce), cap(newSilce))
	//modify(newSilce)
	//fmt.Println(newSilce)

	//subSlice := slice[1:3]
	//fmt.Println(subSlice)
	//subSlice[0] = "boss"
	//fmt.Println(slice)

	// 映射map
	scores := map[string]int{
		"ming":  10,
		"zhang": 13,
		"li":    22,
	}

	delete(scores, "ming")

	//fmt.Println(scores["ming"])
	//fmt.Println(len(scores))
	//
	//score, exist := scores["chong"]
	//fmt.Println(score, exist)
	//
	//modifyMap(scores)
	//fmt.Println(scores)

	//for k, v := range scores {
	//	fmt.Printf("key: %s \t value: %d\n", k, v)
	//}

	multiSlice := make([][]int, 0)
	multiSlice = append(multiSlice, []int{1, 2, 3})
	multiSlice = append(multiSlice, []int{4, 5, 6})

	multiSlice[0][1] = 100

	fmt.Println(multiSlice)

}

func modifyMap(m map[string]int) {
	m["zhang"] = 233
	m["liu"] = 122
}
