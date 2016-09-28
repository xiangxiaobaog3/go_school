package main

import (
	"fmt"
)

func main() {
	var s []int
	s = []int{1, 2, 3}
	fmt.Println(s)

	a := s[0]
	fmt.Println(a)

	s[1] = 4
	fmt.Println(s)

	b := make([]int, 3, 4)
	fmt.Println(len(b), cap(b))
	fmt.Printf("b: %p %d %d %[1]v\n", b, len(b), cap(b))
	fmt.Println(b)

	b1 := append(b, 1)
	fmt.Printf("b1: %p %d %d %[1]v\n", b1, len(b1), cap(b1))

	b1[0] = 9
	fmt.Printf("b: %p %d %d %[1]v\n", b, len(b), cap(b))
	fmt.Printf("b1: %p %d %d %[1]v\n", b1, len(b1), cap(b1))

	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(len(s1), cap(s1), s1)
	fmt.Printf("s1: %p %d %d %[1]v\n", s1, len(s1), cap(s1))

	s5 := append(s1, 6)
	fmt.Printf("s5: %p %d %d %[1]v\n", s5, len(s5), cap(s5))

	s5[0] = 10

	s2 := s1[0:3]
	fmt.Println(len(s2), cap(s2), s2)
	fmt.Printf("s2: %p %d %d %[1]v\n", s2, len(s2), cap(s2))

	s2 = append(s2, 6, 7)
	fmt.Printf("s1: %p %d %d %[1]v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %p %d %d %[1]v\n", s2, len(s2), cap(s2))

	s2 = append(s2, 8)
	fmt.Printf("s1: %p %d %d %[1]v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %p %d %d %[1]v\n", s2, len(s2), cap(s2))

	s3 := s1[2:5]
	fmt.Println(len(s3), cap(s3), s3)

	s4 := s1[2:4]
	fmt.Println(len(s4), cap(s4), s4)

	s4[0] = 6
	s1[3] = 7
	fmt.Println(len(s1), cap(s1), s1)
	fmt.Println(len(s2), cap(s2), s2)
	fmt.Println(len(s3), cap(s3), s3)
	fmt.Println(len(s4), cap(s4), s4)

	s7 := []int{1, 2, 3, 4, 5}
	fmt.Println("s7:", s7)

	s8 := make([]int, len(s1))
	fmt.Println("s8:", s8)
	copy(s8[1:3], s7[2:4])
	fmt.Println("s8:", s8)

	s8[0] = 10
	fmt.Println("s7:", s7)
	fmt.Println("s8:", s8)
}
