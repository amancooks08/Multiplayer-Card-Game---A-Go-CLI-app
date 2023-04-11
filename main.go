package main

import (
	"fmt"
)

// DrawCard draws a card from the deck
func (d *Deck) DrawCard() Card {
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}

func NewPlayer(name string) Player {
	return Player{
		Name: name,
		Hand: []Card{},
	}
}

// DrawHand draws a hand of cards from the deck
func (p *Player) DrawHand(deck *Deck) {
	for i := 0; i < handSize; i++ {
		card := deck.DrawCard()
		p.Hand = append(p.Hand, card)
	}
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

func main() {
	//Initialize the deck
	deck := NewDeck()

	//Shuffle the deck
	deck.Shuffle()

	// Ask for number of players and add them to the players slice
	var numPlayers int
	fmt.Print("Enter number of players: ")
	fmt.Scanln(&numPlayers)

	// Ask for their names and add them to the players slice
	players := []Player{}
	for i := 0; i < numPlayers; i++ {
		var name string
		fmt.Printf("Enter name for player %d: ", i+1)
		fmt.Scanln(&name)
		players = append(players, NewPlayer(name))
	}

	for i := range players {
		players[i].DrawHand(&deck)
	}

	discardPile := []Card{deck.DrawCard()}

	for {
		for i := range players {
			fmt.Printf("%s's Turn\n", players[i].Name)
			fmt.Printf("Hand: %+v\n", players[i].Hand)
			fmt.Printf("Top Card: %+v\n", discardPile[len(discardPile)-1])

			var cardIdx int
			for {
				fmt.Print("Select a card to play (or enter -1 to draw a card): ")
				fmt.Scanln(&cardIdx)
				if cardIdx == -1 {
					players[i].Hand = append(players[i].Hand, deck.DrawCard())
					break
				}

				err := players[i].PlayCard(cardIdx, &discardPile)
				if err == nil {
					break
				} else {
					fmt.Println("Error:", err)
				}
			}
		}

		// Check if any player has an empty hand
		for i := range players {
			if len(players[i].Hand) == 0 {
				fmt.Printf("%s wins!\n", players[i].Name)
				return
			}
		}
	}
}
