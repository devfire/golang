package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("First card must be Ace of Spades but got %v", d[0])
	}
}

func TestSaveToFile(t *testing.T) {
	os.Remove("_decktest")

	deck := newDeck()

	deck.saveToFile("_decktest")

	loadedDeck := newDeckFromFile("_decktest")

	expectedCards := 16
	if len(loadedDeck) != expectedCards {
		t.Errorf("Expected %v cards, got %v", expectedCards, len(loadedDeck))
	}
}
