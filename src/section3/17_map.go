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

	// var m3 map[string]int
	// 以下で、エラーになる。理由:初期化されていない(=入れる対象のメモリがない)ため、nilに値を入れようとしてしまっている
	// m3["pc"] = 5000
	// fmt.Println(m3)

	// 以下両方とも初期化されていないので、nilとなっていることに気を付ける
	var m3 map[string]int
	if m3 == nil {
		fmt.Println("Nil")
	}
	var s []int
	if s == nil {
		fmt.Println("Nil")
	}
}
