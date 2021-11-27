package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello " + "World")
	fmt.Println("Hello World"[0])         // 72とアスキーコードが表示される
	fmt.Println(string("Hello World"[0])) // これで文字列が表示される

	var s string = "Hello World"
	// s[0] = "x"// これだとエラーになるので以下のようにライブラリの「strings」を使う。
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)

	// 他にも以下のようなstringsの使い方がある。
	var t bool = strings.Contains(s, "World")
	fmt.Println(t)

	// 開業のやり方は以下の2つある
	fmt.Println("Test\n" + "Test")
	fmt.Println("----")
	fmt.Println(`Test
Test`)
}
