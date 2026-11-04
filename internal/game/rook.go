package game

import "fmt"

const (
	RookVal = 5
)

type rook struct {
	color     color
	val       int
	pieceRune rune
}

func newRook(color color) *rook {
	var pieceRune rune
	switch color {
	case white:
		pieceRune = '\u2656'
	case black:
		pieceRune = '\u265C'
	}
	return &rook{
		color:     color,
		val:       RookVal,
		pieceRune: pieceRune,
	}
}

func (r *rook) getAvailableMoves(pos coords, game *Game) []coords {
	return []coords{}
}

func (r *rook) getPieceVal() int {
	return r.val
}

func (r *rook) getPieceColor() color {
	return r.color
}

func (r *rook) getPieceRune() rune {
	return r.pieceRune
}

func (r *rook) String() string {
	return fmt.Sprintf("Rook{color: %v}", r.color)
}
