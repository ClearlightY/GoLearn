package main

import "fmt"

func main() {
	// 该案例演示golang如何一次性声明多个变量
	var n1, n2, n3 int
	fmt.Println("n1=", n1, " n2=", n2, " n3=", n3)

	// 一次性声明多个变量方式2
	var n4, name, n5 = 100, "tom", 888
	fmt.Println("n4=", n4, " name=", name, " n5=", n5)

	// 一次性声明多个变量方式3
	n6, name2, n7 := 100, "jack", 888
	fmt.Println("n6=", n6, "name2=", name2, "n7=", n7)
}
