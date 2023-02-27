package main

import "fmt"

// complex64和complex128，分别对应float32和float64两种浮点数精度
// 内置的complex函数用于构建复数，内建的real和imag函数分别返回复数的实部和虚部
func main() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
}
