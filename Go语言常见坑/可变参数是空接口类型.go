package main

import "fmt"

//当参数的可变参数是空接口类型时，传入空接口的切片时需要注意参数展开的问题
//不管是否展开，编译器都无法发现错误，但是输出是不同的
func main() {
	var a = []interface{}{1, 2, 3}
	fmt.Println(a)    // [1 2 3]
	fmt.Println(a...) // 1 2 3
}
