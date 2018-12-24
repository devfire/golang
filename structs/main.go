package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	me := person{firstName: "Igz", lastName: "K"}

	fmt.Println(me.lastName)
}
