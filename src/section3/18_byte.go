package main

import "fmt"

func main() {
	// 以下を参考にするとわかりやすい
	// ref) https://www.ascii-code.com
	b := []byte{72, 73}
	fmt.Println(b) // [72 73]
	fmt.Println(string(b)) // HI

	// HIに関してもbyteのスライスに入れれば72,73のバイト文字列に変換される
	c := []byte("HI")
	fmt.Println(c) // [72 73]
	fmt.Println(string(c)) // HI
}
