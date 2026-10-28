package game

import (
	"errors"
)

const (
	PawnVal   = 1
	KnightVal = 2
	BishopVal = 3
	RookVal   = 5
	QueenVal  = 9
	KingVal   = 20
	EmptyVal  = 0
)

var (
	WhitePawn   = piece{val: PawnVal, isWhite: true, isFirstMove: true}
	BlackPawn   = piece{val: -PawnVal, isWhite: false, isFirstMove: true}
	WhiteKnight = piece{val: KnightVal, isWhite: true, isFirstMove: true}
	BlackKnight = piece{val: -KnightVal, isWhite: false, isFirstMove: true}
	WhiteBishop = piece{val: BishopVal, isWhite: true, isFirstMove: true}
	BlackBishop = piece{val: -BishopVal, isWhite: false, isFirstMove: true}
	WhiteRook   = piece{val: RookVal, isWhite: true, isFirstMove: true}
	BlackRook   = piece{val: -RookVal, isWhite: false, isFirstMove: true}
	WhiteQueen  = piece{val: QueenVal, isWhite: true, isFirstMove: true}
	BlackQueen  = piece{val: -QueenVal, isWhite: false, isFirstMove: true}
	WhiteKing   = piece{val: KingVal, isWhite: true, isFirstMove: true}
	BlackKing   = piece{val: -KingVal, isWhite: false, isFirstMove: true}
	Empty       = piece{val: EmptyVal, isWhite: false, isFirstMove: true}
)

var (
	ErrUnknownPieceCode = errors.New("Unknown Piece")
	ErrInvalidMove      = errors.New("Invalid Move")
)

type piece struct {
	val         int
	isWhite     bool
	isFirstMove bool
}

var piecesMap = map[int]rune{
	WhiteKing.val:   '\u2654', // White King
	WhiteQueen.val:  '\u2655', // White Queen
	WhiteRook.val:   '\u2656', // White Rook
	WhiteBishop.val: '\u2657', // White Bishop
	WhiteKnight.val: '\u2658', // White Knight
	WhitePawn.val:   '\u2659', // White Pawn
	BlackKing.val:   '\u265A', // Black King
	BlackQueen.val:  '\u265B', // Black Queen
	BlackRook.val:   '\u265C', // Black Rook
	BlackBishop.val: '\u265D', // Black Bishop
	BlackKnight.val: '\u265E', // Black Knight
	BlackPawn.val:   '\u265F', // Black Pawn
	Empty.val:       '\u25AA',
}

func pieceRune(pieceVal int) (rune, error) {
	repr, ok := piecesMap[pieceVal]
	if !ok {
		return '\x00', ErrUnknownPieceCode
	}
	return repr, nil
}

func getAvailableMoves(piece piece, pos coords, board board) []coords {
	var moves []coords
	switch piece {
	case WhitePawn, BlackPawn:
		moves = getPawnMoves(piece, pos, board)
	}
	return moves
}

func getPawnMoves(pawn piece, pos coords, board board) []coords {
	verticalOffset := 1
	horizontalOffset := 1

	if !pawn.isWhite {
		verticalOffset = -1
	}

	// TODO: promote when row + 1 == 0 or 7 depending on color

	moves := []coords{
		{row: pos.row + verticalOffset, col: pos.col},
	}

	if pawn.isFirstMove {
		moves = append(moves, coords{row: pos.row + 2*verticalOffset, col: pos.col})
	}

	if pos.col+horizontalOffset <= 7 {
		upRightSquare := board.squares[pos.row+verticalOffset][pos.col+horizontalOffset]
		if upRightSquare.isOccupied {
			moves = append(moves, coords{row: pos.row + verticalOffset, col: pos.col + horizontalOffset})
		}
	}

	if pos.col-horizontalOffset >= 0 {
		upLeftSquare := board.squares[pos.row+verticalOffset][pos.col-horizontalOffset]
		if upLeftSquare.isOccupied {
			moves = append(moves, coords{row: pos.row + verticalOffset, col: pos.col - horizontalOffset})
		}
	}

	return moves
}
