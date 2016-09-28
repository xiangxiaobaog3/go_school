package main

import "fmt"

func main() {
	s1 := make([]int, 3, 10)
	fmt.Println(len(s1), cap(s1))
}


