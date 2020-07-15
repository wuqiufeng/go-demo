package main

//map遍历是顺序不固定
func main() {
	m := map[string]string{"1": "1", "2": "2", "3": "3"}
	for k, v := range m {
		println(k, v)
	}
}
