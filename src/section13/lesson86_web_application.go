package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// /view/test
	title := r.URL.Path[len("/view/"):] // rはクライアントから送られてきたリクエスト情報が入っている
	p, _ := loadPage(title)             // ローカルにtest.txtファイルが必要
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)

	// 第一引数の:の前を省略するとローカルホストになる
	// 第二引数のhandlerをnilとするとスラッシュやview以外のパスでアクセスした時にページNotFoundを表示してくれる
	log.Fatal(http.ListenAndServe(":8888", nil)) // サーバー起動中にエラーが発生した時にエラーが返るようになっているのでエラーログで出力するようにしておく。
}
