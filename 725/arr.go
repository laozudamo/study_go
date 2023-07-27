package main

import "fmt"

func main() {
	handleArr()
	//fuArr()

	//printArr(a)
}

/*
	数组
	var arr [1]string{"jack"}
*/

func handleArr() {
	a := [3]int{1, 2, 4}             // 未初始化元素值为 0。
	b := [3]int{1, 2, 4}             // 通过初始化值确定数组长度。
	c := [5]int{2: 100, 4: 200}      // 使用引号初始化元素。
	d := [...][2]int{{1, 2}, {1, 2}} // 多维数组 第 2 纬度不能用 "..."。
	count := len(d)                  // 内置函数 len 和 cap 都返回数组长度 (元素数量)。
	//arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d[1][0])
	fmt.Printf("count: %v\n", count)
	//rangeArr()

	printArr(&a) // 传递数组指针, 会改变原数组的值
	fmt.Println(a)
	printArr1(b) // 传递数组, 不会改变原数组的值
	fmt.Println(b)
}

func rangeArr() {
	d := [...][2]int{{1, 2}, {3, 4}}
	for _, v := range d {
		//fmt.Printf("i: %v, v: %v\n", i, v)
		for k, item := range v {
			fmt.Printf("k: %v, item: %v\n", k, item)
		}
	}
}

/*
	数组拷贝和传参
*/

//func fuArr() {
// 数组转为切片
//arr1 := [3]int{1, 2, 3}
//println(arr1)
//slice := arr1[:]

// 向数组追加一个元素
//arr2 = append(slice, 4)
//}

func printArr(arr *[3]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
func printArr1(arr [3]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func testFun() {
	arr := [5]int{1, 3, 5, 8, 7}
	length := len(arr)
	println(length)

}
