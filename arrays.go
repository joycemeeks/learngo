package main

import "fmt"

func printArray(a *[5]int) {
	a[0] = 100
	for i, v := range a {
		fmt.Println(i, v)
	}
}
func main() {
	var arr1 [5]int
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("start print arr1")
	printArray(&arr1)
	fmt.Println("start print arr2")
	printArray(&arr2)
	fmt.Println("arr1 and arr2")
	fmt.Println(arr1, arr2)
}
