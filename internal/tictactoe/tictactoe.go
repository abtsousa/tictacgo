package tictactoe

import "fmt"

// Each Board is a 32-bit representation of each players moves
// where 0oXABC represents each row (A,B,C = 1st,2nd,3rd row; X = don't care)
// XPlays indicates whether the next player is X
type State struct {
	XBoard uint32
	OBoard uint32
	XPlays bool
}

const SCORE_PLAYER int = -1
const SCORE_AI int = +1

const CHAR_PLAYERX = "X"
const CHAR_PLAYERO = "O"

func (s *State) FreeStates() uint32 {
	return (^(s.XBoard | s.OBoard)) & 0777
}

func (s *State) Print() {
	// Iterate over each row
	for row := 0; row < 3; row++ {
		// Iterate over each column
		for col := 0; col < 3; col++ {
			// Calculate the position in the board
			pos := row*3 + col
			// Check if the position is occupied by X
			if (s.XBoard>>pos)&1 == 1 {
				fmt.Print(CHAR_PLAYERX)
			} else if (s.OBoard>>pos)&1 == 1 { // Check if the position is occupied by O
				fmt.Print(CHAR_PLAYERO)
			} else { // If the position is empty
				fmt.Print(" ")
			}
			// Print column separators
			if col < 2 {
				fmt.Print("|")
			}
		}
		// Print row separators
		fmt.Println()
		if row < 2 {
			fmt.Println("-----")
		}
	}
}
