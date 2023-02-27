package main

//new函数
//表达式new(T)将创建一个T类型的匿名变量，
//初始化为T类型的零值，然后返回变量地址，返回的指针类型为*T
func newInt() *int {
	return new(int)
}
func newInt2() *int {
	var dummy int
	return &dummy
}

func main() {

}
