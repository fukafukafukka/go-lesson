package main

import "fmt"

type Vertex struct {
	X, Y int
}

func Area(v Vertex) int {
	return v.X * v.Y
}

// 値レシーバー
func (v Vertex) Area() int {
	return v.X * v.Y
}

// ポインタレシーバー
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Area(v))
	fmt.Println(v.Area())

	v.Scale(10)
	fmt.Println(v)
}
