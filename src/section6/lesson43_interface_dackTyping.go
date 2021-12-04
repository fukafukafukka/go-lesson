package main

import "fmt"

type Human interface {
	Say()
}

type Human2 interface {
	Say() string
	SetMr()
}

type Person struct {
	Name string
}

type Person2 struct {
	Name string
}

func (p Person) Say() {
	fmt.Println(p.Name)
}

func (p Person2) Say() string {
	return p.Name
}

func (p *Person2) SetMr() {
	p.Name = "Mr." + p.Name
}

// use case of interface
// ダックタイピングと呼ばれるもの
func DriveCar(h Human2) {
	if h.Say() == "Mr.Tison" {
		fmt.Println("Run!")
	} else {
		fmt.Println("Get out!")
	}
}

func main() {
	// Humanというインターフェースの中にPersonで定義したものを代入する場合、
	// Person自体にinterfaceで定義されているメソッドを定義しないといけない。
	var mike Human = Person{"Mike"}
	mike.Say()

	// ポインタレシーバーとしてポインタが受けつけられないといけない場合、アドレス(&)を渡してあげないといけない。
	var tison Human2 = &Person2{"Tison"}
	tison.SetMr()
	fmt.Println(tison.Say())
	DriveCar(tison)

	var david Human2 = &Person2{"David"}
	david.SetMr()
	fmt.Println(david.Say())
	DriveCar(david)
}
