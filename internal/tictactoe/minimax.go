package tictactoe

import (
	"fmt"
)

type Node struct {
	val      int8
	depth    int8
	isMax    bool
	elem     *State
	children []*Node
	BestMove *Node
}

func NewNode(s *State, d int8, isMax bool) Node {

	valid_states := getValidStates(s)

	var children []*Node
	for _, valid_state := range valid_states {
		n := NewNode(valid_state, d+1, !isMax)
		children = append(children, &n)
	}

	n := Node{
		val:      0,
		depth:    d,
		isMax:    isMax,
		elem:     s,
		children: children,
		BestMove: nil,
	}

	Minimax(&n)
	return n
}

func (n *Node) Print() {
	n.elem.Print()
}

func (n *Node) XPlays() bool {
	return n.elem.XPlays
}

func Minimax(n *Node) int8 {
	if n.BestMove != nil {
		return n.val
	}
	eval := Eval(n)
	if eval == 'O' {
		return n.depth - 10
	} else if eval == 'X' {
		return 10 - n.depth
	}

	if len(n.children) == 0 {
		n.val = 0
		return n.val
	}

	var bestMove *Node

	if n.isMax {
		max := int8(-100)
		for _, child := range n.children {
			eval := Minimax(child)
			if eval > max {
				max = eval
				bestMove = child
			}
		}
		n.val = max
	} else {
		min := int8(100)
		for _, child := range n.children {
			eval := Minimax(child)
			if eval < min {
				min = eval
				bestMove = child
			}
		}
		n.val = min
	}
	n.BestMove = bestMove
	return n.val
}

func Play(n *Node, cell uint32) (*Node, error) {
	s := n.elem
	if s.FreeStates()&cell == 0 {
		return nil, fmt.Errorf("Invalid move!")
	}
	var board uint32
	if s.XPlays {
		board = s.XBoard | cell
	} else {
		board = s.OBoard | cell
	}

	for _, child := range n.children {
		if (s.XPlays && board == child.elem.XBoard) ||
			(!s.XPlays && board == child.elem.OBoard) {
			return child, nil
		}
	}
	return nil, fmt.Errorf("Error updating the board.")
}

func Eval(n *Node) rune {
	s := n.elem
	eval_pos := func(board uint32, pos uint32) bool {
		return (board&pos == pos)
	}

	eval_board := func(board uint32) bool {
		return eval_pos(board, 0o700) || // Horizontal
			eval_pos(board, 0o070) ||
			eval_pos(board, 0o007) ||
			eval_pos(board, 0o444) || // Vertical
			eval_pos(board, 0o222) ||
			eval_pos(board, 0o111) ||
			eval_pos(board, 0o421) || // Diagonal
			eval_pos(board, 0o124)
	}

	if eval_board(s.OBoard) {
		return 'O'
	} else if eval_board(s.XBoard) {
		return 'X'
	} else if s.FreeStates() == 0 {
		return 'D'
	}

	return ' '
}

func getValidStates(s *State) []*State {
	var valid_states []*State
	free_states := s.FreeStates()
	for i := 0; i < 9; i++ {
		digit := uint32(1 << i)
		if free_states&digit != 0 {
			if s.XPlays {
				valid_states = append(valid_states, &State{s.XBoard | digit, s.OBoard, false})
			} else {
				valid_states = append(valid_states, &State{s.XBoard, s.OBoard | digit, true})
			}
		}
	}

	return valid_states
}
