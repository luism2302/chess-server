package game

import (
	"errors"
)

var (
	ErrUnknownPieceCode = errors.New("Unknown Piece")
	ErrInvalidMove      = errors.New("Invalid Move")
)

type piece interface {
	getAvailableMoves(pos coords, board board) []coords
	getPieceColor() bool
	getPieceVal() int
	getPieceRune() rune
	String() string
}
