package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else if num%3 == 0 {
		return "no, by 3"
	} else {
		return "no, else"
	}
}

func main() {
	x, y := 10, 0
	if x == 10 && y == 10 {
		fmt.Println("&&")
	}

	result := by2(10)
	if result == "ok" {
		fmt.Println("great")
	}

	// 上記は以下のようにまとめて書くこともできる
	if result_2 := by2(10); result_2 == "ok" {
		fmt.Println("great 2")
	}
}
