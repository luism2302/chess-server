package game

import (
	"errors"
	"slices"
	"strings"
	"unicode"
)

var (
	ErrInvalidFenPiece = errors.New("invalid fen piece")
	ErrInvalidFenCode  = errors.New("invalid fen code")
)

func parseFEN(fen string) board {
	return board{}
}

func parseFENBoard(fenBoard string) ([][]square, error) {
	ranksText := strings.Split(fenBoard, "/")
	if len(ranksText) != 8 {
		return nil, ErrInvalidFenCode
	}

	boardSquares := [][]square{}

	for _, r := range ranksText {
		rank, err := parseFENRank(r)
		if err != nil {
			return nil, err
		}
		boardSquares = append(boardSquares, rank)
	}
	return boardSquares, nil
}

func parseFENRank(rank string) ([]square, error) {
	row := []square{}
	for _, ch := range rank {
		switch {
		case unicode.IsDigit(ch):
			val := int(ch - '0')
			emptySquares := slices.Repeat([]square{{pieceAt: newEmpty(), isOccupied: false}}, val)
			row = append(row, emptySquares...)
		case unicode.IsLetter(ch):
			switch ch {
			case 'k':
				row = append(row, square{pieceAt: newKing(false), isOccupied: true})
			case 'K':
				row = append(row, square{pieceAt: newKing(true), isOccupied: true})
			case 'q':
				row = append(row, square{pieceAt: newQueen(false), isOccupied: true})
			case 'Q':
				row = append(row, square{pieceAt: newQueen(true), isOccupied: true})
			case 'r':
				row = append(row, square{pieceAt: newRook(false), isOccupied: true})
			case 'R':
				row = append(row, square{pieceAt: newRook(true), isOccupied: true})
			case 'b':
				row = append(row, square{pieceAt: newBishop(false), isOccupied: true})
			case 'B':
				row = append(row, square{pieceAt: newBishop(true), isOccupied: true})
			case 'n':
				row = append(row, square{pieceAt: newKnight(false), isOccupied: true})
			case 'N':
				row = append(row, square{pieceAt: newKnight(true), isOccupied: true})
			case 'p':
				row = append(row, square{pieceAt: newPawn(false), isOccupied: true})
			case 'P':
				row = append(row, square{pieceAt: newPawn(true), isOccupied: true})
			default:
				return []square{}, ErrInvalidFenPiece
			}
		default:
			return []square{}, ErrInvalidFenPiece
		}
	}

	if len(row) != 8 {
		return []square{}, ErrInvalidFenCode
	}

	return row, nil
}
