package main

import "fmt"

// 在循环中使用append函数构建一个由九个rune字符构成的slice
func main() {
	var r []rune
	for _, v := range "hello,世界" {
		r = append(r, v)
	}
	fmt.Printf("%q\n", r)
}
