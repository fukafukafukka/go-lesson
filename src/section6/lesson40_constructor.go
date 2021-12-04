package main

import "fmt"

type Vertex struct {
	x, y int
}

func Area(v Vertex) int {
	return v.x * v.y
}

// 値レシーバー
func (v Vertex) Area() int {
	return v.x * v.y
}

// ポインタレシーバー
func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

// コンストラクター
// 戻り値は構造体のポインタ(アドレスの番地を格納している変数)であることに注意
func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Area(v))
	fmt.Println(v.Area())

	v.Scale(10)
	fmt.Println(v)

	v2 := New(3, 4)
	v2.Scale(10)
	fmt.Println(v.Area())
}
