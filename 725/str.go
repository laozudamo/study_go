package main

import (
	"fmt"
	"strconv"
)

func main() {
	changeString()
	changeType()
}

/*
	字符串方法合集
	len() // 字符串长度
	+	拼接
	// index合集
	index(s, substr) // 子串substr在s中第一次出现的位置，不存在则返回-1
	lastIndex(s, substr) // 子串substr在s中最后一次出现的位置，不存在则返回-1
	indexByte(s, c byte) // 字符c在s中第一次出现的位置，不存在则返回-1
	indexRune(s, r rune) // Unicode字符r在s中第一次出现的位置，不存在则返回-1
	lastIndexByte(s, c byte) // 字符c在s中最后一次出现的位置，不存在则返回-1
	lastIndexRune(s, r rune) // Unicode字符r在s中最后一次出现的位置，不存在则返回-1

	// 大小写转换
	toLower(s) // 转小写
	toUpper(s) // 转大写
	ToTitle(s) // 首字母转换
	ToLowerSpecial(c unicode.SpecialCase, s string) // 转小写
	ToUpperSpecial(c unicode.SpecialCase, s string) // 转大写
	ToTitleSpecial(c unicode.SpecialCase, s string) // 首字母转换

	// 判断前缀和后缀
	hasSuffix() // 判断后缀
	hasPrefix() // 判断前缀

	// 去除空格
	trim(s, set)
	trimLeft() // 去除左边的空格
	trimRight() // 去除右边的空格
	trimRightASCII() // 去除右边ASCII码
	trimLeftASCII() // 去除左边ASCII码
	trimPrefix() // 去除前缀
	trimSuffix() // 去除后缀
	trimSpace() // 去除空格
	trimLeftBytes() // 去除左边字节
	trimRightBytes() // 去除右边字节
	trimLeftFunc() // 去除左边函数
	trimRightFunc() // 去除右边函数
	trimFunc() // 去除函数
	repeat(s, count)
	replace(s, old, new, count)

*/

/*
	修改字符串
	要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string
*/

func changeString() {
	str := "helloWorld网站"
	byteS := []byte(str)
	byteS[0] = 'H'
	str1 := string(byteS)

	runeS := []rune(str)
	runeS[len(runeS)-1] = '玩'
	str2 := string(runeS)
	println(str1)
	println(str2)
}

/*
	类型转换
	Go语言中只有强制类型转换，没有隐式类型转换。
	转换格式：T(表达式) T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等.
*/

func changeType() {
	str := "10"
	num := 20
	a, _ := strconv.Atoi(str)               // string转int
	b := strconv.Itoa(num)                  // int转string
	i64, _ := strconv.ParseInt(str, 10, 64) // string转int64
	uv, _ := strconv.ParseUint(str, 10, 64) // string转uint64
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", i64)
	fmt.Printf("%T\n", uv)
}
