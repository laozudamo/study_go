package main

import "fmt"

func main() {
	//createSlice()
	splSlice()
}

/*
	slice
    1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
    2. 切片的长度可以改变，因此，切片是一个可变的数组。
    3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
    4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
    5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
    6. 如果 slice == nil，那么 len、cap 结果都等于 0。
*/

/*
	切片的创建
*/

func createSlice() {
	var s1 []int                     // 声明一个未指定大小的数组来定义切片
	s2 := []int{1, 2, 3, 4, 5, 6, 7} // 声明并初始化一个切片
	s3 := make([]int, 3)             // 使用make()函数来创建切片: make([]T, length, capacity)，其中capacity为可选参数
	s4 := s2[0:4]                    // 基于数组或切片生成新的切片: slice[开始位置:结束位置]，左包含右不包含，即 [1,4)
	fmt.Println(s1, s2, s3, s4)
}

/*
切片截取
s := arr[startIndex:endIndex]

切片拷贝
copy(s2, s1) // 将 s1 拷贝到 s2 中
*/
func splSlice() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	hs := make([]int, 3)
	fmt.Println(len(hs))
	//ni := []int{1, 2, 3, 4, 5}
	ns := make([]int, 5)   // 创建一个切片
	s1 := s[1:4]           // 2,3,4
	s2 := s[:4]            // 1,2,3,4
	s3 := s[3:]            // 4,5,6,7
	s4 := s[:]             // 1,2,3,4,5,6,7
	n1 := append(ns, s...) // 将 s1 的元素追加到 ns 中
	fmt.Println(s1, s2, s3, s4, ns, n1)
}
