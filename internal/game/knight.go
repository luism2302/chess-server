package game

import "fmt"

const (
	KnightVal = 2
)

type knight struct {
	color     color
	val       int
	pieceRune rune
}

func newKnight(color color) *knight {
	var pieceRune rune
	switch color {
	case white:
		pieceRune = '\u2658'
	case black:
		pieceRune = '\u265E'
	}

	return &knight{
		color:     color,
		val:       KnightVal,
		pieceRune: pieceRune,
	}
}

func (k *knight) getAvailableMoves(pos coords, game *Game) []coords {
	// two up one left
	moves := []coords{}
	if pos.row-2 >= 0 && pos.col-1 >= 0 {
		targetSquare := game.board.squares[pos.row-2][pos.col-1]
		if !targetSquare.isOccupied || targetSquare.pieceAt.getPieceColor() != k.getPieceColor() {
			moves = append(moves, coords{row: pos.row - 2, col: pos.col - 1})
		}
	}
	// two up one right
	if pos.row-2 >= 0 && pos.col+1 <= 7 {
		targetSquare := game.board.squares[pos.row-2][pos.col+1]
		if !targetSquare.isOccupied || targetSquare.pieceAt.getPieceColor() != k.getPieceColor() {
			moves = append(moves, coords{row: pos.row - 2, col: pos.col + 1})
		}
	}
	// two down one left
	if pos.row+2 <= 7 && pos.col-1 >= 0 {
		targetSquare := game.board.squares[pos.row+2][pos.col-1]
		if !targetSquare.isOccupied || targetSquare.pieceAt.getPieceColor() != k.getPieceColor() {
			moves = append(moves, coords{row: pos.row + 2, col: pos.col - 1})
		}
	}
	// two down one right
	if pos.row+2 <= 7 && pos.col+1 <= 7 {
		targetSquare := game.board.squares[pos.row+2][pos.col+1]
		if !targetSquare.isOccupied || targetSquare.pieceAt.getPieceColor() != k.getPieceColor() {
			moves = append(moves, coords{row: pos.row + 2, col: pos.col + 1})
		}
	}
	// one up two left
	if pos.row-1 >= 0 && pos.col-2 >= 0 {
		targetSquare := game.board.squares[pos.row-1][pos.col-2]
		if !targetSquare.isOccupied || targetSquare.pieceAt.getPieceColor() != k.getPieceColor() {
			moves = append(moves, coords{row: pos.row - 1, col: pos.col - 2})
		}
	}
	// one up two right
	if pos.row-1 >= 0 && pos.col+2 <= 7 {
		targetSquare := game.board.squares[pos.row-1][pos.col+2]
		if !targetSquare.isOccupied || targetSquare.pieceAt.getPieceColor() != k.getPieceColor() {
			moves = append(moves, coords{row: pos.row - 1, col: pos.col + 2})
		}
	}
	// one down two left
	if pos.row+1 <= 7 && pos.col-2 >= 0 {
		targetSquare := game.board.squares[pos.row+1][pos.col-2]
		if !targetSquare.isOccupied || targetSquare.pieceAt.getPieceColor() != k.getPieceColor() {
			moves = append(moves, coords{row: pos.row + 1, col: pos.col - 2})
		}
	}
	// one down two right
	if pos.row+1 <= 7 && pos.col+2 <= 7 {
		targetSquare := game.board.squares[pos.row+1][pos.col+2]
		if !targetSquare.isOccupied || targetSquare.pieceAt.getPieceColor() != k.getPieceColor() {
			moves = append(moves, coords{row: pos.row + 1, col: pos.col + 2})
		}
	}

	return moves
}

func (k *knight) getPieceVal() int {
	return k.val
}

func (k *knight) getPieceColor() color {
	return k.color
}

func (k *knight) getPieceRune() rune {
	return k.pieceRune
}

func (k *knight) String() string {
	return fmt.Sprintf("Knight{color: %v}", k.color)
}
