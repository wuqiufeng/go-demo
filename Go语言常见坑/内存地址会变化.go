package main

import (
	"runtime"
	"unsafe"
)

//Go语言中对象的地址可能发生变化，因此指针不能从其它非指针类型的值生成：
//当内存发送变化的时候，相关的指针会同步更新，但是非指针类型的uintptr不会做 同步更新。
//同理CGO中也不能保存Go对象地址。
func main() {
	var x int = 42
	var p uintptr = uintptr(unsafe.Pointer(&x))
	runtime.GC()
	var px *int = (*int)(unsafe.Pointer(p))
	println(*px)
}
