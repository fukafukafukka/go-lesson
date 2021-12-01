package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// goの場合には以下のように呼び出し元でエラーハンドリングを実行する
	file, err := os.Open("/Users/fuka/Training/go/src/section4/30_error_handling.go")
	if err != nil {
		log.Fatalln("Error")
	}
	defer file.Close()

	data := make([]byte, 100)
	// err定義済みだがcountが未定義なので「:」でイニシャライズができる
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error")
	}
	// dataにファイルの中身が入るの要調査
	fmt.Println(count, string(data))

	if err = os.Chdir("test"); err != nil {
		log.Fatalln("Error")
	}
}
