package main

import (
	"errors"
	"fmt"
)

func Bar() (err error) {
	return errors.New("bar error")
}

func Foo() (err error) {
	if err := Bar(); err != nil {
		return
	}
	return
}

//在局部作用域中，命名的返回值内同名的局部变量屏蔽
func main() {
	err := Foo()
	fmt.Println("err:", err)
}
