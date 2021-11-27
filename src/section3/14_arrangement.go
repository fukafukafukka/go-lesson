package main

import "fmt"

func main() {

	// 配列を宣言してから値を代入する方法
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	// 配列宣言時に値をセットする方法
	var b [2]int = [2]int{100, 200}
	fmt.Println(b)

	// 配列は宣言時にサイズを決定する必要があるので、後から変えられない。
	// b = append(b, 300) // 配列はサイズ以上にappendで追加できない。

	// 可変な配列を作りたいなら、次にやるスライスというものを使う必要がある。
	var c []int = []int{100, 200}
	c = append(c, 300)
	fmt.Println(c)
}
