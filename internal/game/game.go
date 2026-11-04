package game

import (
	"errors"
	"fmt"
	"log"
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
	isWhiteTurn         bool
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

	if pieceAtStart.getPieceColor() != g.isWhiteTurn {
		return ErrInvalidTurn
	}

	if pieceAtEnd.getPieceColor() == pieceAtStart.getPieceColor() && pieceAtEnd.getPieceVal() != EmptyVal {
		return ErrInvalidMove
	}

	availableMoves := pieceAtStart.getAvailableMoves(startingCoords, g.board, g)

	log.Printf("available moves: %v", availableMoves)

	if slices.Contains(availableMoves, endingCoords) {
		updatedBoard := g.updateBoard(pieceAtStart, startingCoords, endingCoords)
		g.board = updatedBoard
		g.isWhiteTurn = !g.isWhiteTurn

		switch v := pieceAtStart.(type) {
		case *pawn:
			if distance(startingCoords.row, endingCoords.row) == 2 {
				verticalOffset := -1
				if v.isWhite {
					verticalOffset = 1
				}
				enPassantSquareCoords := coords{row: endingCoords.row + verticalOffset, col: endingCoords.col}
				g.coordsToSquareName(enPassantSquareCoords)
				return nil
			}
		}
		g.enPassantSquare = ""

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

func (g *Game) updateBoard(pieceMoved piece, startAt coords, endAt coords) board {
	updatedBoard := g.board
	updatedStartSquare := square{pieceAt: newEmpty(), isOccupied: false}
	updatedBoard.squares[startAt.row][startAt.col] = updatedStartSquare

	updatedEndSquare := square{pieceAt: pieceMoved, isOccupied: true}
	updatedBoard.squares[endAt.row][endAt.col] = updatedEndSquare

	switch v := pieceMoved.(type) {
	case *pawn:
		if g.enPassantSquare != "" && endAt == g.board.coords[g.enPassantSquare] {
			updatedEnPassantSquare := square{pieceAt: newEmpty(), isOccupied: false}
			updateCoords := g.updatePassantCoords(v, endAt)
			updatedBoard.squares[updateCoords.row][updateCoords.col] = updatedEnPassantSquare
		}
	}

	return updatedBoard
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
	if !pawn.getPieceColor() {
		verticalOffset = 1
	}

	return coords{row: endingCoords.row - verticalOffset, col: endingCoords.col}
}
