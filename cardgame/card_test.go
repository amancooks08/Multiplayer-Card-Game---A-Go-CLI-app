package cardgame

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDrawCard(t *testing.T) {
	// Create a new deck with some cards
	d := Deck{Card{Rank: "A", Suite: "Spades"}, Card{Rank: "2", Suite: "Hearts"}, Card{Rank: "10", Suite: "Diamonds"}}

	// Test case 1: Draw a card from a non-empty deck
	card := d.DrawCard()
	if card.Rank != "A" || card.Suite != "Spades" {
		t.Errorf("Expected card to be Ace of Spades, but got %s of %s", card.Rank, card.Suite)
	}

	// Test case 2: Verify deck length after drawing a card
	if len(d) != 2 {
		t.Errorf("Expected deck length to be 2, but got %d", len(d))
	}

	// Test case 3: Draw a card from an empty deck
	d = Deck{}
	card = d.DrawCard()
	if card.Rank != "" || card.Suite != "" {
		t.Errorf("Expected card to be empty, but got %s of %s", card.Rank, card.Suite)
	}
}

func TestDrawHand(t *testing.T) {
	type fields struct {
		Name string
		Hand []Card
	}
	type args struct {
		deck *Deck
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Draw a hand",
			fields: fields{
				Name: "Player 1",
				Hand: []Card{},
			},
			args: args{
				deck: &Deck{Card{Rank: "A", Suite: "Spades"}, Card{Rank: "2", Suite: "Hearts"}, Card{Rank: "10", Suite: "Diamonds"}, Card{Rank: "K", Suite: "Clubs"}, Card{Rank: "Q", Suite: "Spades"}},
			},
		},

		{
			name: "Draw a hand from a deck with less than 5 cards",
			fields: fields{
				Name: "Player 1",
				Hand: []Card{},
			},
			args: args{
				deck: &Deck{Card{Rank: "A", Suite: "Spades"}, Card{Rank: "2", Suite: "Hearts"}, Card{Rank: "10", Suite: "Diamonds"}, Card{Rank: "K", Suite: "Clubs"}, Card{Rank: "Q", Suite: "Spades"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Player{
				Name: tt.fields.Name,
				Hand: tt.fields.Hand,
			}
			p.DrawHand(tt.args.deck)
		})
	}
}

func Test_PlayCard(t *testing.T) {
	type fields struct {
		Name string
		Hand []Card
	}
	type args struct {
		cardIdx     int
		discardPile *[]Card
		deck        *Deck
		players     []Player
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		expectedErr error
	}{
		{
			name: "Play an unmatched card",
			fields: fields{
				Name: "Player 1",
				Hand: []Card{{Rank: "A", Suite: "Spades"}, {Rank: "2", Suite: "Hearts"}, {Rank: "10", Suite: "Diamonds"}, {Rank: "K", Suite: "Clubs"}, {Rank: "Q", Suite: "Spades"}},
			},
			args: args{
				cardIdx:     1,
				discardPile: &[]Card{{Rank: "2", Suite: "Diamonds"}, {Rank: "4", Suite: "Clubs"}},
				deck:        &Deck{Card{Rank: "A", Suite: "Spades"}, Card{Rank: "2", Suite: "Hearts"}, Card{Rank: "10", Suite: "Diamonds"}, Card{Rank: "K", Suite: "Clubs"}, Card{Rank: "J", Suite: "Spades"}},
				players:    []Player{{Name: "Player 1", Hand: []Card{{Rank: "A", Suite: "Spades"}, {Rank: "2", Suite: "Hearts"}, {Rank: "10", Suite: "Diamonds"}, {Rank: "K", Suite: "Clubs"}, {Rank: "Q", Suite: "Spades"}}}, {Name: "Player 2", Hand: []Card{{Rank: "A", Suite: "Spades"}, {Rank: "2", Suite: "Hearts"}, {Rank: "10", Suite: "Diamonds"}, {Rank: "K", Suite: "Clubs"}, {Rank: "Q", Suite: "Spades"}}}},
			},
			expectedErr: errors.New("cannot play this card"),
		},
		{
			name: "Play a card with an invalid index",
			fields: fields{
				Name: "Player 1",
				Hand: []Card{{Rank: "A", Suite: "Spades"}, {Rank: "2", Suite: "Hearts"}, {Rank: "10", Suite: "Diamonds"}, {Rank: "K", Suite: "Clubs"}, {Rank: "Q", Suite: "Spades"}},
			},
			args: args{
				cardIdx:     10,
				discardPile: &[]Card{{Rank: "2", Suite: "Diamonds"}, {Rank: "4", Suite: "Clubs"}},
				deck:        &Deck{Card{Rank: "A", Suite: "Spades"}, Card{Rank: "2", Suite: "Hearts"}, Card{Rank: "10", Suite: "Diamonds"}, Card{Rank: "K", Suite: "Clubs"}, Card{Rank: "J", Suite: "Spades"}},
				players:    []Player{{Name: "Player 1", Hand: []Card{{Rank: "A", Suite: "Spades"}, {Rank: "2", Suite: "Hearts"}, {Rank: "10", Suite: "Diamonds"}, {Rank: "K", Suite: "Clubs"}, {Rank: "Q", Suite: "Spades"}}}, {Name: "Player 2", Hand: []Card{{Rank: "A", Suite: "Spades"}, {Rank: "2", Suite: "Hearts"}, {Rank: "10", Suite: "Diamonds"}, {Rank: "K", Suite: "Clubs"}, {Rank: "Q", Suite: "Spades"}}}},
			},
			expectedErr: errors.New("invalid card index"),
		},
		
		{
			name: "Play a card with a valid index",
			fields: fields{
				Name: "Player 1",
				Hand: []Card{{Rank: "A", Suite: "Spades"}, {Rank: "2", Suite: "Hearts"}, {Rank: "10", Suite: "Diamonds"}, {Rank: "1", Suite: "Clubs"}, {Rank: "Q", Suite: "Spades"}},
			},
			args: args{
				cardIdx:     3,
				discardPile: &[]Card{{Rank: "2", Suite: "Diamonds"}, {Rank: "4", Suite: "Clubs"}},
				deck:        &Deck{Card{Rank: "A", Suite: "Spades"}, Card{Rank: "2", Suite: "Hearts"}, Card{Rank: "10", Suite: "Diamonds"}, Card{Rank: "K", Suite: "Clubs"}, Card{Rank: "J", Suite: "Spades"}},
				players:    []Player{{Name: "Player 1", Hand: []Card{{Rank: "A", Suite: "Spades"}, {Rank: "2", Suite: "Hearts"}, {Rank: "10", Suite: "Diamonds"}, {Rank: "K", Suite: "Clubs"}, {Rank: "Q", Suite: "Spades"}}}, {Name: "Player 2", Hand: []Card{{Rank: "A", Suite: "Spades"}, {Rank: "2", Suite: "Hearts"}, {Rank: "10", Suite: "Diamonds"}, {Rank: "K", Suite: "Clubs"}, {Rank: "Q", Suite: "Spades"}}}},
			},
			expectedErr: nil,
		},

		{
			name: "Play a card with zero index, which is also an Ace",
			fields: fields{
				Name: "Player 1",
				Hand: []Card{{Rank: "A", Suite: "Clubs"}, {Rank: "2", Suite: "Hearts"}, {Rank: "10", Suite: "Diamonds"}, {Rank: "1", Suite: "Clubs"}, {Rank: "Q", Suite: "Spades"}},
			},
			args: args{
				cardIdx:     0,
				discardPile: &[]Card{{Rank: "2", Suite: "Diamonds"}, {Rank: "4", Suite: "Clubs"}},
				deck:        &Deck{Card{Rank: "A", Suite: "Spades"}, Card{Rank: "2", Suite: "Hearts"}, Card{Rank: "10", Suite: "Diamonds"}, Card{Rank: "K", Suite: "Clubs"}, Card{Rank: "J", Suite: "Spades"}},
				players:    []Player{{Name: "Player 1", Hand: []Card{{Rank: "A", Suite: "Spades"}, {Rank: "2", Suite: "Hearts"}, {Rank: "10", Suite: "Diamonds"}, {Rank: "K", Suite: "Clubs"}, {Rank: "Q", Suite: "Spades"}}}, {Name: "Player 2", Hand: []Card{{Rank: "A", Suite: "Spades"}, {Rank: "2", Suite: "Hearts"}, {Rank: "10", Suite: "Diamonds"}, {Rank: "K", Suite: "Clubs"}, {Rank: "Q", Suite: "Spades"}}}},
			},
			expectedErr: nil,
		},

		{	
			name: "Play a card with a valid index which is also a Jack",
			fields: fields{
				Name: "Player 1",
				Hand: []Card{{Rank: "A", Suite: "Spades"}, {Rank: "J", Suite: "Clubs"}, {Rank: "10", Suite: "Diamonds"}, {Rank: "1", Suite: "Clubs"}, {Rank: "Q", Suite: "Spades"}},
			},
			args: args{
				cardIdx:     1,
				discardPile: &[]Card{{Rank: "2", Suite: "Diamonds"}, {Rank: "4", Suite: "Clubs"}},
				deck:        &Deck{Card{Rank: "A", Suite: "Spades"}, Card{Rank: "2", Suite: "Hearts"}, Card{Rank: "10", Suite: "Diamonds"}, Card{Rank: "K", Suite: "Clubs"}, Card{Rank: "J", Suite: "Spades"}},
				players:    []Player{{Name: "Player 1", Hand: []Card{{Rank: "A", Suite: "Spades"}, {Rank: "2", Suite: "Hearts"}, {Rank: "10", Suite: "Diamonds"}, {Rank: "K", Suite: "Clubs"}, {Rank: "Q", Suite: "Spades"}}}, {Name: "Player 2", Hand: []Card{{Rank: "A", Suite: "Spades"}, {Rank: "2", Suite: "Hearts"}, {Rank: "10", Suite: "Diamonds"}, {Rank: "K", Suite: "Clubs"}, {Rank: "Q", Suite: "Spades"}}}},
			},
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Player{
				Name: tt.fields.Name,
				Hand: tt.fields.Hand,
			}
			// Use require to check for errors
			err := p.PlayCard(tt.args.cardIdx, tt.args.discardPile, tt.args.deck, tt.args.players)
			require.Equal(t, tt.expectedErr, err)
		})

	}
}
