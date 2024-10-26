package main

import "fmt"

func main() {
	var s []int // 此时定义了一个slice类型数据，但是还没有值，但是go语言有一个特点，一旦定义了变量，就有zero value，zero value for slice is nil，go语言中没有null，使用nil表示没有
	// 这里虽然没有分配数值，但是却可以直接使用
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	// 确定值的情况
	s1 := []int{1, 2, 3}
	// 指定长度，没有指定内容的情况
	s2 := make([]int, 16)
	// 指定长度和容量的情况
	s3 := make([]int, 16, 32)

	copy(s2, s1)
	fmt.Println(s2)
	fmt.Println(s3)
	// 删除元素3
	// 下表从0开始，这里跳过下标时2的元素，s2[3:]...对应的时可变类型参数，获取每一个元素
	// 元素的长度会-1
	s2 = append(s2[:2], s2[3:]...)
	fmt.Println(s2)
}
