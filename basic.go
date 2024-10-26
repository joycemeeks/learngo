package main

import (
	"fmt"
	"math"
)

var aa = []int{1, 2, 3, 4, 5}
var bb = []int{6, 7, 8, 9, 10}

var (
	str  = "hello world"
	str2 = "hello world2"
)

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(a, b, c)
}

func enums() {
	const (
		cpp = iota
		_
		_
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

/*func main() {
	fmt.Println("Hello World")

	var a int = 20
	var s string = "hello"
	fmt.Println(a, s)
	// 也可以定义多个值
	var c, d int = 1, 2
	fmt.Println(c, d)
	//var flag, num, str = true, 2, "hello"
	flag, num, str := true, 2, "hello"
	fmt.Println(flag, num, str)

	var (
		str1 = "hello world"
		str2 = "hello world2"
	)
	fmt.Println(str1, str2)

	triangle()
	enums()
}*/
