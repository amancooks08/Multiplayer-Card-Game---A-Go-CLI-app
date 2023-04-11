package main

import "fmt"


type Card struct {
	Rank string
	Suit string
}


// DrawCard draws a card from the deck
func (d *Deck) DrawCard() Card {
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}

// PlayCard plays a card from the player's hand
func (p *Player) PlayCard(cardIdx int, discardPile *[]Card) error {
	if cardIdx < 0 || cardIdx >= len(p.Hand) {
		return fmt.Errorf("invalid card index")
	}

	card := p.Hand[cardIdx]
	topCard := (*discardPile)[len(*discardPile)-1]
	if card.Rank != topCard.Rank && card.Suit != topCard.Suit {
		return fmt.Errorf("cannot play this card")
	}

	*discardPile = append(*discardPile, card)
	p.Hand = append(p.Hand[:cardIdx], p.Hand[cardIdx+1:]...)

	return nil
}

// DrawHand draws a hand of cards from the deck
func (p *Player) DrawHand(deck *Deck) {
	for i := 0; i < handSize; i++ {
		card := deck.DrawCard()
		p.Hand = append(p.Hand, card)
	}
}
