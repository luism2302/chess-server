package game

import "fmt"

const (
	QueenVal = 9
)

type queen struct {
	color     color
	val       int
	pieceRune rune
}

func newQueen(color color) *queen {
	var pieceRune rune
	switch color {
	case white:
		pieceRune = '\u2655'
	case black:
		pieceRune = '\u265B'
	}
	return &queen{
		color:     color,
		val:       QueenVal,
		pieceRune: pieceRune,
	}
}

func (q *queen) getAvailableMoves(pos coords, game *Game) []coords {
	lenBoard := 8
	upOffset := -1
	downOffset := 1
	leftOffset := -1
	rightOffset := 1

	moves := []coords{}

	// up
	for i := 1; i < lenBoard; i++ {
		targetRow := pos.row + i*upOffset

		if targetRow < 0 {
			break
		}

		targetSquare := game.board.squares[targetRow][pos.col]
		if targetSquare.isOccupied {
			if targetSquare.pieceAt.getPieceColor() != q.getPieceColor() {
				moves = append(moves, coords{row: targetRow, col: pos.col})
			}
			break
		}
		moves = append(moves, coords{row: targetRow, col: pos.col})
	}

	// down
	for i := 1; i < lenBoard; i++ {
		targetRow := pos.row + i*downOffset

		if targetRow > 7 {
			break
		}

		targetSquare := game.board.squares[targetRow][pos.col]
		if targetSquare.isOccupied {
			if targetSquare.pieceAt.getPieceColor() != q.getPieceColor() {
				moves = append(moves, coords{row: targetRow, col: pos.col})
			}
			break
		}
		moves = append(moves, coords{row: targetRow, col: pos.col})
	}

	// left
	for i := 1; i < lenBoard; i++ {
		targetCol := pos.col + i*leftOffset

		if targetCol < 0 {
			break
		}
		targetSquare := game.board.squares[pos.row][targetCol]
		if targetSquare.isOccupied {
			if targetSquare.pieceAt.getPieceColor() != q.getPieceColor() {
				moves = append(moves, coords{row: pos.row, col: targetCol})
			}
			break
		}
		moves = append(moves, coords{row: pos.row, col: targetCol})
	}

	// right
	for i := 1; i < lenBoard; i++ {
		targetCol := pos.col + i*rightOffset

		if targetCol > 7 {
			break
		}
		targetSquare := game.board.squares[pos.row][targetCol]
		if targetSquare.isOccupied {
			if targetSquare.pieceAt.getPieceColor() != q.getPieceColor() {
				moves = append(moves, coords{row: pos.row, col: targetCol})
			}
			break
		}
		moves = append(moves, coords{row: pos.row, col: targetCol})
	}

	for i := 1; i < lenBoard; i++ {
		targetRow := pos.row + i*upOffset
		targetCol := pos.col + i*rightOffset

		if targetRow < 0 || targetCol > 7 {
			break
		}
		targetSquare := game.board.squares[targetRow][targetCol]
		if targetSquare.isOccupied {
			if targetSquare.pieceAt.getPieceColor() != q.getPieceColor() {
				moves = append(moves, coords{row: targetRow, col: targetCol})
			}
			break
		}
		moves = append(moves, coords{row: targetRow, col: targetCol})
	}

	// downRight diag
	for i := 1; i < lenBoard; i++ {
		targetRow := pos.row + i*downOffset
		targetCol := pos.col + i*rightOffset

		if targetRow > 7 || targetCol > 7 {
			break
		}
		targetSquare := game.board.squares[targetRow][targetCol]
		if targetSquare.isOccupied {
			if targetSquare.pieceAt.getPieceColor() != q.getPieceColor() {
				moves = append(moves, coords{row: targetRow, col: targetCol})
			}
			break
		}
		moves = append(moves, coords{row: targetRow, col: targetCol})
	}

	// upLeft diag
	for i := 1; i < lenBoard; i++ {
		targetRow := pos.row + i*upOffset
		targetCol := pos.col + i*leftOffset

		if targetRow < 0 || targetCol < 0 {
			break
		}
		targetSquare := game.board.squares[targetRow][targetCol]
		if targetSquare.isOccupied {
			if targetSquare.pieceAt.getPieceColor() != q.getPieceColor() {
				moves = append(moves, coords{row: targetRow, col: targetCol})
			}
			break
		}
		moves = append(moves, coords{row: targetRow, col: targetCol})
	}

	// downLeft diag
	for i := 1; i < lenBoard; i++ {
		targetRow := pos.row + i*downOffset
		targetCol := pos.col + i*leftOffset

		if targetRow > 7 || targetCol < 0 {
			break
		}
		targetSquare := game.board.squares[targetRow][targetCol]
		if targetSquare.isOccupied {
			if targetSquare.pieceAt.getPieceColor() != q.getPieceColor() {
				moves = append(moves, coords{row: targetRow, col: targetCol})
			}
			break
		}
		moves = append(moves, coords{row: targetRow, col: targetCol})
	}

	return moves
}

func (q *queen) getPieceVal() int {
	return q.val
}

func (q *queen) getPieceColor() color {
	return q.color
}

func (q *queen) getPieceRune() rune {
	return q.pieceRune
}

func (q *queen) String() string {
	return fmt.Sprintf("Queen{color: %v}", q.color)
}
