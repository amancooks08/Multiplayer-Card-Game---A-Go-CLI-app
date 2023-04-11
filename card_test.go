package main

import (
	"testing"
)

func TestDrawCard(t *testing.T) {
	// Create a new deck with some cards
	d := Deck{Card{Rank: "Ace", Suite: "Spades"}, Card{Rank: "2", Suite: "Hearts"}, Card{Rank: "10", Suite: "Diamonds"}}

	// Test case 1: Draw a card from a non-empty deck
	card := d.DrawCard()
	if card.Rank != "Ace" || card.Suite != "Spades" {
		t.Errorf("Expected card to be Ace of Spades, but got %s of %s", card.Rank, card.Suite)
	}

	// Test case 2: Verify deck length after drawing a card
	if len(d) != 2 {
		t.Errorf("Expected deck length to be 2, but got %d", len(d))
	}

	// Test case 3: Draw a card from an empty deck
	d = Deck{}
	card = d.DrawCard()
	if card.Rank != "" || card.Suite != "" {
		t.Errorf("Expected card to be empty, but got %s of %s", card.Rank, card.Suite)
	}
}
