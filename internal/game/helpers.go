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

func (g *Game) parseFENPosition(fen string) error {
	parts := strings.Split(fen, " ")

	if len(parts) != 6 {
		return ErrInvalidFenCode
	}

	boardPart := parts[0]
	parsedBoard, err := parseFENBoard(boardPart)
	if err != nil {
		return err
	}

	g.board.squares = parsedBoard

	turnPart := parts[1]
	switch turnPart {
	case "w":
		g.colorTurn = white
	case "b":
		g.colorTurn = black
	default:
		return ErrInvalidFenCode
	}

	castlingPart := parts[2]
	whiteCastlingRights, blackCastlingRights, err := parseFENCastling(castlingPart)
	if err != nil {
		return err
	}

	g.whiteCastlingRights = whiteCastlingRights
	g.blackCastlingRights = blackCastlingRights

	enPassantPart := parts[3]

	_, ok := g.board.coords[enPassantPart]
	if ok {
		g.enPassantSquare = enPassantPart
	} else {
		g.enPassantSquare = ""
	}
	return nil
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
				row = append(row, square{pieceAt: newKing(black), isOccupied: true})
			case 'K':
				row = append(row, square{pieceAt: newKing(white), isOccupied: true})
			case 'q':
				row = append(row, square{pieceAt: newQueen(black), isOccupied: true})
			case 'Q':
				row = append(row, square{pieceAt: newQueen(white), isOccupied: true})
			case 'r':
				row = append(row, square{pieceAt: newRook(black), isOccupied: true})
			case 'R':
				row = append(row, square{pieceAt: newRook(white), isOccupied: true})
			case 'b':
				row = append(row, square{pieceAt: newBishop(black), isOccupied: true})
			case 'B':
				row = append(row, square{pieceAt: newBishop(white), isOccupied: true})
			case 'n':
				row = append(row, square{pieceAt: newKnight(black), isOccupied: true})
			case 'N':
				row = append(row, square{pieceAt: newKnight(white), isOccupied: true})
			case 'p':
				row = append(row, square{pieceAt: newPawn(black), isOccupied: true})
			case 'P':
				row = append(row, square{pieceAt: newPawn(white), isOccupied: true})
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

func parseFENCastling(castlingPart string) (castlingRights, castlingRights, error) {
	if castlingPart == "-" {
		return castlingRights{}, castlingRights{}, nil
	}

	whiteCastlingRights := castlingRights{}
	blackCastlingRights := castlingRights{}

	seen := make(map[rune]struct{})

	for _, ch := range castlingPart {
		if _, ok := seen[ch]; ok {
			return castlingRights{}, castlingRights{}, ErrInvalidFenCode
		} else {
			seen[ch] = struct{}{}
		}

		switch ch {
		case 'K':
			whiteCastlingRights.canKingSide = true
		case 'Q':
			whiteCastlingRights.canQueenSide = true
		case 'k':
			blackCastlingRights.canKingSide = true
		case 'q':
			blackCastlingRights.canQueenSide = true
		default:
			return castlingRights{}, castlingRights{}, ErrInvalidFenCode
		}
	}

	return whiteCastlingRights, blackCastlingRights, nil
}

func distance(a, b int) int {
	result := a - b
	if result < 0 {
		return result * -1
	}
	return result
}
