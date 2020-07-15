package main

import "fmt"

//在函数调用参数中，数组是值传递，无法通过修改数组类型的参数返回结果
func main() {
	x := [3]int{1, 2, 3}
	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)
	fmt.Println(x)
}
