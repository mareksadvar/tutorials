package main

import "fmt"

type Person struct {
	Name string
	Lastname string
}

func (p *Person) Introduce () {
	fmt.Println("Hello " + p.Name)
}

func main() {
	person := &Person{
		Name : "Max",
		Lastname: "Mustermann",
	}
	person.Introduce()
}
