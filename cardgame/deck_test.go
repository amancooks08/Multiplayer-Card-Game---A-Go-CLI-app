package cardgame

import (
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	tests := []struct {
		name string
		want Deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeck(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPutCard(t *testing.T) {
	// Create a new deck
	deck := Deck{}

	// Put a card into the deck
	card := Card{"Ace", "Spades"}
	deck.PutCard(card)

	// Check if the card is in the deck
	found := false
	for _, c := range deck {
		if c == card {
			found = true
			break
		}
	}

	// Assert that the card is in the deck
	if !found {
		t.Errorf("Expected card %+v to be in deck, but not found", card)
	}
}