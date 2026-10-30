package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateMove(t *testing.T) {
	type testCase struct {
		name                   string
		moveInput              string
		expectedStartingCoords coords
		expectedEndingCoords   coords
		expectedError          bool
	}

	testCases := []testCase{
		{name: "Invalid Row", moveInput: "a10-b2", expectedStartingCoords: coords{}, expectedEndingCoords: coords{}, expectedError: true},
		{name: "Invalid Col", moveInput: "a2-p2", expectedStartingCoords: coords{}, expectedEndingCoords: coords{}, expectedError: true},
		{name: "Valid With Spaces", moveInput: "        b2      -     c3", expectedStartingCoords: coords{row: 6, col: 1}, expectedEndingCoords: coords{row: 5, col: 2}, expectedError: false},
		{name: "Valid With No Spaces", moveInput: "h4-a1", expectedStartingCoords: coords{row: 4, col: 7}, expectedEndingCoords: coords{row: 7, col: 0}, expectedError: false},
		{name: "Valid Edges", moveInput: "a8-h1", expectedStartingCoords: coords{row: 0, col: 0}, expectedEndingCoords: coords{row: 7, col: 7}, expectedError: false},
		{name: "All Spaces", moveInput: "         ", expectedStartingCoords: coords{}, expectedEndingCoords: coords{}, expectedError: true},
		{name: "Blank", moveInput: "", expectedStartingCoords: coords{}, expectedEndingCoords: coords{}, expectedError: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			game := New()
			gotStartingCoords, gotEndingCoords, err := game.validateMove(tc.moveInput)

			if tc.expectedError {
				assert.Error(t, err, "expected an error, but got nil")
			} else {
				assert.NoError(t, err, "expected a nil error")
				assert.Equal(t, tc.expectedStartingCoords, gotStartingCoords, fmt.Sprintf("expected starting coords: %v, but got: %v", tc.expectedStartingCoords, gotStartingCoords))
				assert.Equal(t, tc.expectedEndingCoords, gotEndingCoords, fmt.Sprintf("expected ending coords: %v, but got: %v", tc.expectedEndingCoords, gotEndingCoords))
			}
		})
	}
}

func TestParseFENRank(t *testing.T) {
	type testCase struct {
		name         string
		inputFENRank string
		expectedRank []square
		expectedErr  bool
	}

	testCases := []testCase{
		{
			name:         "initial board, black",
			inputFENRank: "rnbqkbnr",
			expectedRank: []square{
				{pieceAt: newRook(false), isOccupied: true},
				{pieceAt: newKnight(false), isOccupied: true},
				{pieceAt: newBishop(false), isOccupied: true},
				{pieceAt: newQueen(false), isOccupied: true},
				{pieceAt: newKing(false), isOccupied: true},
				{pieceAt: newBishop(false), isOccupied: true},
				{pieceAt: newKnight(false), isOccupied: true},
				{pieceAt: newRook(false), isOccupied: true},
			},
			expectedErr: false,
		},
		{
			name:         "initial board, white",
			inputFENRank: "RNBQKBNR",
			expectedRank: []square{
				{pieceAt: newRook(true), isOccupied: true},
				{pieceAt: newKnight(true), isOccupied: true},
				{pieceAt: newBishop(true), isOccupied: true},
				{pieceAt: newQueen(true), isOccupied: true},
				{pieceAt: newKing(true), isOccupied: true},
				{pieceAt: newBishop(true), isOccupied: true},
				{pieceAt: newKnight(true), isOccupied: true},
				{pieceAt: newRook(true), isOccupied: true},
			},
			expectedErr: false,
		},
		{
			name:         "all empty",
			inputFENRank: "8",
			expectedRank: []square{
				{pieceAt: newEmpty(), isOccupied: false},
				{pieceAt: newEmpty(), isOccupied: false},
				{pieceAt: newEmpty(), isOccupied: false},
				{pieceAt: newEmpty(), isOccupied: false},
				{pieceAt: newEmpty(), isOccupied: false},
				{pieceAt: newEmpty(), isOccupied: false},
				{pieceAt: newEmpty(), isOccupied: false},
				{pieceAt: newEmpty(), isOccupied: false},
			},
			expectedErr: false,
		},
		{
			name:         "invalid pieces",
			inputFENRank: "rnbqktlo",
			expectedRank: []square{},
			expectedErr:  true,
		},
		{
			name:         "row too short",
			inputFENRank: "RNB",
			expectedRank: []square{},
			expectedErr:  true,
		},
		{
			name:         "row too long",
			inputFENRank: "RNB8",
			expectedRank: []square{},
			expectedErr:  true,
		},
		{
			name:         "valid row, rand position",
			inputFENRank: "pQ1r1nPP",
			expectedRank: []square{
				{pieceAt: newPawn(false), isOccupied: true},
				{pieceAt: newQueen(true), isOccupied: true},
				{pieceAt: newEmpty(), isOccupied: false},
				{pieceAt: newRook(false), isOccupied: true},
				{pieceAt: newEmpty(), isOccupied: false},
				{pieceAt: newKnight(false), isOccupied: true},
				{pieceAt: newPawn(true), isOccupied: true},
				{pieceAt: newPawn(true), isOccupied: true},
			},
			expectedErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseFENRank(tc.inputFENRank)
			if tc.expectedErr {
				assert.Error(t, err, "expected an error, but got nil")
			} else {
				assert.NoError(t, err, "expected a nil error")
				assert.Equal(t, got, tc.expectedRank)
			}
		})
	}
}

