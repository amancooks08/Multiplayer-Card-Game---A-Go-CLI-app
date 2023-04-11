# Multiplayer Card Game
This is a multiplayer card game implemented in Golang that supports up to 4 players and different types of cards, including number cards and action cards. The game follows the rules described below:

# Rules
1. Each player starts with a hand of 5 cards drawn from a standard deck of 52 cards.
2. The game starts with a discard pile, which initially contains one card drawn from the deck.
3. Players take turns playing cards from their hand, following the following rules:
    i) A player can only play a card if it matches either the suit or the rank of the top card on the discard pile.
   ii) If a player has an action card (Ace, King, Queen, or Jack), they can play it to trigger a special action (see below).
4. If a player cannot play a card, they must draw a card from the draw pile. If the draw pile is empty, the game ends in a draw and no player is declared a winner.
5. The game ends when one player runs out of cards, and that player is declared the winner.

# Action Cards
The game includes the following action cards:

Ace (A): Skip the next player in turn.
King (K): Reverse the sequence of players who play next.
Queen (Q): Add 2 cards to the draw pile. The next player must draw these cards and cannot play a Queen from their hand on that turn, even if available.
Jack (J): Add 4 cards to the draw pile. The next player must draw these cards and cannot play a Jack from their hand on that turn, even if available.

Note that action cards are not stackable, i.e., if a Queen or Jack is played, the next player must draw the specified number of cards and cannot play another Queen or Jack from their hand on that turn, even if available.

# How to Run

To run the game, follow these steps:

1. Clone the repository to your local machine.
2. Ensure that you have Golang installed on your machine.
3. Open a terminal and navigate to the directory where the repository is cloned.
4. Run the "go mod tidy" command to ensure you have all the required modules present.
5. Run the main game file using the go run command: go run main.go.
6. Follow the on-screen instructions to play the game. Use the keyboard to input your choices during the game.
7. Enjoy the multiplayer card game with your friends!


# Contributing
If you would like to contribute to the game, feel free to submit a pull request. Please make sure to follow the coding standards and include appropriate comments in your code. Also, update the README.md file with any necessary changes.

# License
This project is licensed under the MIT License, which allows for free use, modification, and distribution of the code. However, please acknowledge the original authors if you use or modify this project for your own purposes.



