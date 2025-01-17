package main

import "fmt"

type Help interface {
	Print() string
}

type Person struct {
	name string
}

func (p *Person) Print() string {
	return "this is a person interface"
}

func main() {
	p := &Person{
		name: "douzengrui",
	}
	fmt.Println(p.name)
	fmt.Println(p.Print())
	p2 := Person{name: "1"}
	fmt.Println(p2.name)
	fmt.Println(p2.Print())

}
