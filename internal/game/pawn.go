package game

import (
	"fmt"
	"log"
)

const (
	PawnVal = 1
)

type pawn struct {
	isWhite       bool
	val           int
	isFirstMove   bool
	enPassantable bool
	pieceRune     rune
}

func newPawn(color bool) *pawn {
	pieceRune := '\u265F'
	if color {
		pieceRune = '\u2659'
	}
	return &pawn{
		isWhite:       color,
		val:           PawnVal,
		isFirstMove:   true,
		enPassantable: false,
		pieceRune:     pieceRune,
	}
}

func (p *pawn) getAvailableMoves(pos coords, board board) []coords {
	verticalOffset := -1
	horizontalOffset := 1

	if !p.isWhite {
		verticalOffset = 1
	}

	// upper and lower boundries
	if pos.row+verticalOffset < 0 || pos.row+verticalOffset > 7 {
		return []coords{}
	}

	moves := []coords{}

	if !board.squares[pos.row+verticalOffset][pos.col].isOccupied {
		moves = append(moves, coords{row: pos.row + verticalOffset, col: pos.col})
	}

	if p.isFirstMove {
		if !board.squares[pos.row+2*verticalOffset][pos.col].isOccupied {
			moves = append(moves, coords{row: pos.row + 2*verticalOffset, col: pos.col})
		}
	}

	if pos.col+horizontalOffset <= 7 {

		log.Print("checking take at edge h")
		upRightSquare := board.squares[pos.row+verticalOffset][pos.col+horizontalOffset]
		if upRightSquare.isOccupied {
			moves = append(moves, coords{row: pos.row + verticalOffset, col: pos.col + horizontalOffset})
		}
	}

	if pos.col-horizontalOffset >= 0 {
		log.Print("checking take at edge a")
		upLeftSquare := board.squares[pos.row+verticalOffset][pos.col-horizontalOffset]
		if upLeftSquare.isOccupied {
			moves = append(moves, coords{row: pos.row + verticalOffset, col: pos.col - horizontalOffset})
		}
	}

	// TODO: promote when row + 1 == 0 or 7 depending on color

	return moves
}

func (p *pawn) getPieceVal() int {
	return p.val
}

func (p *pawn) getPieceColor() bool {
	return p.isWhite
}

func (p *pawn) getPieceRune() rune {
	return p.pieceRune
}

func (p *pawn) updateFirstMove() {
	p.isFirstMove = false
}

func (p *pawn) String() string {
	return fmt.Sprintf("Pawn{color: %t, isFirstMove: %t, isEnpassantable: %t}", p.isWhite, p.isFirstMove, p.enPassantable)
}
