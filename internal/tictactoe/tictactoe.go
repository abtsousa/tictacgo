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

// Helper functions for the minimax implementation
func IsTerminal(s *State) bool {
	return EvaluateState(s) != ' ' || s.FreeStates() == 0
}

func Utility(s *State) int {
	switch EvaluateState(s) {
	case 'X':
		return -1
	case 'O':
		return 1
	default:
		return 0
	}
}

func GetSuccessors(s *State) []*State {
	var successors []*State
	freeStates := s.FreeStates()

	for i := 0; i < 9; i++ {
		digit := uint32(1 << i)
		if freeStates&digit != 0 {
			newState := &State{
				XBoard: s.XBoard,
				OBoard: s.OBoard,
				XPlays: !s.XPlays,
			}
			if s.XPlays {
				newState.XBoard |= digit
			} else {
				newState.OBoard |= digit
			}
			successors = append(successors, newState)
		}
	}
	return successors
}

func EvaluateState(s *State) rune {
	evalPos := func(board uint32, pos uint32) bool {
		return (board&pos == pos)
	}

	evalBoard := func(board uint32) bool {
		return evalPos(board, 0o700) || // Horizontal
			evalPos(board, 0o070) ||
			evalPos(board, 0o007) ||
			evalPos(board, 0o444) || // Vertical
			evalPos(board, 0o222) ||
			evalPos(board, 0o111) ||
			evalPos(board, 0o421) || // Diagonal
			evalPos(board, 0o124)
	}

	if evalBoard(s.OBoard) {
		return 'O'
	} else if evalBoard(s.XBoard) {
		return 'X'
	} else if s.FreeStates() == 0 {
		return 'D'
	}
	return ' '
}

func Play(s *State, cell uint32) (*State, error) {
	if s.FreeStates()&cell == 0 {
		return nil, fmt.Errorf("invalid move: cell already occupied")
	}

	newState := &State{
		XBoard: s.XBoard,
		OBoard: s.OBoard,
		XPlays: !s.XPlays,
	}

	if s.XPlays {
		newState.XBoard |= cell
	} else {
		newState.OBoard |= cell
	}

	return newState, nil
}
