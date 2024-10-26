package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"runtime"
	"strconv"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic("wrong score: " + strconv.Itoa(score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func div(a, b int) (int, int) {
	return a / b, a % b
}
func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return q, r
}
func forever() {
	for {
		fmt.Println("forever")
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with arguments: %d , %d \n", opName, a, b)
	return op(a, b)
}

func sumArgs(values ...int) int {
	sum := 0
	for value := range values {
		sum += value
	}
	return sum
}

func main() {
	a := 100
	if a > 100 {
		fmt.Println("a>100")
	} else if a == 100 {
		fmt.Println("a==100")
	} else {
		fmt.Println("a<100")
	}

	if b := 100; b > 100 {
		fmt.Println("b>100")
	} else if b == 100 {
		fmt.Println("b==100")
	} else {
		fmt.Println("b<100")
	}
	fmt.Println(eval(1, 2, "x"))
	fmt.Println(grade(100))
	fmt.Println(
		convertToBin(13), convertToBin(14))
	fmt.Println(div2(13, 2))
	q, r := div2(19, 2)
	fmt.Println(q, r)

	if result, err := eval(1, 2, "x"); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	var num int = 2
	var pa *int = &num
	*pa = 4
	fmt.Println(num)

	var arr1 [5]int
	var arr2 [3]int = [3]int{1, 2, 3}
	arr3 := [5]int{1, 2, 3, 4, 5}
	arr4 := [...]int{1, 2, 3, 4, 5}
	var grid [3][3]int = [3][3]int{{1, 2, 3}, {3, 4, 5}, {6, 7, 8}}
	fmt.Println(arr1, arr2, arr3, arr4, grid)
	for i := 0; i < len(arr4); i++ {
		fmt.Println(arr4[i])
	}
	for _, v := range arr4 {
		fmt.Println(v)
	}
}
