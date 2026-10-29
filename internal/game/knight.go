package game

import "fmt"

const (
	KnightVal = 2
)

type knight struct {
	isWhite   bool
	val       int
	pieceRune rune
}

func newKnight(color bool) *knight {
	pieceRune := '\u265E'
	if color {
		pieceRune = '\u2658'
	}
	return &knight{
		isWhite:   color,
		val:       KnightVal,
		pieceRune: pieceRune,
	}
}

func (k *knight) getAvailableMoves(pos coords, board board) []coords {
	return []coords{}
}

func (k *knight) getPieceVal() int {
	return k.val
}

func (k *knight) getPieceColor() bool {
	return k.isWhite
}

func (k *knight) getPieceRune() rune {
	return k.pieceRune
}

func (k *knight) String() string {
	return fmt.Sprintf("Knight{color: %t}", k.isWhite)
}
