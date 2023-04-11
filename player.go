package main

type Player struct {
	Name  string
	Hand  []Card
}

// getNextPlayer returns the next player in the turn sequence
func getNextPlayer(players []Player, currentPlayer *Player) *Player {
	currentPlayerIdx := -1
	for i := 0; i < len(players); i++ {
		if players[i].Name == currentPlayer.Name {
			currentPlayerIdx = i
			break
		}
	}
	if currentPlayerIdx == -1 {
		return nil
	}

	nextPlayerIdx := (currentPlayerIdx + 1) % len(players)
	return &players[nextPlayerIdx]
}

func reversePlayerOrder(players []Player) {
	for i, j := 0, len(players)-1; i < j; i, j = i+1, j-1 {
		players[i], players[j] = players[j], players[i]
	}
}