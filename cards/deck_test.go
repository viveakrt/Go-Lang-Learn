package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected 52 got %v", len(d))
	}

	if d[0] != "Ace Of Diamond" {
		t.Errorf("Expected Ace Of Diamond but got %v", d[0])
	}

	if d[len(d)-1] != "Joker Of Heart" {
		t.Errorf("Expected Joker Of Heart but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := openFromFile(("_decktesting"))
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
