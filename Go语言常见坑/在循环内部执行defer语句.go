package main

import (
	"log"
	"os"
)

//defer在函数退出时才能执行，在for执行defer会导致资源延迟释放：
//func main() {
//	for i := 0; i < 5; i++ {
//		f, err := os.Open("/Users/ace/controlconfig.conf")
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer f.Close()
//	}
//}

//解决的方法可以在for中构造一个局部函数，在局部函数内部执行defer：
func main() {
	for i := 0; i < 5; i++ {
		func() {
			f, err := os.Open("/Users/ace/controlconfig.conf")
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
		}()
	}
}
