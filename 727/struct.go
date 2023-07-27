package main

func main() {

}

/*
	Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。
	Go语言中可以使用type关键字来定义自定义类型。
	type myInt int
*/

/*
	类型别名
  type TypeAlias = Type
	我们之前见过的rune和byte就是类型别名，他们的定义如下：
	type byte = uint8
	type rune = int32
*/

/*
	结构体 表现事物的全部属性
    type 类型名 struct {
        字段名 字段类型
        字段名 字段类型
        …
    }
    1.类型名：标识自定义结构体的名称，在同一个包内不能重复。
    2.字段名：表示结构体字段名。结构体中的字段名必须唯一。
    3.字段类型：表示结构体字段的具体类型。
*/

/*
	结构体实例化
	只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
*/

/*
	匿名结构体
	在定义一些临时数据结构等场景下还可以使用匿名结构体
*/

func noNameStuct() {
	// 匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "小王子"
	user.Age = 18
}

/*
	创建指针类型结构体

*/
