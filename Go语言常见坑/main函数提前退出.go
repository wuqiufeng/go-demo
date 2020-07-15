package main

//后台Goroutine无法保证完成任务。
func main() {
	go println("hello")
}
