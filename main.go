package main

import "fmt"

func main() {
	a := func ()  {
		fmt.Println("Func A")
	}
	a()

	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

func A()  {
	fmt.Println("Func A")
}

func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func (y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}