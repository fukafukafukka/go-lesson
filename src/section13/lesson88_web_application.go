package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"text/template"
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

// キャッシングできる
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// t, _ := template.ParseFiles(tmpl + ".html")
	// t.Execute(w, p)
	// renderTemplate(w, "view", p)

	// 上記を以下に書き換える
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// /view/test
	// 以下は引数で取得するよう修正
	// title := r.URL.Path[len("/view/"):] // rはクライアントから送られてきたリクエスト情報が入っている
	p, err := loadPage(title) // ローカルにtest.txtファイルが必要
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	// http.HandleFunc("/view/", viewHandler)// 下記に変更
	http.HandleFunc("/view/", makeHandler(viewHandler))
	// http.HandleFunc("/edit/", editHandler)// 下記に変更
	http.HandleFunc("/edit/", makeHandler(editHandler))
	// http.HandleFunc("/save/", saveHandler)// 下記に変更
	http.HandleFunc("/save/", makeHandler(saveHandler))

	// 第一引数の:の前を省略するとローカルホストになる
	// 第二引数のhandlerをnilとするとスラッシュやview以外のパスでアクセスした時にページNotFoundを表示してくれる
	log.Fatal(http.ListenAndServe(":8888", nil)) // サーバー起動中にエラーが発生した時にエラーが返るようになっているのでエラーログで出力するようにしておく。
}
