package game

import "fmt"

const (
	BishopVal = 3
)

type bishop struct {
	color     color
	val       int
	pieceRune rune
}

func newBishop(color color) *bishop {
	var pieceRune rune
	switch color {
	case white:
		pieceRune = '\u2657'
	case black:
		pieceRune = '\u265D'
	}

	return &bishop{
		color:     color,
		val:       BishopVal,
		pieceRune: pieceRune,
	}
}

func (b *bishop) getAvailableMoves(pos coords, game *Game) []coords {
	return []coords{}
}

func (b *bishop) getPieceVal() int {
	return b.val
}

func (b *bishop) getPieceColor() color {
	return b.color
}

func (b *bishop) getPieceRune() rune {
	return b.pieceRune
}

func (b *bishop) String() string {
	return fmt.Sprintf("Bishop{color: %v}", b.color)
}
