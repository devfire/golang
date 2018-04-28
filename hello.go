package main

import (
	"fmt"
	"log"
)

func main() {
	c, err := fmt.Println("hello world")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("bytes written", c)
}
