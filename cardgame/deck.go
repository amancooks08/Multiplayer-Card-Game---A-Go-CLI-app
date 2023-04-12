package cardgame

import (
	"math/rand"
	"time"
)

type Deck []Card


// NewDeck creates a new deck of cards
func NewDeck() Deck {
	var deck Deck
	for _, rank := range ranks {
		for _, suite := range suites {
			card := Card{Rank: rank, Suite: suite}
			deck = append(deck, card)
		}
	}
	return deck
}

// Shuffle shuffles the deck
func (d *Deck) Shuffle() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range *d {
		j := rand.Intn(i + 1)
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
}


// PutCard puts a card on the deck
func (d *Deck) PutCard(card Card) {
	*d = append(*d, card)
}
