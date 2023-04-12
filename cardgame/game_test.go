package cardgame

import "testing"

func Test_game(t *testing.T) {
	tests := []struct {
		name string
		numPlayers int
		players []Player
		expectedOutput string
	}{
		{
			name:           "Invalid number of players",
			numPlayers:     1,
			players:    []Player{},
			expectedOutput: "Invalid number of players",
		},
		{
			name:           "Valid number of players",
			numPlayers:     3,
			players:    []Player{
				{
					Name: "Player 1",
					Hand: []Card{},
				},
				{
					Name: "Player 2",
					Hand: []Card{},
				},
			},
			expectedOutput: "", // No error expected
		},

		{
			name:           "Invalid number of players",
			numPlayers:     5,
			players:    []Player{
				{
					Name: "Player 1",
					Hand: []Card{},
				},
				{
					Name: "Player 2",
					Hand: []Card{},
				},
				{
					Name: "Player 3",
					Hand: []Card{},
				},
				{
					Name: "Player 4",
					Hand: []Card{},
				},
				{
					Name: "Player 5",
					Hand: []Card{},
				},
			},
			expectedOutput: "Invalid number of players",
		},

		{
			name:           "Invalid name",
			numPlayers:     3,
			players:    []Player{
				{
					Name: "Player 1",
					Hand: []Card{},
				},
				{
					Name: "Player 2",
					Hand: []Card{},
				},
				{
					Name: "",
					Hand: []Card{},
				},
			},
			expectedOutput: "Invalid name",
		},

		{
			name:           "Valid name",
			numPlayers:     3,
			players:    []Player{
				{
					Name: "Player 1",
					Hand: []Card{},
				},
				{
					Name: "Player 2",
					Hand: []Card{},
				},
				{
					Name: "Player 3",
					Hand: []Card{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Game()
		})
	}
}