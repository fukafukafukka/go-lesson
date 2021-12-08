package main

import (
	"fmt"

	"mylib"
)

// GOPATHモードで紹介されているが、moduleモードが推奨になったため、こちらで対応する。
// https://go.dev/doc/tutorial/create-module
// cd mylib
// go mod init mylib
// cd ..
// go mod init section8
// go mod edit -replace mylib=./mylib
// go mod tidy
// go run .
func main() {
	person := mylib.PublicPerson{Name: "Mike", Age: 20}
	fmt.Println(person)

	// 以下は参照できる
	fmt.Println(mylib.PublicArg)

	// 以下は参照できない
	// fmt.Println(mylib.privateArg)
}
