package main

import "fmt"

// 变量使用的注意事项
func main() {
	// 该区域的数据值可以在同一类型范围内不断变化
	var i int = 10
	i = 30
	i = 50
	fmt.Println("i=", i)
	//i = 1.2 // i原本为int类型，此处错误是因为不能改变数据类型
}
