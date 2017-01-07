package main

import "fmt"
import "net"

const inter  = "en0"

func check(err error)  {
	if err != nil {
		panic(err)
	}
}

func main() {
	ifi, err := net.InterfaceByName(inter)
	check(err)
	addrs, err := ifi.Addrs()
	check(err)
	// _（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃
	// 由于 Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错, 在这种情况下, 可以使用_来丢弃不需要的返回值
	for _, a := range addrs {
		fmt.Printf("Interface %q, address %v\n", ifi.Name, a)
	}
}
