package game

import "fmt"

const (
	QueenVal = 9
)

type queen struct {
	color     color
	val       int
	pieceRune rune
}

func newQueen(color color) *queen {
	var pieceRune rune
	switch color {
	case white:
		pieceRune = '\u2655'
	case black:
		pieceRune = '\u265B'
	}
	return &queen{
		color:     color,
		val:       QueenVal,
		pieceRune: pieceRune,
	}
}

func (q *queen) getAvailableMoves(pos coords, game *Game) []coords {
	return []coords{}
}

func (q *queen) getPieceVal() int {
	return q.val
}

func (q *queen) getPieceColor() color {
	return q.color
}

func (q *queen) getPieceRune() rune {
	return q.pieceRune
}

func (q *queen) String() string {
	return fmt.Sprintf("Queen{color: %v}", q.color)
}
