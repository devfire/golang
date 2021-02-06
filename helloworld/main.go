package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hi there!")
	kuku()
}

func kuku() {
	foobar := true
	fmt.Println("current time is ", time.Now())
	fmt.Println("My favorite number is", add(rand.Intn(10), rand.Intn(10)))
	fmt.Println("Pi today is", foobar)

}
