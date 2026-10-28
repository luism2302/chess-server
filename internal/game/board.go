package game

import (
	"strconv"
	"strings"
)

type board struct {
	squares [][]square
	coords  map[string]coords
}

type square struct {
	pieceAt    piece
	isOccupied bool
}

type coords struct {
	row int
	col int
}

var (
	border = "+  -  -  -  -  -  -  -  -  +\n"
	cols   = "   a   b  c  d  e  f  g  h  "
)

var startingSquares = [][]square{
	{{WhiteRook, true}, {WhiteKnight, true}, {WhiteBishop, true}, {WhiteQueen, true}, {WhiteKing, true}, {WhiteBishop, true}, {WhiteKnight, true}, {WhiteRook, true}},
	{{WhitePawn, true}, {WhitePawn, true}, {WhitePawn, true}, {WhitePawn, true}, {WhitePawn, true}, {WhitePawn, true}, {WhitePawn, true}, {WhitePawn, true}},
	{{Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}},
	{{Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}},
	{{Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}},
	{{Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}, {Empty, false}},
	{{BlackPawn, true}, {BlackPawn, true}, {BlackPawn, true}, {BlackPawn, true}, {BlackPawn, true}, {BlackPawn, true}, {BlackPawn, true}, {BlackPawn, true}},
	{{BlackRook, true}, {BlackKnight, true}, {BlackBishop, true}, {BlackQueen, true}, {BlackKing, true}, {BlackBishop, true}, {BlackKnight, true}, {BlackRook, true}},
}

var letterToCol = map[rune]int{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
	'h': 7,
}

func initializeCoordsMap() map[string]coords {
	m := make(map[string]coords)
	var builder strings.Builder
	for ch := 'a'; ch <= 'h'; ch++ {
		colCoord := letterToCol[ch]
		for rowCoord := 8; rowCoord >= 1; rowCoord-- {
			row := strconv.Itoa(rowCoord)
			builder.WriteRune(ch)
			builder.WriteString(row)
			key := builder.String()
			m[key] = coords{row: 8 - rowCoord, col: colCoord}
			builder.Reset()
		}
	}

	return m
}
