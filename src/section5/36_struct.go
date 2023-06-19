package main

import "fmt"

// 以下のように構造体を表現することもできる
// type Vertex struct {
// X, Y int
// S    string
// }

type Vertex struct {
	// 構造体の変数が大文字ならpublicを表す
	X int
	Y int
	S string
}

// 元のVertexの値は書き変わらない
func changeVertexX(v Vertex) {
	v.X = 1000
}

// 元のVertexの値が書き変わる
func changeVertexX2(v *Vertex) {
	v.X = 1000
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)

	// 以下で書き換えができる
	v.X = 100
	fmt.Println(v)

	// 片方の値だけ入れて宣言することもできる
	v2 := Vertex{Y: 200}
	fmt.Println(v2)

	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Printf("%T %v\n", v4, v4)

	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5) // 構造体の場合はnewしない場合、nilにならない点に注意

	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6)

	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7)

	v8 := Vertex{1, 2, "test"}
	changeVertexX(v8)
	fmt.Println(v8)

	changeVertexX2(&v8)
	fmt.Println(v8)
}
