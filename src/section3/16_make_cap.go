package main

import "fmt"

func main() {
	// makeを使ってスライスを宣言することができる
	// n := make([]型, 長さ, キャパシティ)
	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	fmt.Println("-------------------------------")
	n = append(n, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	fmt.Println("-------------------------------")
	// スライスなので実は5個以上要素を追加できる
	n = append(n, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	fmt.Println("-------------------------------")

	// 引数1つだけだと長さもキャパシティも同一のスライスになる
	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)
	fmt.Println("-------------------------------")

	// こちらは要素無しのスライスをメモリに確保する
	b := make([]int, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
	fmt.Println(b)
	fmt.Println("-------------------------------")

	// こちらはメモリに登録しない
	var c []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)
	fmt.Println(c)
	fmt.Println("-------------------------------")

	c = make([]int, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
}
