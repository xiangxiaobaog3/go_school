package main

import (
	"fmt"
	"strconv"
)

type (
	byte int8
	ByteSize int64
)

const g int = 1
const h  = 'A'

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
)

func main() {
	H(1, 2, 3, 4)
	A()
	//fmt.Println(math.MinInt8)
	fmt.Println(b)

	_, _, c, d := 1, 2, 3, 4
	fmt.Println(c)
	fmt.Println(d)

	var a float32 = 1.1
	b := int(a)
	fmt.Println(b)

	var e int = 65
	f := strconv.Itoa(e)
	//e = strconv.Atoi(f)
	fmt.Println(f)
	fmt.Println(GB)

	z := 1
	var y *int =&z
	fmt.Println(y)	//fmt.Println(*y)

	if x:=1; x > 0 {
		fmt.Println(x)
	}

	//j := 1
	//for {
	//	j++
	//	if j > 3 {
	//		break
	//	}
	//	fmt.Println(j)
	//}
	//fmt.Println("Over")

	j := 1
	for j <= 3 {
		j++
		fmt.Println(j)
	}
	fmt.Println("Over")


	switch l := 1; {
	case l >= 0:
		fmt.Println("l=0")
		fallthrough
	case l >= 1:
		fmt.Println("l=1")
	default:
		fmt.Println("None")
	}

	LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1
			}
		}
	}
	fmt.Println("OK")

	var m [2]int
	fmt.Println(m)

	//n := [20]int{19: 1}
	//fmt.Println(n)

	//n := [...]int{0:1, 1:2, 2:3}
	//fmt.Println(n)
	n := [...]int{99: 1}
	var u *[100]int = &n
	fmt.Println(u)

	p, q := 1, 2
	o := [...]*int{&p, &q}
	fmt.Println(o)

}

func A() (a, b, c int) {
	a, b, c = 1, 2, 3
	return a, b, c
}

func H(a ...int) {
	fmt.Println(a)
}

var a int //

var b int = 321

