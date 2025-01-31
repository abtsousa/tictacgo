package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/abtsousa/minimax-go"
	ttt "github.com/abtsousa/tictacgo/tictactoe"
)

func main() {

	for {
		playGame()
		fmt.Print("Do you want to play again? (y/n) ")
		var input string
		fmt.Scanln(&input)
		if input == "n" || input == "N" {
			break
		}
	}
}

func playGame() {

	var humanPlays bool
	for {
		// Ask who starts
		fmt.Print("Do you want to go first? (y/n) ")
		var input string
		fmt.Scanln(&input)

		if input == "y" || input == "Y" {
			fmt.Println("You will go first.")
			humanPlays = true
			break
		} else if input == "n" || input == "N" {
			fmt.Println("The computer will go first.")
			humanPlays = false
			break
		} else {
			fmt.Println("Invalid input. Please enter 'y' or 'n'.")
		}
	}

	// Initialize the game state
	initialState := &ttt.State{
		XBoard: 0,
		OBoard: 0,
		XPlays: humanPlays, // X (human) plays first
	}

	// Initialize the minimax algorithm
	game := minimax.Make(
		initialState,
		ttt.IsTerminal,
		ttt.Utility,
		ttt.GetSuccessors,
		!humanPlays,
	)

	// Current state in the game
	currentState := initialState

	// Game loop
	for {
		// Print the current board
		fmt.Println("Current Board:")
		currentState.Print()

		// Check if the game is over
		result := ttt.EvaluateState(currentState)
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
		if currentState.XPlays {
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
			nextState, err := ttt.Play(currentState, bitPos)
			if err != nil {
				fmt.Println("Invalid move. Cell already occupied.")
				continue
			}

			// Update the current node
			currentState = nextState
		} else {
			// AI's turn (O)
			fmt.Println("AI's turn (O)...")
			currentState = game.Solve(*currentState)
		}
	}

	fmt.Println("Game over. Thanks for playing!")
}
