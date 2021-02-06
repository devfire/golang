package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	me := person{firstName: "Igz", lastName: "Kay"}

	//me.firstName = "Alex"
	me.print()
}

func (p *person) print() {
	fmt.Printf("%+v", (*p))
}
