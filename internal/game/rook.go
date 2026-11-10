package game

import "fmt"

const (
	RookVal = 5
)

type rook struct {
	color     color
	val       int
	pieceRune rune
}

func newRook(color color) *rook {
	var pieceRune rune
	switch color {
	case white:
		pieceRune = '\u2656'
	case black:
		pieceRune = '\u265C'
	}
	return &rook{
		color:     color,
		val:       RookVal,
		pieceRune: pieceRune,
	}
}

func (r *rook) getAvailableMoves(pos coords, game *Game) []coords {
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
			if targetSquare.pieceAt.getPieceColor() != r.getPieceColor() {
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
			if targetSquare.pieceAt.getPieceColor() != r.getPieceColor() {
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
			if targetSquare.pieceAt.getPieceColor() != r.getPieceColor() {
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
			if targetSquare.pieceAt.getPieceColor() != r.getPieceColor() {
				moves = append(moves, coords{row: pos.row, col: targetCol})
			}
			break
		}
		moves = append(moves, coords{row: pos.row, col: targetCol})
	}

	return moves
}

func (r *rook) getPieceVal() int {
	return r.val
}

func (r *rook) getPieceColor() color {
	return r.color
}

func (r *rook) getPieceRune() rune {
	return r.pieceRune
}

func (r *rook) String() string {
	return fmt.Sprintf("Rook{color: %v}", r.color)
}
