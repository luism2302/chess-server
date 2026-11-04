package game

import (
	"fmt"
)

const (
	PawnVal = 1
)

type pawn struct {
	color     color
	val       int
	pieceRune rune
}

func newPawn(color color) *pawn {
	var pieceRune rune
	switch color {
	case white:
		pieceRune = '\u2659'
	case black:
		pieceRune = '\u265F'
	}
	return &pawn{
		color:     color,
		val:       PawnVal,
		pieceRune: pieceRune,
	}
}

func (p *pawn) getAvailableMoves(pos coords, game *Game) []coords {
	verticalOffset := -1
	horizontalOffset := 1

	if p.color == black {
		verticalOffset = 1
	}

	// upper and lower boundries
	if pos.row+verticalOffset < 0 || pos.row+verticalOffset > 7 {
		return []coords{}
	}

	moves := []coords{}
	upByOneSquare := game.board.squares[pos.row+verticalOffset][pos.col]
	upByOneRow := pos.row + verticalOffset

	if !upByOneSquare.isOccupied {
		moves = append(moves, coords{row: upByOneRow, col: pos.col})
	}

	// if white and in rank 6 (in the grid, not the board notation) , this the pawn's first move, if black, check rank 1
	if p.color == white {
		if pos.row == 6 {
			upByTwoSquare := game.board.squares[pos.row+2*verticalOffset][pos.col]
			upByTwoRow := pos.row + 2*verticalOffset
			if !upByTwoSquare.isOccupied && !upByOneSquare.isOccupied {
				moves = append(moves, coords{row: upByTwoRow, col: pos.col})
			}
		}
	} else {
		if pos.row == 1 {
			upByTwoSquare := game.board.squares[pos.row+2*verticalOffset][pos.col]
			upByTwoRow := pos.row + 2*verticalOffset
			if !upByTwoSquare.isOccupied && !upByOneSquare.isOccupied {
				moves = append(moves, coords{row: upByTwoRow, col: pos.col})
			}

		}
	}

	enPassantCoords := game.board.coords[game.enPassantSquare]
	if pos.col+horizontalOffset <= 7 {
		upRightSquare := game.board.squares[pos.row+verticalOffset][pos.col+horizontalOffset]
		upRightCoords := coords{row: pos.row + verticalOffset, col: pos.col + horizontalOffset}
		if upRightSquare.isOccupied && upRightSquare.pieceAt.getPieceColor() != p.getPieceColor() {
			moves = append(moves, coords{row: pos.row + verticalOffset, col: pos.col + horizontalOffset})
		} else if game.enPassantSquare != "" && enPassantCoords == upRightCoords {
			moves = append(moves, coords{row: pos.row + verticalOffset, col: pos.col + horizontalOffset})
		}
	}

	if pos.col-horizontalOffset >= 0 {
		upLeftSquare := game.board.squares[pos.row+verticalOffset][pos.col-horizontalOffset]
		upLeftCoords := coords{row: pos.row + verticalOffset, col: pos.col - horizontalOffset}
		if upLeftSquare.isOccupied && upLeftSquare.pieceAt.getPieceColor() != p.getPieceColor() {
			moves = append(moves, coords{row: pos.row + verticalOffset, col: pos.col - horizontalOffset})
		} else if game.enPassantSquare != "" && enPassantCoords == upLeftCoords {
			moves = append(moves, coords{row: pos.row + verticalOffset, col: pos.col - horizontalOffset})
		}
	}

	// TODO: promote when row + 1 == 0 or 7 depending on color

	return moves
}

func (p *pawn) getPieceVal() int {
	return p.val
}

func (p *pawn) getPieceColor() color {
	return p.color
}

func (p *pawn) getPieceRune() rune {
	return p.pieceRune
}

func (p *pawn) String() string {
	return fmt.Sprintf("Pawn{color: %v}", p.color)
}
