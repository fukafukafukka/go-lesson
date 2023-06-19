package main

import "fmt"

func lessonPointer(x *int) {
	fmt.Println(x)  // アドレスが出力される
	fmt.Println(*x) // アドレスにある値が出力される
	*x = 1          // アドレスにある値を書き換えている
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
	lessonPointer(&num) // &num: intの2が入ったアドレス
	lessonPointer(p)    // p: intの100が入ったnのアドレス
	lessonPointer(nil)  // メソッド内で、xを出力すると<nil>というアドレスが表示されて、*xを出力すると参照先が無くpanicが起こる
	fmt.Println(num)
}
