package main

import (
	"fmt"
)

func foo() {
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)

	// 以下のように型を調べて出力させることもできる。Printfなことに注意
	fmt.Printf("%T\n", xi)
	fmt.Printf("%T\n", xf64)

	// 再度代入する場合は「:」をつけない。
	xi = 2
	fmt.Println(xi)
	// xi := 2 // エラーになる
	// var i int = 2 // エラーになる
}

func main() {
	// var i int = 1
	// var f64 float64 = 1.2
	// var s string = "test"
	// var t bool = true
	// var f bool = false
	// fmt.Println(i, f64, s, t, f)

	// 上記を以下のようにvarをまとめて書ける。var宣言の場合に限り、main()の外に配置できる。
	var (
		i    int     = 1
		f64  float64 = 1.2
		s    string  = "test"
		t, f bool    = true, false
	)
	fmt.Println(i, f64, s, t, f)

	// 以下の関数内で、varを使わずに変数宣言することもできる。変数に入れるときに自動的に型が決まる
	foo()
}
