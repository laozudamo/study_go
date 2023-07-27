package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//maphandle()
	//mapdelete()
	foreachByOrder()
}

/*
	map数据结构
	1. map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
	2. map的定义：var mapname map[keytype]valuetype
*/

func maphandle() {

	myMap := map[string]int{
		"one": 1,
		"two": 2,
		"ten": 10,
		"ele": 11,
	}

	for key, val := range myMap {
		println(key, val)
	}

	map1 := map[string]string{
		"one": "1",
		"two": "2",
		"ten": "10",
	}

	fmt.Println(map1["one"])

	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100

	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	fmt.Println(scoreMap)

}

/*
	使用delete()内建函数从map中删除一组键值对，delete()函数的格式如下：
	delete(map, key)
*/

func mapdelete() {
	m1 := map[string]int{
		"one": 1,
		"two": 2,
		"ten": 10,
		"ele": 11,
	}
	delete(m1, "one")
	fmt.Println(m1)
}

// 按照指定顺序遍历map
func foreachByOrder() {
	rand.NewSource(time.Now().Unix())
	var m1 = make(map[string]int, 20)
	for i := 0; i < 20; i++ {
		value := rand.Intn(100)
		k := strconv.Itoa(i)
		m1[k] = value
	}

	fmt.Println(m1)
}
