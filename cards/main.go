package main

import "fmt"

func main() {
	cards := newDeck()
	//fmt.Println(cards)

	//cards.print()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
	fmt.Println(hand.toString())
}
