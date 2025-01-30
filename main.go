package main

import (
	"fmt"
	"strconv"
	"strings"
)

import ttt "github.com/abtsousa/tictacgoe/internal/tictactoe"

func main() {
	// Initialize the game state
	initialState := &ttt.State{
		XBoard: 0,
		OBoard: 0,
		XPlays: true, // X (human) plays first
	}

	// Create the root node of the game tree
	root := ttt.NewNode(initialState, 0, true)

	// Current node in the game tree
	currentNode := &root

	// Game loop
	for {
		// Print the current board
		fmt.Println("Current Board:")
		currentNode.Print()

		// Check if the game is over
		result := ttt.Eval(currentNode)
		switch result {
		case 'X':
			fmt.Println("You win!")
		case 'O':
			fmt.Println("AI wins!")
		case 'D':
			fmt.Println("It's a draw!")
		}

		if result != ' ' {
			break
		}

		// Human's turn (X)
		if currentNode.XPlays() {
			fmt.Print("Your turn (X). Enter cell (1-9): ")
			var input string
			fmt.Scanln(&input)

			// Parse the input
			cell, err := strconv.Atoi(strings.TrimSpace(input))
			if err != nil || cell < 1 || cell > 9 {
				fmt.Println("Invalid input. Please enter a number between 1 and 9.")
				continue
			}

			// Convert cell to bit position
			bitPos := uint32(1 << (cell - 1))

			// Play the move
			nextNode, err := ttt.Play(currentNode, bitPos)
			if err != nil {
				fmt.Println("Invalid move. Cell already occupied.")
				continue
			}

			// Update the current node
			currentNode = nextNode
		} else {
			// AI's turn (O)
			fmt.Println("AI's turn (O)...")
			currentNode = currentNode.BestMove
		}
	}

	fmt.Println("Game over. Thanks for playing!")
}
