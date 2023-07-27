package main

const (
	Monday    = iota // 0
	Tuesday          // 1
	Wednesday        // 2
	Thursday         // 3
	Friday
	_        // 4 丢弃
	Saturday // 5
	Sunday   // 6
)

func main() {

	var (
		a int
		b int
		c int = 3
	)

	println(a, b, c)

	println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
}

/*
	iota是go语言的常量计数器，只能在常量的表达式中使用。 iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。
*/
/*
	iota常用于变量的枚举
*/
