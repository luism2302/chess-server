package game

import (
	"errors"
)

var (
	ErrUnknownPieceCode = errors.New("unknown Piece")
	ErrInvalidMove      = errors.New("invalid Move")
)

type piece interface {
	getAvailableMoves(pos coords, board board, game *Game) []coords
	getPieceColor() bool
	getPieceVal() int
	getPieceRune() rune
	String() string
}
