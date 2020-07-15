package main

import (
	"runtime"
)

//休眠并不能保证输出完整的字符串：
//func main() {
//	go println("hello")
//	time.Sleep(time.Second)
//}

//类似的还有通过插入调度语句：
func main() {
	go println("hello")
	runtime.Gosched()
}
