package main

import (
	"fmt"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
	// 名前を付けて使いたい場合には以下のように書く。
	// name "github.com/markcheno/go-talib"
	// 使わない場合には以下のように書く。
	// _ "github.com/markcheno/go-talib"
)

// 以下でgo-talibをgo.mod管理下に置く。(実態は次のディレクトリ配下にダウンロードされる。Users/fuka/go/pkg/mod/github.com)
// go get github.com/markcheno/go-talib
// 以下でgo-talibが依存しているtidyを同じくgo.mod管理下に置く。
// go mod tidy
// 以下で実行
// go run lesson64_third_party_package.go
func main() {
	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}
