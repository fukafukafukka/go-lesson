package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// サーバーサイドのDBと見立てる
var DB = map[string]string{
	"User1Key": "User1Secret",
	"User2Key": "User2Secret",
}

// サーバーサイドと見立てる
func Server(apiKey, sign string, data []byte) {
	apiSecret := DB[apiKey] // DBからシークレットを取得する

	// クライアントサイドと同様にhmac生成する
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedMAC := hex.EncodeToString(h.Sum(nil))

	fmt.Println(sign == expectedMAC)
}

// hmac: apiでサーバーにアクセスする際のauthenticationでヘッダーに含めるというケースが多い
func main() {
	// main内をクライアントサイドと見立てる
	const apiKey = "User1Key"
	const apiSecret = "User1Secret"

	// https://pkg.go.dev/crypto/hmac
	h := hmac.New(sha256.New, []byte(apiSecret)) // sha256というアルゴリズムにkeyを入れてハッシュを作成
	data := []byte("data")
	h.Write(data)                          // 生成したハッシュに送りたいメッセージを入れて
	sign := hex.EncodeToString(h.Sum(nil)) // nilを加えてハッシュを作成する

	fmt.Println(sign) // usecaseとして、この値をサーバーサイドに送って認証してもらう必要がある

	Server(apiKey, sign, data)
}
