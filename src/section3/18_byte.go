package main

import "fmt"

func main() {
	// 以下を参考にするとわかりやすい
	// ref) https://www.ascii-code.com
	b := []byte{72, 73}
	fmt.Println(b)
	fmt.Println(string(b))

	// HIに関してもbyteのスライスに入れれば72,73のバイト文字列に変換される
	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))
}
