package main

import (
	"fmt"
	"math"
)

// 变量除了可以在全局声明中初始化，也可以在 init 函数中初始化。
// 这是一类非常特殊的函数，它不能够被人为调用，而是在每个包完成初始化后自动执行，并且执行优先级比 main 函数高。
var Pi float64

func init() {
	Pi = 4 * math.Atan(1)
}
func main() {
	fmt.Println(Pi)
}
