package game

import "fmt"

const (
	QueenVal = 9
)

type queen struct {
	isWhite   bool
	val       int
	pieceRune rune
}

func newQueen(color bool) *queen {
	pieceRune := '\u265B'
	if color {
		pieceRune = '\u2655'
	}
	return &queen{
		isWhite:   color,
		val:       QueenVal,
		pieceRune: pieceRune,
	}
}

func (q *queen) getAvailableMoves(pos coords, board board) []coords {
	return []coords{}
}

func (q *queen) getPieceVal() int {
	return q.val
}

func (q *queen) getPieceColor() bool {
	return q.isWhite
}

func (q *queen) getPieceRune() rune {
	return q.pieceRune
}

func (q *queen) String() string {
	return fmt.Sprintf("Queen{color: %t}", q.isWhite)
}
