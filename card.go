package main

import "fmt"


type Card struct {
	Rank string
	Suite string
}


// DrawCard draws a card from the deck
func (d *Deck) DrawCard() Card {
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}

// PlayCard plays a card from the player's hand
func (p *Player) PlayCard(cardIdx int, discardPile *[]Card, deck *Deck, players []Player) error {
	if cardIdx < 0 || cardIdx >= len(p.Hand) {
		return fmt.Errorf("invalid card index")
	}

	card := p.Hand[cardIdx]
	topCard := (*discardPile)[len(*discardPile)-1]
	if card.Rank != topCard.Rank && card.Suite != topCard.Suite {
		return fmt.Errorf("cannot play this card")
	}

	// Check for special action cards
	switch card.Rank {
		case "A": // Ace - Skip the next player in turn
			fmt.Printf("%s played Ace - Skipping the next player's turn\n", p.Name)

		case "K": // King - Reverse the sequence of who plays next
			fmt.Printf("%s played King - Reversing the turn sequence\n", p.Name)
			// Reverse the player order
			reversePlayerOrder(players)

		case "Q": // Queen - +2
			fmt.Printf("%s played Queen - Drawing 2 cards for the next player\n", p.Name)
			nextPlayer := getNextPlayer(players, p)
			for i := 0; i < 2; i++ {
				nextPlayer.Hand = append(nextPlayer.Hand, deck.DrawCard())
			}

		case "J": // Jack - +4
			fmt.Printf("%s played Jack - Drawing 4 cards for the next player\n", p.Name)
			nextPlayer := getNextPlayer(players, p)
			for i := 0; i < 4; i++ {
				nextPlayer.Hand = append(nextPlayer.Hand, deck.DrawCard())
			}
	}
	*discardPile = append(*discardPile, card)
	if cardIdx == 0{
		p.Hand = p.Hand[1:]
	} else {
		p.Hand = append(p.Hand[:cardIdx], p.Hand[cardIdx+1:]...)
	}
	
	return nil
}

// DrawHand draws a hand of cards from the deck
func (p *Player) DrawHand(deck *Deck) {
	for i := 0; i < handSize; i++ {
		card := deck.DrawCard()
		p.Hand = append(p.Hand, card)
	}
}
