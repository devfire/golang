package main

import (
	"fmt"
	"log"
)

func main() {
	const mystr = "hello world"
	c, err := fmt.Println(mystr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("bytes written", c)

}
