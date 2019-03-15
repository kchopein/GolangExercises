package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 20 {
		t.Errorf("Expected 20 but got %v!", len(deck))
	}
}

func TestDeckToFile(t *testing.T) {
	d := newDeck()

	d.saveToFile("_deckTesting")

	ld, _ := newDeckFromFile("_deckTesting")

	if len(d) != len(ld) {
		t.Error("Decks are not the same size!")
	}

	os.Remove("_deckTesting")
}
