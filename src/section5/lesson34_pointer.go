package main

import "fmt"

func lessonPointer(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	// 変数内の100が取得できる
	fmt.Println(n)
	// メモリのアドレスが取得できる
	fmt.Println(&n)
	// 以下は実行できない
	// fmt.Println(*n)

	// アドレスを変数に詰める。その際、型は「*int」というintegerのポイント型を使うことで変数に詰めることができる
	var p *int = &n
	// 変数内の番地が取得できる
	fmt.Println(p)
	// 変数内の番地が指す値が取得できる。
	fmt.Println(*p)

	var num int = 2
	// numのアドレスをメソッドに渡して、メソッド内でアドレスに対して値を代入する。
	lessonPointer(&num)
	fmt.Println(num)
}
