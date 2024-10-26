package main

import "fmt"

func main() {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	m2 := make(map[string]int) // m2 == empty map
	var m3 map[string]int      // m3 == nil
	fmt.Println(m, m2, m3)
	fmt.Println(m2 == nil) // false
	fmt.Println(m3 == nil) // true

	fourValue := m["four"]
	fmt.Println(fourValue)

	if value, ok := m["one"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("value not found")
	}

	println("len:", len(m))

	delete(m, "one")
	if value, ok := m["one"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("value not found")
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
