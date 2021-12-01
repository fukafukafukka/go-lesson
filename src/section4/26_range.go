package main

import "fmt"

func main() {
	l := []string{"python", "go", "java"}

	// slice
	// for文
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	fmt.Println("--------")

	// slice
	// range文
	for i, v := range l {
		fmt.Println(i, v)
	}

	fmt.Println("--------")

	// slice
	// lのインデックス番号使わない時は「_」にしてあげないとコンパイルエラー
	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200, "grappe": 300}

	// map
	// range文
	for k, v := range m {
		fmt.Println(k, v)
	}

	// keyだけ取り出す
	for k := range m {
		fmt.Println(k)
	}

	// valueだけ取り出す
	for _, v := range m {
		fmt.Println(v)
	}
}
