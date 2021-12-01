package main

import (
	"fmt"
	"os"
)

func doDefer() {
	defer fmt.Println("world from doDefer")

	fmt.Println("hello from doDefer")
}

func useCase() {
	// ファイルの閉じ忘れ防止のためによく使う
	file, _ := os.Open("./lesson.go")
	defer file.Close()

	// 上記のように閉じる操作を示した後、以下の具体的な処理に移る
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(data)
}

func main() {
	defer fmt.Println("world")

	doDefer()

	fmt.Println("hello")

	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	useCase()
}
