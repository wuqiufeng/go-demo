package main

//func main() {
//	for i := 0; i < 5; i++ {
//		defer func() {
//			println(i)
//		}()
//	}
//}

//每轮迭代中生成一个局部变量：
//func main() {
//	for i := 0; i < 5; i++ {
//		i := i
//		defer func() {
//			println(i)
//		}()
//	}
//}

//通过函数参数传入：
func main() {
	for i := 0; i < 5; i++ {
		defer func(i int) {
			println(i)
		}(i)
	}
}
