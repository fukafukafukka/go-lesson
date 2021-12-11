package main

import "fmt"

const (
	c1 = iota
	// c2 = iota
	c2 // 上記をこのように省略可能
	c3 // 同様に省略
)

const (
	_ = iota
	// シフト演算
	KB int = 1 << (10 * iota)
	MB
	GB
)

func main() {
	fmt.Println(c1, c2, c3)

	fmt.Println(KB, MB, GB)
}
