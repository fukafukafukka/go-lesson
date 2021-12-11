package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	cur, _ := os.Getwd()     // なぜかカレンとディレクトリがsrc/section9まで読み込まれない。
	fmt.Println(string(cur)) // /Users/fuka/Training/go

	// osパッケージでもファイルを読み込める。
	// ioutilはネットワーク関連でよく見かける。ex)パケット読み込み出力
	content, err := ioutil.ReadFile("src/section9/lesson72_ioutil.go")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(content))

	// byte配列のファイルへの書き込み
	// if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil {
	// log.Fatalln(err)
	// }

	// ネットワークからbyteを読み込んで一時的に記憶しておくのとかに良い
	r := bytes.NewBuffer([]byte("abc"))
	content2, _ := ioutil.ReadAll(r)
	// https://www.k-cube.co.jp/wakaba/server/ascii_code.html
	fmt.Println(content2)         // [97 98 99]
	fmt.Println(string(content2)) // abc
}
