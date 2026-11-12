package game

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

const (
	initialFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)

var (
	ErrInvalidTurn   = errors.New("other player's turn")
	ErrSelectedEmpty = errors.New("selected an empty square")
)

type castlingRights struct {
	canKingSide  bool
	canQueenSide bool
}
type Game struct {
	board               board
	blackCastlingRights castlingRights
	whiteCastlingRights castlingRights
	enPassantSquare     string
	colorTurn           color
	isWhiteInCheck      bool
	isBlackInCheck      bool
}

func New(fen string) (*Game, error) {
	game := &Game{
		board: board{
			coords: initializeCoordsMap(),
		},
	}

	if fen == "" {
		fen = initialFen
	}

	if err := game.parseFENPosition(fen); err != nil {
		return nil, err
	}
	return game, nil
}

func (g *Game) BoardRepr() string {
	var builder strings.Builder
	builder.WriteString(border)
	for i, row := range g.board.squares {
		builder.WriteString(fmt.Sprintf("%d | ", 8-i))
		for _, square := range row {
			builder.WriteString(" ")
			repr := square.pieceAt.getPieceRune()
			builder.WriteRune(repr)
			builder.WriteString(" ")
		}
		builder.WriteString(" |")
		builder.WriteString("\n\n")
	}
	builder.WriteString(border)
	builder.WriteString(cols)
	builder.WriteString("\n\n")
	return builder.String()
}

func (g *Game) Update(move string) error {
	startingCoords, endingCoords, err := g.validateMove(move)
	if err != nil {
		return err
	}

	pieceAtStart := g.getPieceAtCoords(startingCoords)
	pieceAtEnd := g.getPieceAtCoords(endingCoords)

	if pieceAtStart.getPieceVal() == EmptyVal {
		return ErrSelectedEmpty
	}

	if pieceAtStart.getPieceColor() != g.colorTurn {
		return ErrInvalidTurn
	}

	if pieceAtEnd.getPieceColor() == pieceAtStart.getPieceColor() && pieceAtEnd.getPieceVal() != EmptyVal {
		return ErrInvalidMove
	}

	availableMoves := pieceAtStart.getAvailableMoves(startingCoords, g)

	if slices.Contains(availableMoves, endingCoords) {
		updatedSquares, err := g.updateBoard(pieceAtStart, startingCoords, endingCoords)
		if err != nil {
			return err
		}

		g.board.squares = updatedSquares
		if g.colorTurn == white {
			g.colorTurn = black
		} else {
			g.colorTurn = white
		}

		a1RookCoords := coords{row: 7, col: 0}
		h1RookCoords := coords{row: 7, col: 7}
		a8RookCoords := coords{row: 0, col: 0}
		h8RookCoords := coords{row: 0, col: 7}

		switch v := pieceAtStart.(type) {
		case *pawn:
			if distance(startingCoords.row, endingCoords.row) == 2 {
				verticalOffset := -1
				if v.getPieceColor() == white {
					verticalOffset = 1
				}
				enPassantSquareCoords := coords{row: endingCoords.row + verticalOffset, col: endingCoords.col}
				g.coordsToSquareName(enPassantSquareCoords)
				return nil
			}
		case *rook:
			if v.getPieceColor() == white {
				if startingCoords == a1RookCoords {
					g.whiteCastlingRights.canQueenSide = false
				} else if startingCoords == h1RookCoords {
					g.whiteCastlingRights.canKingSide = false
				}
			}

			if v.getPieceColor() == black {
				if startingCoords == a8RookCoords {
					g.blackCastlingRights.canQueenSide = false
				} else if startingCoords == h8RookCoords {
					g.blackCastlingRights.canKingSide = false
				}
			}
		case *king:
			if v.getPieceColor() == white {
				g.whiteCastlingRights.canKingSide = false
				g.whiteCastlingRights.canQueenSide = false
			} else if v.getPieceColor() == black {
				g.blackCastlingRights.canKingSide = false
				g.blackCastlingRights.canQueenSide = false
			}
		}

		g.enPassantSquare = ""

		switch v := pieceAtEnd.(type) {
		case *rook:
			if endingCoords == a1RookCoords && v.getPieceColor() == white {
				g.whiteCastlingRights.canQueenSide = false
			} else if endingCoords == h1RookCoords && v.getPieceColor() == white {
				g.whiteCastlingRights.canKingSide = false
			} else if endingCoords == a8RookCoords && v.getPieceColor() == black {
				g.blackCastlingRights.canQueenSide = false
			} else if endingCoords == h8RookCoords && v.getPieceColor() == black {
				g.blackCastlingRights.canKingSide = false
			}
		}

	} else {
		return ErrInvalidMove
	}

	return nil
}

