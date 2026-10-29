package game

import "fmt"

const (
	BishopVal = 3
)

type bishop struct {
	isWhite   bool
	val       int
	pieceRune rune
}

func newBishop(color bool) *bishop {
	pieceRune := '\u265D'
	if color {
		pieceRune = '\u2657'
	}
	return &bishop{
		isWhite:   color,
		val:       BishopVal,
		pieceRune: pieceRune,
	}
}

func (b *bishop) getAvailableMoves(pos coords, board board) []coords {
	return []coords{}
}

func (b *bishop) getPieceVal() int {
	return b.val
}

func (b *bishop) getPieceColor() bool {
	return b.isWhite
}

func (b *bishop) getPieceRune() rune {
	return b.pieceRune
}

func (b *bishop) String() string {
	return fmt.Sprintf("Bishop{color: %t}", b.isWhite)
}
