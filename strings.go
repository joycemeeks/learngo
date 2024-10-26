package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello这里是图书馆"
	fmt.Println(str)
	// 输出23，获取字节的数量，一个汉字占用3个字节
	fmt.Println(len(str))
	// 使用utf8.RuneCountInString获得字符数量，输出11
	fmt.Println(utf8.RuneCountInString(str))

	// 遍历字节，输出字节对应的ascII码，48 65 6C 6C 6F E8 BF 99 E9 87 8C E6 98 AF E5 9B BE E4 B9 A6 E9 A6 86
	for _, ch := range []byte(str) {
		fmt.Printf("%X ", ch)
	}
	fmt.Println()
	// 遍历字符串中的每个字符，位置是字节的位置，注意中文字符占用三个字节
	// 这里的 ch 是 int32 类型，就相当于一个 rune
	// (0 48) (1 65) (2 6C) (3 6C) (4 6F) (5 8FD9) (8 91CC) (11 662F) (14 56FE) (17 4E66) (20 9986)
	for i, ch := range str {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	// 转换字符遍历输出
	// H  e  l  l  o  这  里  是  图  书  馆
	bytes := []byte(str)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf(" %c ", ch)
	}
	fmt.Println()

	// 转rune数组遍历输出字符
	// (0 H) (1 e) (2 l) (3 l) (4 o) (5 这) (6 里) (7 是) (8 图) (9 书) (10 馆)
	for i, ch := range []rune(str) {
		fmt.Printf("(%d %c) ", i, ch)
	}
}