func (g *Game) validateMove(move string) (coords, coords, error) {
	trimmed := strings.ReplaceAll(move, " ", "")

	if len(trimmed) != 5 {
		return coords{}, coords{}, ErrInvalidMove
	}

	parts := strings.Split(trimmed, "-")

	if len(parts) != 2 {
		return coords{}, coords{}, ErrInvalidMove
	}

	startingPos, endingPos := parts[0], parts[1]
	if len(startingPos) != 2 || len(endingPos) != 2 {
		return coords{}, coords{}, ErrInvalidMove
	}

	startingPosCoords, ok := g.board.coords[startingPos]
	if !ok {
		return coords{}, coords{}, ErrInvalidMove
	}

	endingPosCoords, ok := g.board.coords[endingPos]
	if !ok {
		return coords{}, coords{}, ErrInvalidMove
	}

	return startingPosCoords, endingPosCoords, nil
}

func (g *Game) getPieceAtCoords(startCoords coords) piece {
	piece := g.board.squares[startCoords.row][startCoords.col].pieceAt
	return piece
}

func (g *Game) updateBoard(pieceMoved piece, startAt coords, endAt coords) ([][]square, error) {
	originalSquares := make([][]square, 0, len(g.board.squares))
	for _, r := range g.board.squares {
		originalSquares = append(originalSquares, slices.Clone(r))
	}

	updatedStartSquare := square{pieceAt: newEmpty(), isOccupied: false}
	g.board.squares[startAt.row][startAt.col] = updatedStartSquare

	updatedEndSquare := square{pieceAt: pieceMoved, isOccupied: true}
	g.board.squares[endAt.row][endAt.col] = updatedEndSquare

	switch v := pieceMoved.(type) {
	case *pawn:
		if g.enPassantSquare != "" && endAt == g.board.coords[g.enPassantSquare] {
			updatedEnPassantSquare := square{pieceAt: newEmpty(), isOccupied: false}
			updateCoords := g.updatePassantCoords(v, endAt)
			g.board.squares[updateCoords.row][updateCoords.col] = updatedEnPassantSquare

		}
	case *king:
		if distance(startAt.col, endAt.col) == 2 {
			err := g.validateCastling(endAt, v)
			if err != nil {
				g.board.squares = originalSquares
				return originalSquares, ErrInvalidMove
			}
			// starting rook coords
			a1RookCoords := coords{row: 7, col: 0}
			h1RookCoords := coords{row: 7, col: 7}
			a8RookCoords := coords{row: 0, col: 0}
			h8RookCoords := coords{row: 0, col: 7}

			// ending king coords
			c1Coords := coords{row: 7, col: 2}
			g1Coords := coords{row: 7, col: 6}
			c8Coords := coords{row: 0, col: 2}
			g8Coords := coords{row: 0, col: 6}

			// ending rook coords
			d1Coords := coords{row: 7, col: 3}
			f1Coords := coords{row: 7, col: 5}
			d8Coords := coords{row: 0, col: 3}
			f8Coords := coords{row: 0, col: 5}

			// starting king coords
			e1Coords := coords{row: 7, col: 4}
			e8Coords := coords{row: 0, col: 4}

			switch v.getPieceColor() {
			case white:
				if endAt == c1Coords {
					updatedEmptyRookSquare := square{pieceAt: newEmpty(), isOccupied: false}
					updatedRookSquare := square{pieceAt: newRook(white), isOccupied: true}
					updatedKingSquare := square{pieceAt: newKing(white), isOccupied: true}
					g.board.squares[a1RookCoords.row][a1RookCoords.col] = updatedEmptyRookSquare
					g.board.squares[c1Coords.row][c1Coords.col] = updatedKingSquare
					g.board.squares[d1Coords.row][d1Coords.col] = updatedRookSquare
				} else if endAt == g1Coords {
					updatedEmptyRookSquare := square{pieceAt: newEmpty(), isOccupied: false}
					updatedRookSquare := square{pieceAt: newRook(white), isOccupied: true}
					updatedKingSquare := square{pieceAt: newKing(white), isOccupied: true}
					g.board.squares[h1RookCoords.row][h1RookCoords.col] = updatedEmptyRookSquare
					g.board.squares[g1Coords.row][g1Coords.col] = updatedKingSquare
					g.board.squares[f1Coords.row][f1Coords.col] = updatedRookSquare
				}
				updatedEmptyKingSquare := square{pieceAt: newEmpty(), isOccupied: false}
				g.board.squares[e1Coords.row][e1Coords.col] = updatedEmptyKingSquare
			case black:
				if endAt == c8Coords {
					updatedEmptyRookSquare := square{pieceAt: newEmpty(), isOccupied: false}
					updatedRookSquare := square{pieceAt: newRook(black), isOccupied: true}
					updatedKingSquare := square{pieceAt: newKing(black), isOccupied: true}
					g.board.squares[a8RookCoords.row][a8RookCoords.col] = updatedEmptyRookSquare
					g.board.squares[c8Coords.row][c8Coords.col] = updatedKingSquare
					g.board.squares[d8Coords.row][d8Coords.col] = updatedRookSquare
				} else if endAt == g8Coords {
					updatedEmptyRookSquare := square{pieceAt: newEmpty(), isOccupied: false}
					updatedRookSquare := square{pieceAt: newRook(black), isOccupied: true}
					updatedKingSquare := square{pieceAt: newKing(black), isOccupied: true}
					g.board.squares[h8RookCoords.row][h8RookCoords.col] = updatedEmptyRookSquare
					g.board.squares[g8Coords.row][g8Coords.col] = updatedKingSquare
					g.board.squares[f8Coords.row][f8Coords.col] = updatedRookSquare
				}
				updatedEmptyKingSquare := square{pieceAt: newEmpty(), isOccupied: false}
				g.board.squares[e8Coords.row][e8Coords.col] = updatedEmptyKingSquare

			}
		}
	}

	// check if move leaves own king in check
	for r, row := range g.board.squares {
		for c, square := range row {
			piece := square.pieceAt
			if piece.getPieceColor() == pieceMoved.getPieceColor() {
				continue
			}
			availableMoves := piece.getAvailableMoves(coords{row: r, col: c}, g)
			for _, move := range availableMoves {
				pieceAtPossibleMove := g.board.squares[move.row][move.col].pieceAt
				switch v := pieceAtPossibleMove.(type) {
				case *king:
					if v.getPieceColor() != piece.getPieceColor() {
						g.board.squares = originalSquares
						return originalSquares, ErrInvalidMove
					}
				}
			}
		}
	}

	return g.board.squares, nil
}

