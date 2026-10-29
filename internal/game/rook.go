package game

import "fmt"

const (
	RookVal = 5
)

type rook struct {
	isWhite     bool
	val         int
	isFirstMove bool
	pieceRune   rune
}

func newRook(color bool) *rook {
	pieceRune := '\u265C'
	if color {
		pieceRune = '\u2656'
	}
	return &rook{
		isWhite:     color,
		val:         RookVal,
		isFirstMove: true,
		pieceRune:   pieceRune,
	}
}

func (r *rook) getAvailableMoves(pos coords, board board) []coords {
	return []coords{}
}

func (r *rook) getPieceVal() int {
	return r.val
}

func (r *rook) getPieceColor() bool {
	return r.isWhite
}

func (r *rook) getPieceRune() rune {
	return r.pieceRune
}

func (r *rook) String() string {
	return fmt.Sprintf("Rook{color: %t, isFirstMove: %t}", r.isWhite, r.isFirstMove)
}
