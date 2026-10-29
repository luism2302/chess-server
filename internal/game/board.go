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
	{{newRook(true), true}, {newKnight(true), true}, {newBishop(true), true}, {newQueen(true), true}, {newKing(true), true}, {newBishop(true), true}, {newKnight(true), true}, {newRook(true), true}},
	{{newPawn(true), true}, {newPawn(true), true}, {newPawn(true), true}, {newPawn(true), true}, {newPawn(true), true}, {newPawn(true), true}, {newPawn(true), true}, {newPawn(true), true}},
	{{newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
	{{newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
	{{newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
	{{newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
	{{newPawn(false), false}, {newPawn(false), false}, {newPawn(false), false}, {newPawn(false), false}, {newPawn(false), false}, {newPawn(false), false}, {newPawn(false), false}, {newPawn(false), false}},
	{{newRook(false), false}, {newKnight(false), false}, {newBishop(false), false}, {newQueen(false), false}, {newKing(false), false}, {newBishop(false), false}, {newKnight(false), false}, {newRook(false), false}},
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
