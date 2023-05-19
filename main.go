package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) eat() {
	h.Name = "Jiro"
	h.Age = 20

}

func main() {
	h := Human{Name: "Ichiro", Age: 10}
	// 通常
	fmt.Printf("%sは%d歳です\n", h.Name, h.Age)

	// *の使い方
	h.eat()
	fmt.Printf("%sは%d歳です\n", h.Name, h.Age)

	// &の使い方
	u := &h
	u.Name = "Saburo"
	u.Age = 30
	fmt.Printf("%sは%d歳です", u.Name, u.Age)
}
