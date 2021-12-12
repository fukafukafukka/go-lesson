package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mxk/go-sqlite/sqlite3"
)

// コネクションプール
var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

// # sqliteのインストール
// - brew install sqlite
// # Xcodeのインストール
// - https://developer.apple.com/xcode/
// # Xcodeがインストールされているか確認
// - xcode-select --install
// # sqliteライブラリのインストール(https://pkg.go.dev/github.com/mxk/go-sqlite/sqlite3#hdr-Installation)
// - go get github.com/mxk/go-sqlite/sqlite3
// # 使わないけど、import文へ上記モジュールを追加
// - import (_ "github.com/mxk/go-sqlite/sqlite3")
func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	// 1.テーブル作成
	cmd := `CREATE TABLE IF NOT EXISTS person(
				name STRING,
				age INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	// 実行したら下記コマンドでテーブルが作成されていることが確認できる
	// $ sqlite3 example.sql
	// $ .table

	// 2.テーブルへレコードをインサート
	cmd = `INSERT INTO person (name, age) VALUES (?, ?)`
	_, err = DbConnection.Exec(cmd, "Mike", 20)
	if err != nil {
		log.Fatalln(err)
	}
	// こちらも実行したら下記コマンドでテーブルの中身を確認できる
	// $ sqlite3 example.sql
	// $ select * from person

	// 3.投入したレコードを更新
	cmd = `UPDATE person SET age = ? WHERE name = ?`
	_, err = DbConnection.Exec(cmd, 25, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}

	// 4.レコードをマルチセレクト
	cmd = "SELECT * FROM person"
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	var pp []Person
	// 配列ppへQuery取得した要素を詰め込む
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	// 要素詰め込んだ配列ppの要素全て表示する
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

	// 5.シングルセレクト
	cmd = "SELECT * FROM person where age = ?"
	row := DbConnection.QueryRow(cmd, 20)
	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)

	// 6.削除
	cmd = "DELETE FROM person WHERE name = ?"
	_, err = DbConnection.Exec(cmd, "Nancy")
	if err != nil {
		fmt.Println(err)
	}
}
