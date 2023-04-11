package main

import "fmt"

func game() {
	//Initialize the deck
	deck := NewDeck()

	//Shuffle the deck
	deck.Shuffle()

	// Ask for number of players and add them to the players slice
	var numPlayers int
	fmt.Print("Enter number of players: ")
	fmt.Scanln(&numPlayers)

	// Check if the number of players is between 2 and 4
	if numPlayers > 4 || numPlayers < 2 {
		fmt.Println("Invalid number of players")
		return
	}
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

	//DiscardPile are the cards that have been played already.
	firstCard := deck.DrawCard()

	// If the first card is a Queen(Q) or a Jack(J), put the card back in the deck, reshuffle
	// it and draw another card.
	if firstCard.Rank == "Q" || firstCard.Rank == "J" {
		deck.PutCard(firstCard)
		deck.Shuffle()
		firstCard = deck.DrawCard()
	}
	discardPile := []Card{firstCard}

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

				err := players[i].PlayCard(cardIdx, &discardPile, &deck, players)
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
