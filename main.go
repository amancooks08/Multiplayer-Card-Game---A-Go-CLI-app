package main

import (
	game "github.com/amancooks08/Multiplayer-Card-Game/cardgame"
)
func main() {
	for {
		exit := game.Game()
		if exit {
			break
		}
	}
}
