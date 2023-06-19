package main

import "fmt"

func main() {
	var n int = 100
	fmt.Println(n)

	var p *int = &n
	// 変数のポインタが入っていることがわかる
	fmt.Println(p)

	var q *int = new(int)
	// 変数のポインタが入っていることがわかる
	fmt.Println(q)

	var q2 *int
	// 変数のポインタが入っていないことがわかる
	fmt.Println(q2)

	// new宣言の変数のポインタは計算できる
	*q++
	fmt.Println(*q)

	// newしていない変数は計算できない
	// q2++
}
