package main

func main() {
	//cards := newDeck()
	//fmt.Println(cards)

	//cards.print()

	//hand, remainingCards := deal(cards, 5)

	//hand.print()
	//remainingCards.print()
	//fmt.Println(cards.toString())

	//cards.saveToFile("temp_cards.txt")

	//cards := newDeckFromFile("temp_cards.txt")

	//cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}
