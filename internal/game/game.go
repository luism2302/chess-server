package game

import (
	"fmt"
	"slices"
	"strings"
)

type Game struct {
	board       board
	isWhiteTurn bool
}

func New() *Game {
	return &Game{
		board: board{
			squares: startingSquares,
			coords:  initializeCoordsMap(),
		},
		isWhiteTurn: true,
	}
}

func (g *Game) BoardRepr() string {
	var builder strings.Builder
	builder.WriteString(border)
	for i, row := range g.board.squares {
		builder.WriteString(fmt.Sprintf("%d | ", 8-i))
		for _, square := range row {
			builder.WriteString(" ")
			repr, err := pieceRune(square.pieceAt.val)
			if err != nil {
				return ""
			}
			builder.WriteRune(repr)
			builder.WriteString(" ")
		}
		builder.WriteString(" |")
		builder.WriteString("\n\n")
	}
	builder.WriteString(border)
	builder.WriteString(cols)
	return builder.String()
}

func (g *Game) Update(move string) error {
	startingCoords, endingCoords, err := g.validateMove(move)
	if err != nil {
		return err
	}

	pieceAtStart := g.getPieceAtCoords(startingCoords)
	if pieceAtStart.isWhite == g.isWhiteTurn {
		return nil
	}

	availableMoves := getAvailableMoves(pieceAtStart, startingCoords, g.board)

	if slices.Contains(availableMoves, endingCoords) {
		updatedBoard := g.UpdateBoard(pieceAtStart, startingCoords, endingCoords)
		g.board = updatedBoard
		g.isWhiteTurn = !g.isWhiteTurn
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

func (g *Game) UpdateBoard(piece piece, startAt coords, endAt coords) board {
	updatedBoard := g.board
	updatedStartSquare := square{pieceAt: Empty, isOccupied: false}
	updatedBoard.squares[startAt.row][startAt.col] = updatedStartSquare

	updatedEndSquare := square{pieceAt: piece, isOccupied: true}
	updatedBoard.squares[endAt.row][endAt.col] = updatedEndSquare

	return updatedBoard
}
