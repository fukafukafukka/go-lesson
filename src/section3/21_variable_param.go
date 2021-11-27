package main

import "fmt"

// 固定長引数
func staticParam(param1, param2 int) {
	fmt.Println(param1)
	fmt.Println(param2)
}

// 可変長引数
func variableParam(params ...int) {
	fmt.Println(len(params), params)
}

func main() {
	staticParam(10, 20)
	variableParam(10, 20, 30)

	// スライスも引数に使うことができる˜
	s := []int{1, 2, 3, 4, 5}
	variableParam(s...)
}
