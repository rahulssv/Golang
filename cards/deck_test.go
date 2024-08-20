package main

import (
	"os"
	"testing"
)

func TestNewDeck (t *testing.T){
	d := newDeck()

	if len(d) !=16 {
		t.Errorf("Expected deck length 16, but found %v", len(d))
	}

	if (d[0]!="Ace of Spades"){
		t.Errorf("Expected Ace of Spades, but found %v", d[0])
	}
}

func TestNewDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_deckTesting")
	deck:= newDeck()
	deck.saveToFile("_deckTesting")

	loadDeckFromFile:= newDeckFromFile("_deckTesting")

	if len(loadDeckFromFile)!=16{
		t.Errorf("Expected deck length 16, but found %v", len(loadDeckFromFile))
	}

	os.Remove("_deckTesting")

}