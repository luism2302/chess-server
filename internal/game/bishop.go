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
	lenBoard := 8
	upOffset := -1
	downOffset := 1
	leftOffset := -1
	rightOffset := 1

	moves := []coords{}

	// rightUp diag
	for i := 1; i < lenBoard; i++ {
		targetRow := pos.row + i*upOffset
		targetCol := pos.col + i*rightOffset

		if targetRow < 0 || targetCol > 7 {
			break
		}
		targetSquare := game.board.squares[targetRow][targetCol]
		if targetSquare.isOccupied {
			if targetSquare.pieceAt.getPieceColor() != b.getPieceColor() {
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
			if targetSquare.pieceAt.getPieceColor() != b.getPieceColor() {
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
			if targetSquare.pieceAt.getPieceColor() != b.getPieceColor() {
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
			if targetSquare.pieceAt.getPieceColor() != b.getPieceColor() {
				moves = append(moves, coords{row: targetRow, col: targetCol})
			}
			break
		}
		moves = append(moves, coords{row: targetRow, col: targetCol})
	}

	return moves
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
