package main

import (
	"fmt"
)

func menu() (exit bool) {
	fmt.Println("************************************")
	fmt.Println("1. New Game")
	fmt.Println("2. Exit")
	var choice string
	for {
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case "1":
			fmt.Println("Starting a new game...")
			game()

		case "2":
			fmt.Println("Exiting the game...")
			return true

		default:
			fmt.Println("Invalid choice. Please try again.")

		}
	}
}
