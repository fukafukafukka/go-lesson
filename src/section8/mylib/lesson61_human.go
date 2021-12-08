package mylib

var PublicArg string = "public arg"
var privateArg string = "private arg"

// publicなストラクト(importして初期化できる)
type PublicPerson struct {
	// publicなフィールド
	Name string
	Age  int
}
