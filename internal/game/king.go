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
	upOffset := -1
	downOffset := 1
	leftOffset := -1
	rightOffset := 1

	moves := []coords{}

	// up
	if pos.row+upOffset >= 0 {
		targetSquare := game.board.squares[pos.row+upOffset][pos.col]
		if targetSquare.isOccupied {
			if targetSquare.pieceAt.getPieceColor() != k.getPieceColor() {
				moves = append(moves, coords{row: pos.row + upOffset, col: pos.col})
			}
		} else {
			moves = append(moves, coords{row: pos.row + upOffset, col: pos.col})
		}
	}

	// down
	if pos.row+downOffset <= 7 {
		targetSquare := game.board.squares[pos.row+downOffset][pos.col]
		if targetSquare.isOccupied {
			if targetSquare.pieceAt.getPieceColor() != k.getPieceColor() {
				moves = append(moves, coords{row: pos.row + downOffset, col: pos.col})
			}
		} else {
			moves = append(moves, coords{row: pos.row + downOffset, col: pos.col})
		}
	}

	// left
	if pos.col+leftOffset >= 0 {
		targetSquare := game.board.squares[pos.row][pos.col+leftOffset]
		if targetSquare.isOccupied {
			if targetSquare.pieceAt.getPieceColor() != k.getPieceColor() {
				moves = append(moves, coords{row: pos.row, col: pos.col + leftOffset})
			}
		} else {
			moves = append(moves, coords{row: pos.row, col: pos.col + leftOffset})
		}
	}

	// right
	if pos.col+rightOffset <= 7 {
		targetSquare := game.board.squares[pos.row][pos.col+rightOffset]
		if targetSquare.isOccupied {
			if targetSquare.pieceAt.getPieceColor() != k.getPieceColor() {
				moves = append(moves, coords{row: pos.row, col: pos.col + rightOffset})
			}
		} else {
			moves = append(moves, coords{row: pos.row, col: pos.col + rightOffset})
		}
	}

	// upRight diag
	if pos.col+rightOffset <= 7 && pos.row+upOffset >= 0 {
		targetSquare := game.board.squares[pos.row+upOffset][pos.col+rightOffset]
		if targetSquare.isOccupied {
			if targetSquare.pieceAt.getPieceColor() != k.getPieceColor() {
				moves = append(moves, coords{row: pos.row + upOffset, col: pos.col + rightOffset})
			}
		} else {
			moves = append(moves, coords{row: pos.row + upOffset, col: pos.col + rightOffset})
		}
	}

	// downRight diag
	if pos.col+rightOffset <= 7 && pos.row+downOffset <= 7 {
		targetSquare := game.board.squares[pos.row+downOffset][pos.col+rightOffset]
		if targetSquare.isOccupied {
			if targetSquare.pieceAt.getPieceColor() != k.getPieceColor() {
				moves = append(moves, coords{row: pos.row + downOffset, col: pos.col + rightOffset})
			}
		} else {
			moves = append(moves, coords{row: pos.row + downOffset, col: pos.col + rightOffset})
		}
	}

	// upLeft diag
	if pos.col+leftOffset >= 0 && pos.row+upOffset >= 0 {
		targetSquare := game.board.squares[pos.row+upOffset][pos.col+leftOffset]
		if targetSquare.isOccupied {
			if targetSquare.pieceAt.getPieceColor() != k.getPieceColor() {
				moves = append(moves, coords{row: pos.row + upOffset, col: pos.col + leftOffset})
			}
		} else {
			moves = append(moves, coords{row: pos.row + upOffset, col: pos.col + leftOffset})
		}
	}

	// downLeft diag
	if pos.col+leftOffset >= 0 && pos.row+downOffset <= 7 {
		targetSquare := game.board.squares[pos.row+downOffset][pos.col+leftOffset]
		if targetSquare.isOccupied {
			if targetSquare.pieceAt.getPieceColor() != k.getPieceColor() {
				moves = append(moves, coords{row: pos.row + downOffset, col: pos.col + leftOffset})
			}
		} else {
			moves = append(moves, coords{row: pos.row + downOffset, col: pos.col + leftOffset})
		}
	}

	// castling
	switch k.getPieceColor() {
	case white:
		if game.whiteCastlingRights.canQueenSide {
			if pos.col+2*leftOffset >= 0 && !game.board.squares[pos.row][pos.col+leftOffset].isOccupied && !game.board.squares[pos.row][pos.col+2*leftOffset].isOccupied && !game.board.squares[pos.row][pos.col+3*leftOffset].isOccupied {
				moves = append(moves, coords{row: 7, col: 2})
			}
		}
		if game.whiteCastlingRights.canKingSide {
			if pos.col+2*rightOffset >= 0 && !game.board.squares[pos.row][pos.col+rightOffset].isOccupied && !game.board.squares[pos.row][pos.col+2*rightOffset].isOccupied {
				moves = append(moves, coords{row: 7, col: 6})
			}
		}
	case black:
		if game.blackCastlingRights.canQueenSide {
			if pos.col+2*leftOffset >= 0 && !game.board.squares[pos.row][pos.col+leftOffset].isOccupied && !game.board.squares[pos.row][pos.col+2*leftOffset].isOccupied && !game.board.squares[pos.row][pos.col+3*leftOffset].isOccupied {
				moves = append(moves, coords{row: 0, col: 2})
			}
		}
		if game.blackCastlingRights.canKingSide {
			if pos.col+2*rightOffset >= 0 && !game.board.squares[pos.row][pos.col+rightOffset].isOccupied && !game.board.squares[pos.row][pos.col+2*rightOffset].isOccupied {
				moves = append(moves, coords{row: 0, col: 6})
			}
		}

	}

	return moves
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
