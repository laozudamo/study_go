package main

import (
	"fmt"
	"strings"
)

func main() {
	//str := getMoreLineStr()
	//fmt.Print(str)

	handleStr()
}

// rune int32 rune类型，代表一个 UTF-8字符
/*
	在 Go 语言中，rune 类型用于表示 Unicode 码点，
	即 Unicode 字符集中的一个字符，它实际上是一个 int32 类型的别名，用于表示 Unicode 编码的字符。
	因此，rune 类型的数据是一种整数类型，但它可以表示 Unicode 字符，而不仅仅是 ASCII 字符。

  s := "pprof.cn博客"
        for i := 0; i < len(s); i++ { //byte
            fmt.Printf("%v(%c) ", s[i], s[i])
        }
*/

// bool

// byte unit8 0 ~ 255

// int8 int16 int32 int64 uint8 uint16 uint32 uint64

// float32 flot64

// array

// struct

// string

// slice

// map

// channel

// interface

// function

/*
	字符串转义符
	1. /r
	2. /n
	3. \t	制表符
	4. \' 单引号
	5. \"	双引号
	6. \ 反斜杠
*/

/*
	go语言中定义多行字符串用反引号
*/

func getMoreLineStr() string {
	str := `
	hello
	jack
	`
	return str
}

/*
	字符串的常用方法
	len(str)
	+
	strings.Split()
	strings.clone()
	strings.Contains()
	strings.HasPrefix()
	strings.HasSuffix()
	strings.lastIndex()
	strings.Index()
	strings.Join(a[]string, sep string)
*/

func handleStr() {
	str := "_ hello the world"
	length := len(str)
	str1 := str + "!!!"
	arr1 := strings.Split(str, " ") //
	hasStr := strings.Contains(str, "llo")
	clone := strings.Clone(str)
	cv := strings.Compare("b", "a")
	rs := strings.Repeat(str, 2)
	hasChars := strings.ContainsAny(str, "m")
	count := strings.Count(str, "l")
	//strings.ToLowerSpecial("", str)
	upStr := strings.ToUpper(str)
	lowStr := strings.ToLower(upStr)
	trimStr := strings.Trim(str, "_")
	fmt.Println(length, str1, arr1, hasStr, clone, cv, rs, hasChars, count, upStr, lowStr, trimStr)
}
