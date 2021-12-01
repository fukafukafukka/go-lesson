package main

import "fmt"

// Golangでは、そもそもpanicを実行しないようにエラーハンドリングしてプログラムを書くことを推奨している。
func thirdPartyConnectDB() {
	// 自分で例外を投げられる
	panic("Unable to connect database!")
}

func save() {
	defer func() {
		// panicが実行された場合に、そのpanicを以下のrecoverでキャッチできる
		s := recover()
		fmt.Println(s)
		// 以上のrecover処理より引き続き処理を継続できる。
	}()

	// panicが起きる可能性のある処理の前に、deferでrecoverを定義しておかないといけない。
	thirdPartyConnectDB()
}

func main() {
	save()
	fmt.Println("ok?")
}
