// David Walshe
// 16/04/2020

package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck lenght of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	err := os.Remove("_decktesting")
	if err != nil {
		fmt.Println("_decktesting does not exist")
	}

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := d.newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in loaded deck, got %v", len(loadedDeck))
	}

	err = os.Remove("_decktesting")
	if err != nil {
		fmt.Println("_decktesting does not exist")
	}
}
