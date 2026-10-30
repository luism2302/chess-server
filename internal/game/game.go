package game

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"
)

var (
	ErrInvalidTurn   = errors.New("other player's turn")
	ErrSelectedEmpty = errors.New("selected an empty square")
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
	log.Println(builder.String())
	return builder.String()
}

func (g *Game) Update(move string) error {
	startingCoords, endingCoords, err := g.validateMove(move)
	if err != nil {
		return err
	}

	pieceAtStart := g.getPieceAtCoords(startingCoords)
	pieceAtEnd := g.getPieceAtCoords(endingCoords)

	if pieceAtEnd.getPieceColor() == pieceAtStart.getPieceColor() && pieceAtEnd.getPieceVal() != EmptyVal {
		return ErrInvalidMove
	}

	if pieceAtStart.getPieceVal() == EmptyVal {
		return ErrSelectedEmpty
	}

	if pieceAtStart.getPieceColor() == g.isWhiteTurn {
		return ErrInvalidTurn
	}

	availableMoves := pieceAtStart.getAvailableMoves(startingCoords, g.board)

	log.Printf("available moves: %v", availableMoves)

	if slices.Contains(availableMoves, endingCoords) {
		updatedBoard := g.updateBoard(pieceAtStart, startingCoords, endingCoords)
		g.board = updatedBoard
		g.isWhiteTurn = !g.isWhiteTurn

		switch pieceAtStart {
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

func (g *Game) updateBoard(piece piece, startAt coords, endAt coords) board {
	updatedBoard := g.board
	updatedStartSquare := square{pieceAt: newEmpty(), isOccupied: false}
	updatedBoard.squares[startAt.row][startAt.col] = updatedStartSquare

	updatedEndSquare := square{pieceAt: piece, isOccupied: true}
	updatedBoard.squares[endAt.row][endAt.col] = updatedEndSquare

	g.updatePiece(piece, startAt, endAt)

	return updatedBoard
}

func (g *Game) updatePiece(piece piece, startAt coords, endAt coords) {
	switch v := piece.(type) {
	case *pawn:
		if v.isFirstMove {
			v.updateFirstMove()
			if endAt.row-startAt.row == 2 || endAt.row-startAt.row == -2 {
				v.enPassantable = true
			}
		}
	}
}
