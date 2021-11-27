package main

import "fmt"

// 最初に実行される関数
func init() {
	fmt.Println("Init!")
}

func bazz() {
	fmt.Println("Bazz")
}

func main() {
	bazz()
	fmt.Println("Hello world!", "TEST TEST")
}
