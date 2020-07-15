package main

//var msg string
//var done bool
//
//func setup() {
//	msg = "hello, world"
//	done = true
//}
//
////因为在不同的Goroutine，main函数中无法保证能打印出 hello, world :
//func main() {
//	go setup()
//	for !done {
//	}
//	println(msg)
//}

//解决的办法是用显式同步：
var msg string
var done = make(chan bool)

func setup() {
	msg = "hello, world"
	done <- true
}
func main() {
	go setup()
	<-done
	println(msg)
}
