package main

//This file contains all the required constants and structs used to run the game

const (
	deckSize = 52
	handSize = 5
)

var (
	ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
	suits = []string{"Hearts", "Diamonds", "Clubs", "Spades"}
)

type Player struct {
	Name  string
	Hand  []Card
	Score int
}