func (g *Game) coordsToSquareName(c coords) {
	for squareName, coord := range g.board.coords {
		if c == coord {
			g.enPassantSquare = squareName
			return
		}
	}
}

func (g *Game) updatePassantCoords(pawn *pawn, endingCoords coords) coords {
	verticalOffset := -1
	if pawn.getPieceColor() == black {
		verticalOffset = 1
	}

	return coords{row: endingCoords.row - verticalOffset, col: endingCoords.col}
}

func (g *Game) validateCastling(endAt coords, king *king) error {
	// ending squares after castle
	c1Coords := coords{row: 7, col: 2}
	g1Coords := coords{row: 7, col: 6}
	c8Coords := coords{row: 0, col: 2}
	g8Coords := coords{row: 0, col: 6}
	// goes through check
	d1Coords := coords{row: 7, col: 3}
	f1Coords := coords{row: 7, col: 5}
	d8Coords := coords{row: 0, col: 3}
	f8Coords := coords{row: 0, col: 5}
	// starting king coords
	e1Coords := coords{row: 7, col: 4}
	e8Coords := coords{row: 0, col: 4}

	for r, row := range g.board.squares {
		for c, square := range row {
			pieceAtSquare := square.pieceAt
			if pieceAtSquare.getPieceColor() == king.getPieceColor() {
				continue
			}
			squaresAttacked := pieceAtSquare.getAvailableMoves(coords{row: r, col: c}, g)
			switch pieceAtSquare.(type) {
			case *pawn:
				if pieceAtSquare.getPieceColor() == white {
					if r-1 < 0 {
						break
					}
					if c-1 >= 0 {
						squaresAttacked = append(squaresAttacked, coords{row: r - 1, col: c - 1})
					}
					if c+1 <= 7 {
						squaresAttacked = append(squaresAttacked, coords{row: r - 1, col: c + 1})
					}
				} else if pieceAtSquare.getPieceColor() == black {
					if r+1 > 7 {
						break
					}
					if c-1 >= 0 {
						squaresAttacked = append(squaresAttacked, coords{row: r + 1, col: c - 1})
					}
					if c+1 <= 7 {
						squaresAttacked = append(squaresAttacked, coords{row: r + 1, col: c + 1})
					}
				}
			}

			if endAt == c1Coords && slices.Contains(squaresAttacked, d1Coords) {
				return ErrInvalidMove
			}
			if endAt == c1Coords && slices.Contains(squaresAttacked, e1Coords) {
				return ErrInvalidMove
			}
			if endAt == g1Coords && slices.Contains(squaresAttacked, f1Coords) {
				return ErrInvalidMove
			}
			if endAt == g1Coords && slices.Contains(squaresAttacked, e1Coords) {
				return ErrInvalidMove
			}
			if endAt == c8Coords && slices.Contains(squaresAttacked, d8Coords) {
				return ErrInvalidMove
			}
			if endAt == c8Coords && slices.Contains(squaresAttacked, e8Coords) {
				return ErrInvalidMove
			}
			if endAt == g8Coords && slices.Contains(squaresAttacked, f8Coords) {
				return ErrInvalidMove
			}
			if endAt == g8Coords && slices.Contains(squaresAttacked, e8Coords) {
				return ErrInvalidMove
			}

		}
	}
	return nil
}
