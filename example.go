package main

import (
	"fmt"
)

var a int = 10 //全局变量

//函数example
func example(x int) int {
	var a int = 5 //函数内变量
	a = x + a     //操作的是函数内变量
	return a      //返回结果
}

func main() {
	var a int = 10
	fmt.Println("变量的地址: %x\n", &a)
	fmt.Println(example(3)) //打印结果
}
