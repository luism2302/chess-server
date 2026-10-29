package game

import "fmt"

const (
	KingVal = 20
)

type king struct {
	isWhite     bool
	val         int
	isFirstMove bool
	pieceRune   rune
}

func newKing(color bool) *king {
	pieceRune := '\u265A'

	if color {
		pieceRune = '\u2654'
	}
	return &king{
		isWhite:     color,
		val:         KingVal,
		isFirstMove: true,
		pieceRune:   pieceRune,
	}
}

func (k *king) getAvailableMoves(pos coords, board board) []coords {
	return []coords{}
}

func (k *king) getPieceVal() int {
	return k.val
}

func (k *king) getPieceColor() bool {
	return k.isWhite
}

func (k *king) getPieceRune() rune {
	return k.pieceRune
}

func (k *king) String() string {
	return fmt.Sprintf("King{color: %t, isFirstMove: %t}", k.isWhite, k.isFirstMove)
}
