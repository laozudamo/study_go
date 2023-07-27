package main

import "fmt"

func main() {
	difference()
}

/*
	pointer 指针地址、指针类型和指针取值

	Go语言中的函数传参都是值拷贝，当我们想要修改某个变量的时候，我们可以创建一个指向该变量地址的指针变量;
	Go语言中的指针操作非常简单，只需要记住两个符号：&（取地址）和*（根据地址取值);
	取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。

*/

/*
	空指针
*/

/*
	make和new的区别
	new只用于分配内存，返回一个指向地址的指针
	make只可用于slice,map,channel的初始化,返回的是引用。
*/

func difference() {
	var p *int   // 声明一个指针变量
	p = new(int) // 分配内存,初始化为0
	fmt.Printf("%T\n", *p)
	fmt.Println(*p)

	var s []int = make([]int, 3) // 创建一个切片
	fmt.Println(s)               // [0 0 0] 返回的是引用
}
