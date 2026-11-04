package game

import (
	"errors"
)

var (
	ErrUnknownPieceCode = errors.New("unknown Piece")
	ErrInvalidMove      = errors.New("invalid Move")
)

type color int

const (
	black color = iota
	white
	noColor
)

type piece interface {
	getAvailableMoves(pos coords, game *Game) []coords
	getPieceColor() color
	getPieceVal() int
	getPieceRune() rune
	String() string
}