func TestParseFENBoard(t *testing.T) {
	type testCase struct {
		name            string
		inputFENBoard   string
		expectedSquares [][]square
		expectedErr     bool
	}

	testCases := []testCase{
		{
			name:          "starting board",
			inputFENBoard: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
			expectedSquares: [][]square{
				{{newRook(false), true}, {newKnight(false), true}, {newBishop(false), true}, {newQueen(false), true}, {newKing(false), true}, {newBishop(false), true}, {newKnight(false), true}, {newRook(false), true}},
				{{newPawn(false), true}, {newPawn(false), true}, {newPawn(false), true}, {newPawn(false), true}, {newPawn(false), true}, {newPawn(false), true}, {newPawn(false), true}, {newPawn(false), true}},
				{{newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
				{{newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
				{{newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
				{{newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
				{{newPawn(true), true}, {newPawn(true), true}, {newPawn(true), true}, {newPawn(true), true}, {newPawn(true), true}, {newPawn(true), true}, {newPawn(true), true}, {newPawn(true), true}},
				{{newRook(true), true}, {newKnight(true), true}, {newBishop(true), true}, {newQueen(true), true}, {newKing(true), true}, {newBishop(true), true}, {newKnight(true), true}, {newRook(true), true}},
			},
			expectedErr: false,
		},
		{
			name:          "random valid position",
			inputFENBoard: "rp3pk1/8/npq3br/1n1pb3/K1R2PPR/2N2Q2/P4B1N/B2P4",
			expectedSquares: [][]square{
				{{newRook(false), true}, {newPawn(false), true}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newPawn(false), true}, {newKing(false), true}, {newEmpty(), false}},
				{{newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
				{{newKnight(false), true}, {newPawn(false), true}, {newQueen(false), true}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newBishop(false), true}, {newRook(false), true}},
				{{newEmpty(), false}, {newKnight(false), true}, {newEmpty(), false}, {newPawn(false), true}, {newBishop(false), true}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
				{{newKing(true), true}, {newEmpty(), false}, {newRook(true), true}, {newEmpty(), false}, {newEmpty(), false}, {newPawn(true), true}, {newPawn(true), true}, {newRook(true), true}},
				{{newEmpty(), false}, {newEmpty(), false}, {newKnight(true), true}, {newEmpty(), false}, {newEmpty(), false}, {newQueen(true), true}, {newEmpty(), false}, {newEmpty(), false}},
				{{newPawn(true), true}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newBishop(true), true}, {newEmpty(), false}, {newKnight(true), true}},
				{{newBishop(true), true}, {newEmpty(), false}, {newEmpty(), false}, {newPawn(true), true}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
			},
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseFENBoard(tc.inputFENBoard)
			if tc.expectedErr {
				assert.Error(t, err, "expected an error, but got nil")
			} else {
				assert.NoError(t, err, "expected a nil error")
				assert.Equal(t, got, tc.expectedSquares)
			}
		})
	}
}
