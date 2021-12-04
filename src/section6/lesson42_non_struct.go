package main

import "fmt"

type MyInt int

// int型を返す場合は、MyInt型をint型に変換してあげないといけない。
func (i MyInt) Double() int {
	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", 1, 1)
	return int(i * 2)
}

// typeで宣言したMyInt型を返す場合は以下のようにする。
func (i MyInt) DoubleReturnType() MyInt {
	return i * 2
}

func main() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double())
	fmt.Println(myInt.DoubleReturnType())
}
