package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	// stringをintにする場合にはライブラリを使う必要がある
	var s string = "14"
	i, _ := strconv.Atoi(s) // Atoi=「アスキーto integer」という意味
	// i, err := strconv.Atoi(s) // errは今回返ってこないし使わないので、この使い方はしない。
	fmt.Printf("%T %v", i, i)

	// 復習だが、文字列とアスキーは関係性があるので以下のようにライブラリを使わずとも変換ができた
	h := "Hello World"
	fmt.Println(string(h[0]))
	// まとめるとintとfloat、stringとアスキーコードはすぐ変換できる。
}
