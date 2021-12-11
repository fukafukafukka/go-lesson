package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// 以下の一連の流れでgetリクエスト実行できる
	// resp, _ := http.Get("http://example.com")
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	// Parseはurlの値が正しいかどうかチェックするもの->空白とか入っているとエラー出る
	base, _ := url.Parse("http://example.com/syouryaku/sareru")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint) // http://example.com/test?a=1&b=2

	// リクエストの作成
	req, _ := http.NewRequest("GET", endpoint, nil)
	// req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("password")))// パスワードとかはPOSTリクエストで送る

	// リクエストにヘッダーを設定
	req.Header.Add("If-None-Match", `W/"wyzzy"`) // https://developer.mozilla.org/ja/docs/Web/HTTP/Headers/If-None-Match

	// リクエストにクエリを設定
	q := req.URL.Query()
	fmt.Println(q) // map[a:[1] b:[2]]
	q.Add("c", "3&%")
	fmt.Println(q)                   // map[a:[1] b:[2] c:[3&%]]
	fmt.Println(q.Encode())          // a=1&b=2&c=3%26%25
	req.URL.RawFragment = q.Encode() // エンコードしておく
	fmt.Println(req)

	// クライアント生成 -> クライアントで作成しておいたリクエスト実行 -> ioutilでレスポンスbodyを取得
	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	// レスポンスのbyte[]をstringに変換して確認
	fmt.Println(string(body))
}
