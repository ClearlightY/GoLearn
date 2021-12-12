package main

import "fmt"

func main() {
	// 定义变量/声明变量
	// 第一种： 指定变量类型， 声明后若不赋值， 使用默认值
	// int的默认值是0，String默认值是空串，小数默认值为0
	var i int
	fmt.Println("i = ", i)
	// 第二种： 根据值自行排定变量类型（类型推导)
	var num = 10.11
	fmt.Println("num=", num)

	// 第三种：省略var，注意 := 左侧的变量不应该是已经声明过的，否则会导致编译错误
	// 下面的方式等价var name string name = "tom"
	// := 的 : 不能省略，否则错误
	name := "tom"
	fmt.Println("name=", name)
}
