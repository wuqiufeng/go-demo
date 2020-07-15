package main

import (
	"fmt"
	"os"
	"runtime"
)

//Goroutine是协作式抢占调度，Goroutine本身不会主动放弃CPU：
//func main() {
//	runtime.GOMAXPROCS(1)
//	go func() {
//		for i := 0; i < 10; i++ {
//			fmt.Println(i)
//		}
//	}()
//	//占用CPU
//	for {
//	}
//}

//解决的方法是在for循环加入runtime.Gosched()调度函数：
//func main() {
//	runtime.GOMAXPROCS(1)
//	go func() {
//		for i := 0; i < 10; i++ {
//			fmt.Println(i)
//		}
//	}()
//	for {
//		runtime.Gosched()
//	}
//	fmt.Println("hello")
//}

//是通过阻塞的方式避免CPU占用：
func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		os.Exit(0)
	}()
	select {}
}
