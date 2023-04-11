package main

import "fmt"

func game() (exit bool){
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
		if name == "" || name == " "{
			fmt.Println("Invalid name")
			return false
		}
		players = append(players, NewPlayer(name))
	}

	for i := range players {
		players[i].DrawHand(&deck)
	}

	//DiscardPile are the cards that have been played already.
	firstCard := deck.DrawCard()

	// If the first card is a Queen(Q) or a Jack(J), put the card back in the deck, reshuffle
	// it and draw another card.
	if firstCard.Rank == "Q" || firstCard.Rank == "J" || firstCard.Rank == "K" || firstCard.Rank == "A"{
		deck.PutCard(firstCard)
		deck.Shuffle()
		firstCard = deck.DrawCard()
	}
	//DiscardPile are the cards that have been played already.
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
					if len(players[i].Hand) == 0 {
						fmt.Printf("%s wins!\n", players[i].Name)
						return true
					}
					break
				} else {
					fmt.Println("Error:", err)
				}
			}
		}

		// Check if the deck is empty, then declare that no one wins
		if len(deck) == 0 {
			fmt.Println("No one wins! :P")
			return true
		}
	}
}
