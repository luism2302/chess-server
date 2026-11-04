package game

import "fmt"

const (
	KingVal = 20
)

type king struct {
	color     color
	val       int
	pieceRune rune
}

func newKing(color color) *king {
	var pieceRune rune
	switch color {
	case white:
		pieceRune = '\u2654'
	case black:
		pieceRune = '\u265A'
	}

	return &king{
		color:     color,
		val:       KingVal,
		pieceRune: pieceRune,
	}
}

func (k *king) getAvailableMoves(pos coords, game *Game) []coords {
	return []coords{}
}

func (k *king) getPieceVal() int {
	return k.val
}

func (k *king) getPieceColor() color {
	return k.color
}

func (k *king) getPieceRune() rune {
	return k.pieceRune
}

func (k *king) String() string {
	return fmt.Sprintf("King{color: %v}", k.color)
}
