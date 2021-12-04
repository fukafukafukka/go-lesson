package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// structの標準出力の表示内容を変えたい場合は「String」メソッドを作成する
func (p Person) String() string {
	return fmt.Sprintf("My name is %v.", p.Name)
}

func main() {
	mike := Person{"Mike", 22}
	fmt.Println(mike)
}
