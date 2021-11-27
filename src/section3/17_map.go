package main

import "fmt"

func main() {
	// map[key]value{"key":value, ...}
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)

	// 上書き
	m["banana"] = 300

	// 新規追加
	m["new"] = 500

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	// mapの初期化もできる
	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	// 以下は入れる対象のメモリがないのでエラーになる
	// var m3 map[string]int
	// m3["pc"] = 5000
	// fmt.Println(m3)

	// 以下両方ともnilとなることに気を付ける
	var m3 map[string]int
	if m3 == nil {
		fmt.Println("Nil")
	}
	var s []int
	if s == nil {
		fmt.Println("Nil")
	}
}
