package game

import (
	"fmt"
	"strings"
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
			game, err := New("")
			if err != nil {
				t.Fatal(err)
			}
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
				{pieceAt: newRook(black), isOccupied: true},
				{pieceAt: newKnight(black), isOccupied: true},
				{pieceAt: newBishop(black), isOccupied: true},
				{pieceAt: newQueen(black), isOccupied: true},
				{pieceAt: newKing(black), isOccupied: true},
				{pieceAt: newBishop(black), isOccupied: true},
				{pieceAt: newKnight(black), isOccupied: true},
				{pieceAt: newRook(black), isOccupied: true},
			},
			expectedErr: false,
		},
		{
			name:         "initial board, white",
			inputFENRank: "RNBQKBNR",
			expectedRank: []square{
				{pieceAt: newRook(white), isOccupied: true},
				{pieceAt: newKnight(white), isOccupied: true},
				{pieceAt: newBishop(white), isOccupied: true},
				{pieceAt: newQueen(white), isOccupied: true},
				{pieceAt: newKing(white), isOccupied: true},
				{pieceAt: newBishop(white), isOccupied: true},
				{pieceAt: newKnight(white), isOccupied: true},
				{pieceAt: newRook(white), isOccupied: true},
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
				{pieceAt: newPawn(black), isOccupied: true},
				{pieceAt: newQueen(white), isOccupied: true},
				{pieceAt: newEmpty(), isOccupied: false},
				{pieceAt: newRook(black), isOccupied: true},
				{pieceAt: newEmpty(), isOccupied: false},
				{pieceAt: newKnight(black), isOccupied: true},
				{pieceAt: newPawn(white), isOccupied: true},
				{pieceAt: newPawn(white), isOccupied: true},
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
				{{newRook(black), true}, {newKnight(black), true}, {newBishop(black), true}, {newQueen(black), true}, {newKing(black), true}, {newBishop(black), true}, {newKnight(black), true}, {newRook(black), true}},
				{{newPawn(black), true}, {newPawn(black), true}, {newPawn(black), true}, {newPawn(black), true}, {newPawn(black), true}, {newPawn(black), true}, {newPawn(black), true}, {newPawn(black), true}},
				{{newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
				{{newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
				{{newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
				{{newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
				{{newPawn(white), true}, {newPawn(white), true}, {newPawn(white), true}, {newPawn(white), true}, {newPawn(white), true}, {newPawn(white), true}, {newPawn(white), true}, {newPawn(white), true}},
				{{newRook(white), true}, {newKnight(white), true}, {newBishop(white), true}, {newQueen(white), true}, {newKing(white), true}, {newBishop(white), true}, {newKnight(white), true}, {newRook(white), true}},
			},
			expectedErr: false,
		},
		{
			name:          "random valid position",
			inputFENBoard: "rp3pk1/8/npq3br/1n1pb3/K1R2PPR/2N2Q2/P4B1N/B2P4",
			expectedSquares: [][]square{
				{{newRook(black), true}, {newPawn(black), true}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newPawn(black), true}, {newKing(black), true}, {newEmpty(), false}},
				{{newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
				{{newKnight(black), true}, {newPawn(black), true}, {newQueen(black), true}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newBishop(black), true}, {newRook(black), true}},
				{{newEmpty(), false}, {newKnight(black), true}, {newEmpty(), false}, {newPawn(black), true}, {newBishop(black), true}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
				{{newKing(white), true}, {newEmpty(), false}, {newRook(white), true}, {newEmpty(), false}, {newEmpty(), false}, {newPawn(white), true}, {newPawn(white), true}, {newRook(white), true}},
				{{newEmpty(), false}, {newEmpty(), false}, {newKnight(white), true}, {newEmpty(), false}, {newEmpty(), false}, {newQueen(white), true}, {newEmpty(), false}, {newEmpty(), false}},
				{{newPawn(white), true}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newBishop(white), true}, {newEmpty(), false}, {newKnight(white), true}},
				{{newBishop(white), true}, {newEmpty(), false}, {newEmpty(), false}, {newPawn(white), true}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
			},
			expectedErr: false,
		},
		{
			name:          "random invalid position",
			inputFENBoard: "rp3pk1/8/npq3br/1n1pb/K1R2PPR/2N2Q2/P4B1N/B2P4",
			expectedSquares: [][]square{
				{{newRook(black), true}, {newPawn(black), true}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newPawn(black), true}, {newKing(black), true}, {newEmpty(), false}},
				{{newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
				{{newKnight(black), true}, {newPawn(black), true}, {newQueen(black), true}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newBishop(black), true}, {newRook(black), true}},
				{{newEmpty(), false}, {newKnight(black), true}, {newEmpty(), false}, {newPawn(black), true}, {newBishop(black), true}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
				{{newKing(white), true}, {newEmpty(), false}, {newRook(white), true}, {newEmpty(), false}, {newEmpty(), false}, {newPawn(white), true}, {newPawn(white), true}, {newRook(white), true}},
				{{newEmpty(), false}, {newEmpty(), false}, {newKnight(white), true}, {newEmpty(), false}, {newEmpty(), false}, {newQueen(white), true}, {newEmpty(), false}, {newEmpty(), false}},
				{{newPawn(white), true}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newBishop(white), true}, {newEmpty(), false}, {newKnight(black), true}},
				{{newBishop(white), true}, {newEmpty(), false}, {newEmpty(), false}, {newPawn(white), true}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}, {newEmpty(), false}},
			},
			expectedErr: true,
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

func TestParseFENCastling(t *testing.T) {
	type testCase struct {
		name          string
		inputCastling string
		expectedWhite castlingRights
		expectedBlack castlingRights
		expectedError bool
	}

	testCases := []testCase{
		{
			name:          "initial rights",
			inputCastling: "KQkq",
			expectedWhite: castlingRights{canKingSide: true, canQueenSide: true},
			expectedBlack: castlingRights{canKingSide: true, canQueenSide: true},
			expectedError: false,
		},
		{
			name:          "only white, queen side",
			inputCastling: "Q",
			expectedWhite: castlingRights{canKingSide: false, canQueenSide: true},
			expectedBlack: castlingRights{canKingSide: false, canQueenSide: false},
			expectedError: false,
		},
		{
			name:          "white king side, black queen side",
			inputCastling: "Kq",
			expectedWhite: castlingRights{canKingSide: true, canQueenSide: false},
			expectedBlack: castlingRights{canKingSide: false, canQueenSide: true},
			expectedError: false,
		},
		{
			name:          "invalid castling",
			inputCastling: "KKKKKKq",
			expectedWhite: castlingRights{},
			expectedBlack: castlingRights{},
			expectedError: true,
		},
		{
			name:          "invalid castling symbol",
			inputCastling: "Kl",
			expectedWhite: castlingRights{},
			expectedBlack: castlingRights{},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotWhite, gotBlack, err := parseFENCastling(tc.inputCastling)
			if tc.expectedError {
				assert.Error(t, err, "expected an error but got nil")
			} else {
				assert.NoError(t, err, "expected a nil error")
				assert.Equal(t, gotWhite, tc.expectedWhite)
				assert.Equal(t, gotBlack, tc.expectedBlack)
			}
		})
	}
}

func TestEnPassantFromFEN(t *testing.T) {
	type testCase struct {
		name               string
		initialFEN         string
		captureMove        string
		takeSquare         string
		defenderPawnSquare string
		attackerPawnSquare string
		validEnPassant     bool
		turnAfterEnPassant color // true white false black
	}

	testCases := []testCase{
		{
			name:               "enPassant in f6, to the right, white",
			initialFEN:         "rnbqkbnr/pppp2pp/8/4Pp2/8/8/PPP1PPPP/RNBQKBNR w KQkq f6 0 1",
			captureMove:        "e5-f6",
			takeSquare:         "f6",
			defenderPawnSquare: "f5",
			attackerPawnSquare: "e5",
			validEnPassant:     true,
			turnAfterEnPassant: black,
		},
		{
			name:               "enPassant in a6, to the left, white",
			initialFEN:         "rnbqkbnr/1p1ppppp/8/pPp5/8/8/P1PPPPPP/RNBQKBNR w KQkq a6 0 1",
			captureMove:        "b5-a6",
			takeSquare:         "a6",
			defenderPawnSquare: "a5",
			attackerPawnSquare: "b5",
			validEnPassant:     true,
			turnAfterEnPassant: black,
		},
		{
			name:               "enPassant in e3, to the right, black",
			initialFEN:         "rnbqkbnr/ppp1pppp/8/2P5/3pP3/8/PP1P1PPP/RNBQKBNR b KQkq e3 0 1",
			captureMove:        "d4-e3",
			takeSquare:         "e3",
			defenderPawnSquare: "e4",
			attackerPawnSquare: "d4",
			validEnPassant:     true,
			turnAfterEnPassant: white,
		},
		{
			name:               "enPassant in a3, to the left, black",
			initialFEN:         "rnbqkbnr/pp1ppppp/8/8/PpP5/8/3PPPPP/RNBQKBNR b KQkq a3 0 1",
			captureMove:        "b4-a3",
			takeSquare:         "a3",
			defenderPawnSquare: "a4",
			attackerPawnSquare: "b4",
			validEnPassant:     true,
			turnAfterEnPassant: white,
		},
		{
			name:               "enPassant in a3, to the left, black, but not performed",
			initialFEN:         "rnbqkbnr/pp1ppppp/8/8/PpP5/8/3PPPPP/RNBQKBNR b KQkq a3 0 1",
			captureMove:        "b4-b3",
			takeSquare:         "a3",
			defenderPawnSquare: "a4",
			attackerPawnSquare: "b4",
			validEnPassant:     false,
			turnAfterEnPassant: white,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			game, err := New(tc.initialFEN)
			if err != nil {
				t.Fatalf("error creating game: %s", err.Error())
			}

			if err := callMoves(t, game, tc.captureMove); err != nil {
				t.Fatal(err.Error())
			}

			takeToSquareCoords := game.board.coords[tc.takeSquare]
			attackerSquareCoords := game.board.coords[tc.attackerPawnSquare]
			defenderSquareCoords := game.board.coords[tc.defenderPawnSquare]

			takeToSquare := game.board.squares[takeToSquareCoords.row][takeToSquareCoords.col]
			attackerSquare := game.board.squares[attackerSquareCoords.row][attackerSquareCoords.col]
			defenderSquare := game.board.squares[defenderSquareCoords.row][defenderSquareCoords.col]

			if !tc.validEnPassant { // nothing should happen
				// square where attacker pawn will end up
				assert.Equal(t, newEmpty(), takeToSquare.pieceAt, "no enpassant, expected en passant square to be an empty square")
				assert.False(t, takeToSquare.isOccupied, "no enpassant, expected enpassant ssquare to not be occupied")
				// square where attacker starts
				assert.Equal(t, newPawn(tc.turnAfterEnPassant), defenderSquare.pieceAt, "no enpasant, expected defender square to be the same")
				assert.True(t, defenderSquare.isOccupied, "no enpassant, expected defender square to not be occupied")

			} else {
				// square where attacker pawn will end up
				testColor := white
				if tc.turnAfterEnPassant == white {
					testColor = black
				}
				assert.Equal(t, newPawn(testColor), takeToSquare.pieceAt)
				assert.True(t, takeToSquare.isOccupied)
				// square where attacker starts
				assert.Equal(t, newEmpty(), attackerSquare.pieceAt)
				assert.False(t, attackerSquare.isOccupied)
				// square where defender starts
				assert.Equal(t, newEmpty(), defenderSquare.pieceAt)
				assert.False(t, defenderSquare.isOccupied)
				assert.True(t, game.enPassantSquare == "")
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	type testCase struct {
		name                    string
		startingFEN             string
		moveToTest              string
		expectedEmptySquareErr  bool
		expectedWrongTurnErr    bool
		expectedInvalidMoveErr  bool
		expectedPass            bool
		pieceMoved              piece
		expectedEnPassantSquare string
	}

	testCases := []testCase{
		{
			name:                   "starting empty square",
			startingFEN:            "rnbqkbnr/pppp2pp/8/5p2/2PQ4/8/PP2PPPP/RNB1KBNR w KQkq f6 0 1",
			moveToTest:             "e3-e4",
			expectedEmptySquareErr: true,
			expectedWrongTurnErr:   false,
			expectedInvalidMoveErr: false,
			expectedPass:           false,
		},
		{
			name:                   "not blacks turn",
			startingFEN:            "rnbqk2r/ppp2ppp/5n2/2bPp3/2P5/P4N2/1P1P1PPP/RNBQKB1R w KQkq - 0 1",
			moveToTest:             "e5-e4",
			expectedEmptySquareErr: false,
			expectedWrongTurnErr:   true,
			expectedInvalidMoveErr: false,
			expectedPass:           false,
		},
		{
			name:                   "not whites turn",
			startingFEN:            "rnbqkbnr/pppp1ppp/8/4p3/3P1P2/8/PPP1P1PP/RNBQKBNR b KQkq f3 0 1",
			moveToTest:             "d4-e5",
			expectedEmptySquareErr: false,
			expectedWrongTurnErr:   true,
			expectedInvalidMoveErr: false,
			expectedPass:           false,
		},
		{
			name:                   "white tries to caputre same color piece",
			startingFEN:            "rnbqkbnr/ppp3pp/3p1p2/4P3/5P2/8/PPP1P1PP/RNBQKBNR w KQkq - 0 1",
			moveToTest:             "f4-e5",
			expectedEmptySquareErr: false,
			expectedWrongTurnErr:   false,
			expectedInvalidMoveErr: true,
			expectedPass:           false,
		},
		{
			name:                   "invalid move from black",
			startingFEN:            "rnbqkbnr/ppp2ppp/8/3pP3/2P5/8/PP2PPPP/RNBQKBNR b KQkq c3 0 1",
			moveToTest:             "d5-c5",
			expectedEmptySquareErr: false,
			expectedWrongTurnErr:   false,
			expectedInvalidMoveErr: true,
			expectedPass:           false,
		},
		{
			name:                   "pawn move to an occupied square",
			startingFEN:            "rnbqkbnr/ppp1pppp/8/3p4/3P4/8/PPP1PPPP/RNBQKBNR w KQkq d6 0 1",
			moveToTest:             "d4-d5",
			expectedEmptySquareErr: false,
			expectedWrongTurnErr:   false,
			expectedInvalidMoveErr: true,
			expectedPass:           false,
		},
		{
			name:                   "invalid opening move, left",
			startingFEN:            "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			moveToTest:             "e2-d3",
			expectedEmptySquareErr: false,
			expectedWrongTurnErr:   false,
			expectedInvalidMoveErr: true,
			expectedPass:           false,
		},
		{
			name:                   "invalid opening move, right",
			startingFEN:            "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			moveToTest:             "e2-f3",
			expectedEmptySquareErr: false,
			expectedWrongTurnErr:   false,
			expectedInvalidMoveErr: true,
			expectedPass:           false,
		},
		{
			name:                   "invalid two square move",
			startingFEN:            "rnbqkbnr/pppppppp/8/8/8/4P3/PPPP1PPP/RNBQKBNR w KQkq - 0 1",
			moveToTest:             "e3-e5",
			expectedEmptySquareErr: false,
			expectedWrongTurnErr:   false,
			expectedInvalidMoveErr: true,
			expectedPass:           false,
		},
		{
			name:                   "wrong formatted move",
			startingFEN:            "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			moveToTest:             "e2e4",
			expectedEmptySquareErr: false,
			expectedWrongTurnErr:   false,
			expectedInvalidMoveErr: true,
			expectedPass:           false,
		},
		{
			name:                    "valid white take move, right",
			startingFEN:             "rnbqkbnr/pppp2pp/8/4pp2/2PP4/8/PP2PPPP/RNBQKBNR w KQkq f6 0 1",
			moveToTest:              "d4-e5",
			expectedEmptySquareErr:  false,
			expectedWrongTurnErr:    false,
			expectedInvalidMoveErr:  false,
			pieceMoved:              newPawn(white),
			expectedPass:            true,
			expectedEnPassantSquare: "",
		},
		{
			name:                    "valid white take move, left",
			startingFEN:             "rnbqkbnr/1ppppppp/8/p7/1P6/8/P1PPPPPP/RNBQKBNR w KQkq a6 0 1",
			moveToTest:              "b4-a5",
			expectedEmptySquareErr:  false,
			expectedWrongTurnErr:    false,
			expectedInvalidMoveErr:  false,
			pieceMoved:              newPawn(white),
			expectedPass:            true,
			expectedEnPassantSquare: "",
		},
		{
			name:                    "valid black move",
			startingFEN:             "rnbqkbnr/ppppp1pp/8/5p2/3PP3/8/PPP2PPP/RNBQKBNR b KQkq d3 0 1",
			moveToTest:              "f5-f4",
			expectedEmptySquareErr:  false,
			expectedWrongTurnErr:    false,
			expectedInvalidMoveErr:  false,
			pieceMoved:              newPawn(black),
			expectedPass:            true,
			expectedEnPassantSquare: "",
		},
		{
			name:                    "valid white 2 square move",
			startingFEN:             "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			moveToTest:              "a2-a4",
			expectedEmptySquareErr:  false,
			expectedWrongTurnErr:    false,
			expectedInvalidMoveErr:  false,
			pieceMoved:              newPawn(white),
			expectedPass:            true,
			expectedEnPassantSquare: "a3",
		},
		{
			name:                    "valid black 2 square move",
			startingFEN:             "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq d3 0 1",
			moveToTest:              "e7-e5",
			expectedEmptySquareErr:  false,
			expectedWrongTurnErr:    false,
			expectedInvalidMoveErr:  false,
			pieceMoved:              newPawn(black),
			expectedPass:            true,
			expectedEnPassantSquare: "e6",
		},
		{
			name:                   "invalid two square move from white, piece in the middle",
			startingFEN:            "rnbqkbnr/ppp1pppp/8/8/5P2/2Pp4/PP1PP1PP/RNBQKBNR w KQkq - 0 1",
			moveToTest:             "d2-d4",
			expectedEmptySquareErr: false,
			expectedWrongTurnErr:   false,
			expectedInvalidMoveErr: true,
			pieceMoved:             newPawn(white),
			expectedPass:           false,
		},
		{
			name:                   "invalid two square move from white, piece at the end",
			startingFEN:            "rnbqkbnr/pppp1ppp/8/3P4/4p3/8/PPP1PPPP/RNBQKBNR w KQkq - 0 1",
			moveToTest:             "e2-e4",
			expectedEmptySquareErr: false,
			expectedWrongTurnErr:   false,
			expectedInvalidMoveErr: true,
			pieceMoved:             newPawn(white),
			expectedPass:           false,
		},
		{
			name:                   "invalid two square move from black, already moved",
			startingFEN:            "rnbqkbnr/ppppp2p/6p1/4Pp2/3P4/8/PPP2PPP/RNBQKBNR b KQkq d3 0 1",
			moveToTest:             "g6-g4",
			expectedEmptySquareErr: false,
			expectedWrongTurnErr:   false,
			expectedInvalidMoveErr: true,
			pieceMoved:             newPawn(black),
			expectedPass:           false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gameToUpdate, err := New(tc.startingFEN)
			if err != nil {
				t.Fatalf("error creating game: %s", err.Error())
			}
			gameToCheckAgainst, err := New(tc.startingFEN)
			if err != nil {
				t.Fatalf("error creating game: %s", err.Error())
			}

			err = gameToUpdate.Update(tc.moveToTest)
			if tc.expectedPass {
				assert.NoError(t, err, "expected a nil error")
				moveParts := strings.Split(tc.moveToTest, "-")
				startingCoords := gameToUpdate.board.coords[moveParts[0]]
				endingCoords := gameToUpdate.board.coords[moveParts[1]]

				// piece is at ending
				assert.Equal(t, gameToUpdate.board.squares[endingCoords.row][endingCoords.col].pieceAt, tc.pieceMoved, "piece moved not at ending coords")
				assert.True(t, gameToUpdate.board.squares[endingCoords.row][endingCoords.col].isOccupied, "ending coords are empty")
				// starting square is empty
				assert.Equal(t, gameToUpdate.board.squares[startingCoords.row][startingCoords.col].pieceAt, newEmpty(), "starting square is not empty")
				assert.False(t, gameToUpdate.board.squares[startingCoords.row][startingCoords.col].isOccupied, "starting square is occupied")
				// turn changed
				assert.NotEqual(t, tc.pieceMoved.getPieceColor(), gameToUpdate.colorTurn, "turn didnt change")
				// chech enPassant square
				assert.Equal(t, tc.expectedEnPassantSquare, gameToUpdate.enPassantSquare, "expected different enpassant square")

			} else {
				assert.Error(t, err, "expected a non nil error")
				if tc.expectedEmptySquareErr {
					assert.ErrorIs(t, err, ErrSelectedEmpty, "expected an empty square error")
					assert.Equal(t, gameToUpdate, gameToCheckAgainst, "expected games to be the same")
				} else if tc.expectedWrongTurnErr {
					assert.ErrorIs(t, err, ErrInvalidTurn, "expected a wrong turn error")
					assert.Equal(t, gameToUpdate, gameToCheckAgainst, "expected games to be the same")
				} else if tc.expectedInvalidMoveErr {
					assert.ErrorIs(t, err, ErrInvalidMove, "expected an invalid move error")
					assert.Equal(t, gameToUpdate, gameToCheckAgainst, "expected games to be the same")
				}
			}
		})
	}
}

func TestGetMovesKnight(t *testing.T) {
	type testCase struct {
		name                   string
		startingFen            string
		startingCoords         coords
		expectedAvailableMoves []coords
	}

	testCases := []testCase{
		{
			name:                   "b1 white knight starting moves",
			startingFen:            "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			startingCoords:         coords{row: 7, col: 1},
			expectedAvailableMoves: []coords{{row: 5, col: 0}, {row: 5, col: 2}},
		},
		{
			name:                   "g1 white knight starting moves",
			startingFen:            "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			startingCoords:         coords{row: 7, col: 6},
			expectedAvailableMoves: []coords{{row: 5, col: 5}, {row: 5, col: 7}},
		},
		{
			name:                   "b8 knight starting moves",
			startingFen:            "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq d3 0 1",
			startingCoords:         coords{row: 0, col: 1},
			expectedAvailableMoves: []coords{{row: 2, col: 0}, {row: 2, col: 2}},
		},
		{
			name:                   "g8 knight starting moves",
			startingFen:            "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq d3 0 1",
			startingCoords:         coords{row: 0, col: 6},
			expectedAvailableMoves: []coords{{row: 2, col: 5}, {row: 2, col: 7}},
		},
		{
			name:                   "white knight in the middle, all moves available",
			startingFen:            "rnbqkbnr/pppppp2/8/3N2pp/8/8/PPPPPPPP/R1BQKBNR w KQkq g6 0 1",
			startingCoords:         coords{row: 3, col: 3},
			expectedAvailableMoves: []coords{{row: 1, col: 2}, {row: 2, col: 1}, {row: 1, col: 4}, {row: 4, col: 1}, {row: 5, col: 2}, {row: 5, col: 4}, {row: 4, col: 5}, {row: 2, col: 5}},
		},
		{
			name:                   "black knight in the middle, some moves available",
			startingFen:            "rnbqkbr1/pppp1p1p/5n2/8/2P1p1p1/1PNPP3/PBQ2PPP/R3KBNR b KQq - 0 1",
			startingCoords:         coords{row: 2, col: 5},
			expectedAvailableMoves: []coords{{row: 3, col: 3}, {row: 3, col: 7}},
		},
		{
			name:                   "black knight at the edge of board",
			startingFen:            "rnbqkb1r/pppppppp/7n/8/5P2/8/PPPPP1PP/RNBQKBNR w KQkq - 0 1",
			startingCoords:         coords{row: 2, col: 7},
			expectedAvailableMoves: []coords{{row: 0, col: 6}, {row: 3, col: 5}, {row: 4, col: 6}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			game, err := New(tc.startingFen)
			if err != nil {
				t.Fatalf("error starting game: %s", err.Error())
			}

			knight := game.getPieceAtCoords(tc.startingCoords)
			moves := knight.getAvailableMoves(tc.startingCoords, game)

			assert.ElementsMatch(t, tc.expectedAvailableMoves, moves, "expected different moves...")
		})
	}
}

func TestKnightUpdateGame(t *testing.T) {
	type expectedErr struct {
		expected bool
		errType  error
	}

	type testCase struct {
		name        string
		startingFEN string
		move        string
		expectedErr expectedErr
		pieceMoved  piece
	}

	testCases := []testCase{
		{
			name:        "invalid white knight move, starting pos",
			move:        "b1-d2",
			startingFEN: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expectedErr: expectedErr{expected: true, errType: ErrInvalidMove},
			pieceMoved:  newKnight(white),
		},
		{
			name:        "black knight opening move, valid",
			move:        "g8-f6",
			startingFEN: "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq d3 0 1",
			expectedErr: expectedErr{expected: false},
			pieceMoved:  newKnight(black),
		},
		{
			name:        "white knight takes pawn at d5, valid",
			move:        "c3-d5",
			startingFEN: "rnbqkb1r/ppp1pppp/5n2/3p4/2P5/2N5/PP1PPPPP/R1BQKBNR w KQkq d6 0 1",
			expectedErr: expectedErr{expected: false},
			pieceMoved:  newKnight(white),
		},
		{
			name:        "invalid turn",
			move:        "c3-d5",
			startingFEN: "r1bqkb1r/ppppppp1/2n2n2/7p/2PP4/2N2N2/PP2PPPP/R1BQKB1R b KQkq - 0 1",
			expectedErr: expectedErr{expected: true, errType: ErrInvalidTurn},
			pieceMoved:  newKnight(white),
		},
		{
			name:        "wrong formatted move",
			move:        "f3e4",
			startingFEN: "rnbqkbnr/ppp2ppp/8/3pp3/8/2N2N2/PPPPPPPP/R1BQKB1R w KQkq e6 0 1",
			expectedErr: expectedErr{expected: true, errType: ErrInvalidMove},
			pieceMoved:  newKnight(white),
		},
		{
			name:        "black knight takes pawn at d4",
			move:        "c6-d4",
			startingFEN: "r1bqkbnr/pppppppp/2n5/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq - 0 1",
			expectedErr: expectedErr{expected: false},
			pieceMoved:  newKnight(black),
		},
		{
			name:        "white knight takes pawn at f3",
			move:        "e5-f3",
			startingFEN: "rnbqkbnr/pppp2pp/8/4N3/8/2N1Pp2/PPPP1PPP/R1BQKB1R w KQkq - 0 1",
			expectedErr: expectedErr{expected: false},
			pieceMoved:  newKnight(white),
		},
		{
			name:        "white knight has no moves",
			move:        "f3-d2",
			startingFEN: "5b1r/prqkppp1/b1np1n1p/4B1P1/3P3N/1P3N2/P2KPP1P/R4BR1 w - - 0 1",
			expectedErr: expectedErr{expected: true, errType: ErrInvalidMove},
			pieceMoved:  newKnight(white),
		},
		{
			name:        "black knight at edege of board",
			move:        "a1-b3",
			startingFEN: "r1bqkbnr/pppppppp/8/3P4/8/1P3N2/P2QPPPP/nNB1KB1R b Kkq - 0 1",
			expectedErr: expectedErr{expected: false},
			pieceMoved:  newKnight(black),
		},
		{
			name:        "black knight takes at h1",
			move:        "f2-h1",
			startingFEN: "rnbqkb1r/pppppppp/8/8/4P3/2N2N2/PPPPBnPP/R1BQK2R b KQkq - 0 1",
			expectedErr: expectedErr{expected: false},
			pieceMoved:  newKnight(black),
		},
		{
			name:        "white knight takes at h8",
			move:        "f7-h8",
			startingFEN: "r1q1kbnr/pppb1Npp/2np4/4p2Q/4P3/8/PPPP1PPP/RNB1KB1R w KQkq - 0 1",
			expectedErr: expectedErr{expected: false},
			pieceMoved:  newKnight(white),
		},
		{
			name:        "white knight surrounded",
			move:        "d5-c3",
			startingFEN: "rnb4r/p1k1pp2/1p3np1/2pN4/1q3b2/1Pp1p2P/P1PPPPP1/R1BQKBNR w KQ - 0 1",
			expectedErr: expectedErr{expected: false},
			pieceMoved:  newKnight(white),
		},
		{
			name:        "black knight surrounded",
			move:        "e4-g3",
			startingFEN: "rnbqkb1r/pppppppp/3Q1N2/2B3N1/1P2n3/2P3P1/P2PPP1P/R3KB1R b KQkq - 0 1",
			expectedErr: expectedErr{expected: false},
			pieceMoved:  newKnight(black),
		},
		{
			name:        "white knight surrounded by own pieces",
			move:        "d3-b4",
			startingFEN: "rnbqkbnr/pppppp2/6pp/2P1P3/1Q3N2/3N4/PP1P1PPP/1RB1KB1R w Kkq - 0 1",
			expectedErr: expectedErr{expected: true, errType: ErrInvalidMove},
			pieceMoved:  newKnight(white),
		},
		{
			name:        "invalid white starting move",
			move:        "g1-g3",
			startingFEN: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expectedErr: expectedErr{expected: true, errType: ErrInvalidMove},
			pieceMoved:  newKnight(white),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gameToUpdate, err := New(tc.startingFEN)
			if err != nil {
				t.Fatalf("error starting game: %s", err.Error())
			}
			gameToCompareAgainst, err := New(tc.startingFEN)
			if err != nil {
				t.Fatalf("error starting game: %s", err.Error())
			}

			err = gameToUpdate.Update(tc.move)

			if tc.expectedErr.expected {
				assert.Error(t, err, "expected a non nil error")
				assert.ErrorIs(t, err, tc.expectedErr.errType, "expected a different error...")
				assert.Equal(t, gameToCompareAgainst, gameToUpdate, "expected both games to be the same")
			} else {
				assert.NoError(t, err, "expected a nil error")
				moveParts := strings.Split(tc.move, "-")
				startingCoords := gameToUpdate.board.coords[moveParts[0]]
				endingCoords := gameToUpdate.board.coords[moveParts[1]]
				// knight is at ending coords
				endingSquare := gameToUpdate.board.squares[endingCoords.row][endingCoords.col]
				assert.True(t, endingSquare.isOccupied, "expected ending square to be occupied")
				assert.Equal(t, tc.pieceMoved, endingSquare.pieceAt, "expected ending square to contain the knight moved")
				// starting square is empty
				startingSquare := gameToUpdate.board.squares[startingCoords.row][startingCoords.col]
				assert.False(t, startingSquare.isOccupied, "expected starting square to be empty")
				assert.Equal(t, newEmpty(), startingSquare.pieceAt, "expected starting square to be an empty square")
				// turn changed
				assert.NotEqual(t, tc.pieceMoved.getPieceColor(), gameToUpdate.colorTurn, "turn didn't change")

				// knight moves always clear enpassant square
				assert.Equal(t, "", gameToUpdate.enPassantSquare, "enpassant square didn't clear")
			}
		})
	}
}

func TestKnightTriesEnPassant(t *testing.T) {
	type expectedErr struct {
		expected bool
		errType  error
	}

	type testCase struct {
		name        string
		startingFEN string
		move        string
		pieceMoved  piece
	}

	testCases := []testCase{
		{
			name:        "black knight tries capturing enpassant square at b3, invalid",
			move:        "d4-b3",
			startingFEN: "r1bqkb1r/pppppppp/5n2/8/1PPn4/2N2N2/P2PPPPP/R1BQKB1R b KQkq b3 0 1",
			pieceMoved:  newKnight(black),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gameToUpdate, err := New(tc.startingFEN)
			if err != nil {
				t.Fatalf("error starting game: %s", err.Error())
			}
			gameToUpdate.Update(tc.move)

			moveParts := strings.Split(tc.move, "-")
			startingCoords := gameToUpdate.board.coords[moveParts[0]]
			endingCoords := gameToUpdate.board.coords[moveParts[1]]
			// knight is at ending coords
			endingSquare := gameToUpdate.board.squares[endingCoords.row][endingCoords.col]
			assert.True(t, endingSquare.isOccupied, "expected ending square to be occupied")
			assert.Equal(t, tc.pieceMoved, endingSquare.pieceAt, "expected ending square to contain the knight moved")
			// starting square is empty
			startingSquare := gameToUpdate.board.squares[startingCoords.row][startingCoords.col]
			assert.False(t, startingSquare.isOccupied, "expected starting square to be empty")
			assert.Equal(t, newEmpty(), startingSquare.pieceAt, "expected starting square to be an empty square")
			// turn changed
			assert.NotEqual(t, tc.pieceMoved.getPieceColor(), gameToUpdate.colorTurn, "turn didn't change")

			// knight moves always clear enpassant square
			assert.Equal(t, "", gameToUpdate.enPassantSquare, "enpassant square didn't clear")
			// pawn still is at enpassant square
			pawnSquareCoords := gameToUpdate.board.coords["b4"]
			pawnSquare := gameToUpdate.board.squares[pawnSquareCoords.row][pawnSquareCoords.col]
			assert.True(t, pawnSquare.isOccupied, "expected pawn square to be occupied")
			assert.Equal(t, newPawn(white), pawnSquare.pieceAt, "expeted pawn square to contain a pawn")
		})
	}
}

func TestGetAvailableMovesBishop(t *testing.T) {
	type testCase struct {
		name           string
		startingFEN    string
		startingSquare string
		expectedMoves  []coords
	}

	testCases := []testCase{
		{
			name:           "c1 bishop starting moves",
			startingFEN:    "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			startingSquare: "c1",
			expectedMoves:  []coords{},
		},
		{
			name:           "c1 bishop with d2 not occupied",
			startingFEN:    "rnbqkbnr/ppp1pppp/3p4/8/8/3P4/PPP1PPPP/RNBQKBNR w KQkq - 0 1",
			startingSquare: "c1",
			expectedMoves:  []coords{{row: 6, col: 3}, {row: 5, col: 4}, {row: 4, col: 5}, {row: 3, col: 6}, {row: 2, col: 7}},
		},
		{
			name:           "c1 bishop with d2 not occupied, but own pawn blocking at f4",
			startingFEN:    "rnbqkbnr/pppp2p1/4p2p/5p2/5P2/2PP4/PP2P1PP/RNBQKBNR w KQkq - 0 1",
			startingSquare: "c1",
			expectedMoves:  []coords{{row: 6, col: 3}, {row: 5, col: 4}},
		},
		{
			name:           "c1 bishop with d2 not occupied, can take black pawn at g5",
			startingFEN:    "rnbqkbnr/ppp2pp1/8/3p2p1/4p3/3P1N1P/PPP1PP2/RNBQKB1R w KQkq - 0 1",
			startingSquare: "c1",
			expectedMoves:  []coords{{row: 6, col: 3}, {row: 5, col: 4}, {row: 4, col: 5}, {row: 3, col: 6}},
		},
		{
			name:           "c8 bishop starting moves",
			startingFEN:    "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
			startingSquare: "c8",
			expectedMoves:  []coords{},
		},
		{
			name:           "c8 bishop starting moves with d7 not occupied",
			startingFEN:    "rnbqkbnr/ppp1pppp/3p4/8/3PP3/8/PPP2PPP/RNBQKBNR b KQkq d3 0 1",
			startingSquare: "c8",
			expectedMoves:  []coords{{row: 1, col: 3}, {row: 2, col: 4}, {row: 3, col: 5}, {row: 4, col: 6}, {row: 5, col: 7}},
		},
		{
			name:           "c8 bishop starting moves with d7 not occupied, but own pawn blocking at g4",
			startingFEN:    "rnbqkbnr/ppp1pp1p/3p4/8/3P2p1/4P3/PPPN1PPP/RNBQKB1R b KQkq - 0 1",
			startingSquare: "c8",
			expectedMoves:  []coords{{row: 1, col: 3}, {row: 2, col: 4}, {row: 3, col: 5}},
		},
		{
			name:           "c8 bishop starting moves with d7 not occupied, can take at g4",
			startingFEN:    "rnbqkbnr/ppp1pp1p/3p4/8/3P2P1/P3P3/1PPN1PP1/RNBQKB1R b KQkq - 0 1",
			startingSquare: "c8",
			expectedMoves:  []coords{{row: 1, col: 3}, {row: 2, col: 4}, {row: 3, col: 5}, {row: 4, col: 6}},
		},
		{
			name:           "f1 bishop starting moves with e2 not occuppied",
			startingFEN:    "rnbqkbnr/pppppppp/8/8/8/4P3/PPPP1PPP/RNBQKBNR b KQkq - 0 1",
			startingSquare: "f1",
			expectedMoves:  []coords{{row: 6, col: 4}, {row: 5, col: 3}, {row: 4, col: 2}, {row: 3, col: 1}, {row: 2, col: 0}},
		},
		{
			name:           "f8 bishop starting moves with e7 not occuppied",
			startingFEN:    "rnbqkbnr/pppp1ppp/4p3/8/3PP3/8/PPP2PPP/RNBQKBNR b KQkq d3 0 1",
			startingSquare: "f8",
			expectedMoves:  []coords{{row: 1, col: 4}, {row: 2, col: 3}, {row: 3, col: 2}, {row: 4, col: 1}, {row: 5, col: 0}},
		},
		{
			name:           "white bishop at c4, lots many moves available",
			startingFEN:    "rnbqkb1r/ppp1p2p/3p1ppn/8/2B5/P3P2P/1PPP1PP1/RNBQK1NR w KQkq - 0 1",
			startingSquare: "c4",
			expectedMoves:  []coords{{row: 6, col: 0}, {row: 5, col: 1}, {row: 3, col: 3}, {row: 2, col: 4}, {row: 1, col: 5}, {row: 0, col: 6}, {row: 2, col: 0}, {row: 3, col: 1}, {row: 5, col: 3}, {row: 6, col: 4}, {row: 7, col: 5}},
		},
		{
			name:           "black bishop at f5, lots many moves available",
			startingFEN:    "rn1qkb1r/ppp1p2p/3p1ppn/5b2/2B3P1/P3P2P/1PPP1P2/RNBQK1NR b KQkq g3 0 1",
			startingSquare: "f5",
			expectedMoves:  []coords{{row: 6, col: 2}, {row: 5, col: 3}, {row: 4, col: 4}, {row: 0, col: 2}, {row: 1, col: 3}, {row: 2, col: 4}, {row: 4, col: 6}},
		},
		{
			name:           "white bishop at d4, lots of moves available",
			startingFEN:    "rnbqkbnr/p1pppppp/1p6/8/3B4/8/PPPPPPPP/RN1QKBNR w KQkq - 0 1",
			startingSquare: "d4",
			expectedMoves:  []coords{{row: 5, col: 2}, {row: 3, col: 4}, {row: 2, col: 5}, {row: 5, col: 4}, {row: 1, col: 6}, {row: 3, col: 2}, {row: 2, col: 1}},
		},
		{
			name:           "white bishop at a1, whole diagonal free",
			startingFEN:    "rnbqkbn1/p2ppp1r/1pp3pp/8/1P6/P7/R1PPPPPP/BN1QKBNR w Kq - 0 1",
			startingSquare: "a1",
			expectedMoves:  []coords{{row: 6, col: 1}, {row: 5, col: 2}, {row: 4, col: 3}, {row: 3, col: 4}, {row: 2, col: 5}, {row: 1, col: 6}, {row: 0, col: 7}},
		},
		{
			name:           "white bishop at h1, whole diagonal free",
			startingFEN:    "1nbqkbnr/r1pppppp/8/p7/1p6/6PP/PPPPPP1R/RNBQK1NB w Qk - 0 1",
			startingSquare: "h1",
			expectedMoves:  []coords{{row: 6, col: 6}, {row: 5, col: 5}, {row: 4, col: 4}, {row: 3, col: 3}, {row: 2, col: 2}, {row: 1, col: 1}, {row: 0, col: 0}},
		},
		{
			name:           "black bishop at a8, whole diagonal free",
			startingFEN:    "bn1qkbnr/2pppppp/rp6/p7/7P/2PP2P1/PP2PP1R/RNBQKBN1 w Qk - 0 1",
			startingSquare: "a8",
			expectedMoves:  []coords{{row: 6, col: 6}, {row: 5, col: 5}, {row: 4, col: 4}, {row: 3, col: 3}, {row: 2, col: 2}, {row: 1, col: 1}, {row: 7, col: 7}},
		},
		{
			name:           "white bishop at h8, whole diagonal free",
			startingFEN:    "rnbqk1nb/pppppp2/6pr/7p/P7/RP1P1P2/2P1P1PP/1NBQKBNR w Kq - 0 1",
			startingSquare: "h8",
			expectedMoves:  []coords{{row: 6, col: 1}, {row: 5, col: 2}, {row: 4, col: 3}, {row: 3, col: 4}, {row: 2, col: 5}, {row: 1, col: 6}, {row: 7, col: 0}},
		},
		{
			name:           "white bishop in the middle, 4 moves available",
			startingFEN:    "4k3/8/1P3P2/8/3B4/8/1P3P2/4K3 w - - 0 1",
			startingSquare: "d4",
			expectedMoves:  []coords{{row: 3, col: 2}, {row: 3, col: 4}, {row: 5, col: 2}, {row: 5, col: 4}},
		},
		{
			name:           "black bishop in the middle, 4 moves available",
			startingFEN:    "3k4/8/1p3p2/8/3b4/6N1/1p3p2/3K4 b - - 0 1",
			startingSquare: "d4",
			expectedMoves:  []coords{{row: 3, col: 2}, {row: 3, col: 4}, {row: 5, col: 2}, {row: 5, col: 4}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			game, err := New(tc.startingFEN)
			if err != nil {
				t.Fatalf("error starting game: %s", err.Error())
			}

			startingCoords := game.board.coords[tc.startingSquare]
			bishop := game.getPieceAtCoords(startingCoords)
			gotMoves := bishop.getAvailableMoves(startingCoords, game)

			assert.ElementsMatch(t, tc.expectedMoves, gotMoves, "expected different moves")
		})
	}
}

func TestBishopUpdateGame(t *testing.T) {
	type expectedErr struct {
		expected bool
		errType  error
	}

	type testCase struct {
		name        string
		startingFEN string
		move        string
		expectedErr expectedErr
		pieceMoved  piece
	}

	testCases := []testCase{
		{
			name:        "invalid white bishop move, no moves available",
			move:        "c1-f4",
			startingFEN: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			expectedErr: expectedErr{expected: true, errType: ErrInvalidMove},
			pieceMoved:  newBishop(white),
		},
		{
			name:        "invalid white bishop move, square is occupied by ally",
			move:        "c3-b2",
			startingFEN: "3K4/8/3b1N2/8/8/2B5/1P6/3k4 w - - 0 1",
			expectedErr: expectedErr{expected: true, errType: ErrInvalidMove},
			pieceMoved:  newBishop(white),
		},
		{
			name:        "invalid white bishop move, illegal move",
			move:        "c1-d3",
			startingFEN: "rnbqkbnr/pppp1ppp/4p3/8/3P4/8/PPP1PPPP/RNBQKBNR w KQkq - 0 1",
			expectedErr: expectedErr{expected: true, errType: ErrInvalidMove},
			pieceMoved:  newBishop(white),
		},
		{
			name:        "valid white bishop move, takes at f6",
			move:        "c3-f6",
			startingFEN: "3k4/8/5n2/8/8/2B5/8/3K4 w - - 0 1",
			expectedErr: expectedErr{expected: false},
			pieceMoved:  newBishop(white),
		},
		{
			name:        "invalid black bishop move, no squares available",
			move:        "f8-b4",
			startingFEN: "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq d3 0 1",
			expectedErr: expectedErr{expected: true, errType: ErrInvalidMove},
			pieceMoved:  newBishop(black),
		},
		{
			name:        "invalid black bishop move, square is occuppied by ally",
			move:        "e8-a4",
			startingFEN: "4b2k/8/8/8/n7/4B3/8/2K5 b - - 0 1",
			expectedErr: expectedErr{expected: true, errType: ErrInvalidMove},
			pieceMoved:  newBishop(black),
		},
		{
			name:        "invalid black bishop move, illegal move",
			move:        "e6-d4",
			startingFEN: "rn1qkb1r/ppp1pp1p/4bn2/3p2p1/8/1PP2N1P/PB1PPPP1/RN1QKB1R b KQkq - 0 1",
			expectedErr: expectedErr{expected: true, errType: ErrInvalidMove},
			pieceMoved:  newBishop(black),
		},
		{
			name:        "valid black bishop move, takes at a1",
			move:        "g7-a1",
			startingFEN: "rnbqk1nr/ppppppbp/6p1/8/8/1P3N1P/P1PPPPP1/RNBQKB1R b KQkq - 0 1",
			expectedErr: expectedErr{expected: false},
			pieceMoved:  newBishop(black),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gameToUpdate, err := New(tc.startingFEN)
			if err != nil {
				t.Fatalf("error starting game: %s", err.Error())
			}
			gameToCompareAgainst, err := New(tc.startingFEN)
			if err != nil {
				t.Fatalf("error starting game: %s", err.Error())
			}

			moveParts := strings.Split(tc.move, "-")
			startingCoords := gameToUpdate.board.coords[moveParts[0]]
			startingSquare := gameToUpdate.board.squares[startingCoords.row][startingCoords.col]
			endingCoords := gameToUpdate.board.coords[moveParts[1]]
			endingSquare := gameToUpdate.board.squares[endingCoords.row][endingCoords.col]

			assert.IsType(t, &bishop{}, startingSquare.pieceAt, "expected a bishop at starting square, before moving")
			assert.Equal(t, tc.pieceMoved.getPieceColor(), startingSquare.pieceAt.getPieceColor())
			err = gameToUpdate.Update(tc.move)

			startingCoords = gameToUpdate.board.coords[moveParts[0]]
			startingSquare = gameToUpdate.board.squares[startingCoords.row][startingCoords.col]
			endingCoords = gameToUpdate.board.coords[moveParts[1]]
			endingSquare = gameToUpdate.board.squares[endingCoords.row][endingCoords.col]

			if tc.expectedErr.expected {
				assert.Error(t, err, "expected a non nil error")
				assert.ErrorIs(t, err, tc.expectedErr.errType, "expected a different error...")
				assert.Equal(t, gameToCompareAgainst, gameToUpdate, "expected both games to be the same")
			} else {
				assert.NoError(t, err, "expected a nil error")
				// knight is at ending coords
				assert.True(t, endingSquare.isOccupied, "expected ending square to be occupied")
				assert.Equal(t, tc.pieceMoved, endingSquare.pieceAt, "expected ending square to contain the bishop moved")
				// starting square is empty
				assert.False(t, startingSquare.isOccupied, "expected starting square to be empty")
				assert.Equal(t, newEmpty(), startingSquare.pieceAt, "expected starting square to be an empty square")
				// turn changed
				assert.NotEqual(t, tc.pieceMoved.getPieceColor(), gameToUpdate.colorTurn, "turn didn't change")

				// bishop moves always clear enpassant square
				assert.Equal(t, "", gameToUpdate.enPassantSquare, "enpassant square didn't clear")
			}
		})
	}
}

func TestGetAvailableMovesRook(t *testing.T) {
	type testCase struct {
		name           string
		startingFEN    string
		startingSquare string
		expectedMoves  []coords
	}

	testCases := []testCase{
		{
			name:           "a1 white rook starting moves",
			startingFEN:    "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			startingSquare: "a1",
			expectedMoves:  []coords{},
		},
		{
			name:           "h1 white rook starting moves",
			startingFEN:    "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			startingSquare: "h1",
			expectedMoves:  []coords{},
		},
		{
			name:           "a8 white rook starting moves",
			startingFEN:    "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq d3 0 1",
			startingSquare: "a8",
			expectedMoves:  []coords{},
		},
		{
			name:           "h8 white rook starting moves",
			startingFEN:    "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq d3 0 1",
			startingSquare: "h8",
			expectedMoves:  []coords{},
		},
		{
			name:           "white rook in the middle, many moves available",
			startingFEN:    "4k3/8/8/8/3R4/8/8/1K6 w - - 0 1",
			startingSquare: "d4",
			expectedMoves:  []coords{{row: 7, col: 3}, {row: 6, col: 3}, {row: 5, col: 3}, {row: 3, col: 3}, {row: 2, col: 3}, {row: 1, col: 3}, {row: 0, col: 3}, {row: 4, col: 0}, {row: 4, col: 1}, {row: 4, col: 2}, {row: 4, col: 4}, {row: 4, col: 5}, {row: 4, col: 6}, {row: 4, col: 7}},
		},
		{
			name:           "black rook in the middle, many moves available",
			startingFEN:    "4k3/8/8/8/3r4/5R2/8/1K6 b - - 0 1",
			startingSquare: "d4",
			expectedMoves:  []coords{{row: 7, col: 3}, {row: 6, col: 3}, {row: 5, col: 3}, {row: 3, col: 3}, {row: 2, col: 3}, {row: 1, col: 3}, {row: 0, col: 3}, {row: 4, col: 0}, {row: 4, col: 1}, {row: 4, col: 2}, {row: 4, col: 4}, {row: 4, col: 5}, {row: 4, col: 6}, {row: 4, col: 7}},
		},
		{
			name:           "white rook at h8, all edges available",
			startingFEN:    "7R/4k3/8/8/3r4/8/1K6/8 w - - 0 1",
			startingSquare: "h8",
			expectedMoves:  []coords{{row: 0, col: 0}, {row: 0, col: 1}, {row: 0, col: 2}, {row: 0, col: 3}, {row: 0, col: 4}, {row: 0, col: 5}, {row: 0, col: 6}, {row: 1, col: 7}, {row: 2, col: 7}, {row: 3, col: 7}, {row: 4, col: 7}, {row: 5, col: 7}, {row: 6, col: 7}, {row: 7, col: 7}},
		},
		{
			name:           "white rook at a8, all edges available",
			startingFEN:    "R7/4k3/8/8/8/3r4/1K6/8 w - - 0 1",
			startingSquare: "a8",
			expectedMoves:  []coords{{row: 0, col: 7}, {row: 0, col: 1}, {row: 0, col: 2}, {row: 0, col: 3}, {row: 0, col: 4}, {row: 0, col: 5}, {row: 0, col: 6}, {row: 1, col: 0}, {row: 2, col: 0}, {row: 3, col: 0}, {row: 4, col: 0}, {row: 5, col: 0}, {row: 6, col: 0}, {row: 7, col: 0}},
		},
		{
			name:           "black rook at a1, all edges available",
			startingFEN:    "8/4k3/6R1/8/8/8/1K6/r7 b - - 0 1",
			startingSquare: "a1",
			expectedMoves:  []coords{{row: 7, col: 7}, {row: 7, col: 1}, {row: 7, col: 2}, {row: 7, col: 3}, {row: 7, col: 4}, {row: 7, col: 5}, {row: 7, col: 6}, {row: 1, col: 0}, {row: 2, col: 0}, {row: 3, col: 0}, {row: 4, col: 0}, {row: 5, col: 0}, {row: 6, col: 0}, {row: 0, col: 0}},
		},
		{
			name:           "white rook in the middle, take available in four directions",
			startingFEN:    "8/4k3/4p3/8/2n1R1b1/8/1K2pr2/8 w - - 0 1",
			startingSquare: "e4",
			expectedMoves:  []coords{{row: 2, col: 4}, {row: 3, col: 4}, {row: 5, col: 4}, {row: 6, col: 4}, {row: 4, col: 2}, {row: 4, col: 3}, {row: 4, col: 5}, {row: 4, col: 6}},
		},
		{
			name:           "black rook in the middle, take available in four directions",
			startingFEN:    "8/1k6/5P2/8/3P1r1P/8/1K3P2/8 b - - 0 1",
			startingSquare: "f4",
			expectedMoves:  []coords{{row: 4, col: 3}, {row: 4, col: 4}, {row: 4, col: 6}, {row: 4, col: 7}, {row: 2, col: 5}, {row: 3, col: 5}, {row: 5, col: 5}, {row: 6, col: 5}},
		},
		{
			name:           "white rook in the middle, surrounded by allies",
			startingFEN:    "8/3P1K2/8/1P1R3P/8/8/1k1P4/8 w - - 0 1",
			startingSquare: "d5",
			expectedMoves:  []coords{{row: 3, col: 2}, {row: 3, col: 4}, {row: 3, col: 5}, {row: 3, col: 6}, {row: 2, col: 3}, {row: 4, col: 3}, {row: 5, col: 3}},
		},
		{
			name:           "black rook in the middle, surrounded by allies",
			startingFEN:    "8/5k2/4b3/8/2q1r2n/8/1K6/4n3 b - - 0 1",
			startingSquare: "e4",
			expectedMoves:  []coords{{row: 4, col: 3}, {row: 4, col: 5}, {row: 4, col: 6}, {row: 3, col: 4}, {row: 5, col: 4}, {row: 6, col: 4}},
		},
		{
			name:           "white rook surrounded by enemy pawns",
			startingFEN:    "2k5/8/8/4p3/3pRp2/4p3/1K6/8 w - - 0 1",
			startingSquare: "e4",
			expectedMoves:  []coords{{row: 4, col: 3}, {row: 4, col: 5}, {row: 3, col: 4}, {row: 5, col: 4}},
		},
		{
			name:           "black rook surrounded by enemy pawns",
			startingFEN:    "2k5/5P2/4PrP1/5P2/8/8/1K6/8 b - - 0 1",
			startingSquare: "f6",
			expectedMoves:  []coords{{row: 2, col: 4}, {row: 2, col: 6}, {row: 1, col: 5}, {row: 3, col: 5}},
		},
		{
			name:           "white rook in the middle, with ally pawn at the right",
			startingFEN:    "4k3/8/8/8/3RP3/8/8/1K6 w - - 0 1",
			startingSquare: "d4",
			expectedMoves:  []coords{{row: 0, col: 3}, {row: 1, col: 3}, {row: 2, col: 3}, {row: 3, col: 3}, {row: 5, col: 3}, {row: 6, col: 3}, {row: 7, col: 3}, {row: 4, col: 0}, {row: 4, col: 1}, {row: 4, col: 2}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			game, err := New(tc.startingFEN)
			if err != nil {
				t.Fatalf("error starting game: %s", err.Error())
			}

			startingCoords := game.board.coords[tc.startingSquare]
			r := game.getPieceAtCoords(startingCoords)

			assert.IsType(t, &rook{}, r, "expected a rook as piece in starting square")
			gotMoves := r.getAvailableMoves(startingCoords, game)

			assert.ElementsMatch(t, tc.expectedMoves, gotMoves, "expected different moves")
		})
	}
}

func TestRookUpdateGame(t *testing.T) {
	type expectedErr struct {
		expected bool
		errType  error
	}

	type testCase struct {
		name                        string
		startingFEN                 string
		move                        string
		expectedErr                 expectedErr
		pieceMoved                  piece
		expectedWhiteCastlingRights castlingRights
		expectedBlackCastlingRights castlingRights
	}

	testCases := []testCase{
		{
			name:                        "white rook, right caputre",
			startingFEN:                 "2b5/8/4k3/8/p1R3n1/2r5/8/1K6 w - - 0 1",
			move:                        "c4-g4",
			expectedErr:                 expectedErr{expected: false},
			pieceMoved:                  newRook(white),
			expectedWhiteCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
			expectedBlackCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
		},
		{
			name:                        "white rook, left caputre",
			startingFEN:                 "2b5/8/4k3/8/p1R3n1/2r5/8/1K6 w - - 0 1",
			move:                        "c4-a4",
			expectedErr:                 expectedErr{expected: false},
			pieceMoved:                  newRook(white),
			expectedWhiteCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
			expectedBlackCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
		},
		{
			name:                        "white rook, up capture",
			startingFEN:                 "2b5/8/4k3/8/p1R3n1/2r5/8/1K6 w - - 0 1",
			move:                        "c4-c8",
			expectedErr:                 expectedErr{expected: false},
			pieceMoved:                  newRook(white),
			expectedWhiteCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
			expectedBlackCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
		},
		{
			name:                        "white rook, down capture",
			startingFEN:                 "2b5/8/4k3/8/p1R3n1/2r5/8/1K6 w - - 0 1",
			move:                        "c4-c3",
			expectedErr:                 expectedErr{expected: false},
			pieceMoved:                  newRook(white),
			expectedWhiteCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
			expectedBlackCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
		},
		{
			name:                        "black rook, right capture",
			startingFEN:                 "2P4k/Q1r4N/8/8/8/8/2P5/K7 b - - 0 1",
			move:                        "c7-h7",
			expectedErr:                 expectedErr{expected: false},
			pieceMoved:                  newRook(black),
			expectedWhiteCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
			expectedBlackCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
		},
		{
			name:                        "black rook, left capture",
			startingFEN:                 "2P4k/Q1r4N/8/8/8/8/2P5/K7 b - - 0 1",
			move:                        "c7-a7",
			expectedErr:                 expectedErr{expected: false},
			pieceMoved:                  newRook(black),
			expectedWhiteCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
			expectedBlackCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
		},
		{
			name:                        "black rook, up capture",
			startingFEN:                 "2P4k/Q1r4N/8/8/8/8/2P5/K7 b - - 0 1",
			move:                        "c7-c8",
			expectedErr:                 expectedErr{expected: false},
			pieceMoved:                  newRook(black),
			expectedWhiteCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
			expectedBlackCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
		},
		{
			name:                        "black rook, down capture",
			startingFEN:                 "2P4k/Q1r4N/8/8/8/8/2P5/K7 b - - 0 1",
			move:                        "c7-c2",
			expectedErr:                 expectedErr{expected: false},
			pieceMoved:                  newRook(black),
			expectedWhiteCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
			expectedBlackCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
		},
		{
			name:                        "white rook, invalid move",
			startingFEN:                 "8/8/4k3/2n5/8/3R4/1K6/8 w - - 0 1",
			move:                        "d3-c5",
			expectedErr:                 expectedErr{expected: true, errType: ErrInvalidMove},
			pieceMoved:                  newRook(white),
			expectedWhiteCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
			expectedBlackCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
		},
		{
			name:                        "black rook, tries to jump over piece",
			startingFEN:                 "8/8/4k3/2n3r1/8/3R2B1/1K6/8 b - - 0 1",
			move:                        "g5-g2",
			expectedErr:                 expectedErr{expected: true, errType: ErrInvalidMove},
			pieceMoved:                  newRook(black),
			expectedWhiteCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
			expectedBlackCastlingRights: castlingRights{canKingSide: false, canQueenSide: false},
		},
		{
			name:                        "white rook, valid move with enpassant active and updates white queen side castling",
			startingFEN:                 "r1bqkbnr/pppp3p/2n1ppp1/8/1P1P4/4P3/P1PN1PPP/R1BQKBNR w KQkq - 0 1",
			move:                        "a1-b1",
			expectedErr:                 expectedErr{expected: false},
			pieceMoved:                  newRook(white),
			expectedWhiteCastlingRights: castlingRights{canKingSide: true, canQueenSide: false},
			expectedBlackCastlingRights: castlingRights{canKingSide: true, canQueenSide: true},
		},
		{
			name:                        "black rook, valid move with enpassant active and updates white queen side castling",
			startingFEN:                 "r1bqkb1r/ppppn3/2n1ppp1/7p/1P1P4/4PN1P/P1PNBPP1/1RBQK2R b Kkq - 0 1",
			move:                        "h8-g8",
			expectedErr:                 expectedErr{expected: false},
			pieceMoved:                  newRook(black),
			expectedWhiteCastlingRights: castlingRights{canKingSide: true, canQueenSide: false},
			expectedBlackCastlingRights: castlingRights{canKingSide: false, canQueenSide: true},
		},
		{
			name:                        "white rook, valid move with enpassant active",
			startingFEN:                 "rnbqkbnr/pppp1ppp/8/4p3/8/7P/PPPPPPP1/RNBQKBNR w KQkq e6 0 2",
			move:                        "h1-h2",
			expectedErr:                 expectedErr{expected: false},
			pieceMoved:                  newRook(white),
			expectedWhiteCastlingRights: castlingRights{canKingSide: false, canQueenSide: true},
			expectedBlackCastlingRights: castlingRights{canKingSide: true, canQueenSide: true},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gameToUpdate, err := New(tc.startingFEN)
			if err != nil {
				t.Fatalf("error starting game: %s", err.Error())
			}
			gameToCompareAgainst, err := New(tc.startingFEN)
			if err != nil {
				t.Fatalf("error starting game: %s", err.Error())
			}

			moveParts := strings.Split(tc.move, "-")
			startingCoords := gameToUpdate.board.coords[moveParts[0]]
			startingSquare := gameToUpdate.board.squares[startingCoords.row][startingCoords.col]
			endingCoords := gameToUpdate.board.coords[moveParts[1]]
			endingSquare := gameToUpdate.board.squares[endingCoords.row][endingCoords.col]

			assert.IsType(t, &rook{}, startingSquare.pieceAt, "expected a rook at starting square, before moving")
			assert.Equal(t, tc.pieceMoved.getPieceColor(), startingSquare.pieceAt.getPieceColor())
			err = gameToUpdate.Update(tc.move)

			startingCoords = gameToUpdate.board.coords[moveParts[0]]
			startingSquare = gameToUpdate.board.squares[startingCoords.row][startingCoords.col]
			endingCoords = gameToUpdate.board.coords[moveParts[1]]
			endingSquare = gameToUpdate.board.squares[endingCoords.row][endingCoords.col]

			if tc.expectedErr.expected {
				assert.Error(t, err, "expected a non nil error")
				assert.ErrorIs(t, err, tc.expectedErr.errType, "expected a different error...")
				assert.Equal(t, gameToCompareAgainst, gameToUpdate, "expected both games to be the same")
			} else {
				assert.NoError(t, err, "expected a nil error")
				// knight is at ending coords
				assert.True(t, endingSquare.isOccupied, "expected ending square to be occupied")
				assert.Equal(t, tc.pieceMoved, endingSquare.pieceAt, "expected ending square to contain the rook moved")
				// starting square is empty
				assert.False(t, startingSquare.isOccupied, "expected starting square to be empty")
				assert.Equal(t, newEmpty(), startingSquare.pieceAt, "expected starting square to be an empty square")
				// turn changed
				assert.NotEqual(t, tc.pieceMoved.getPieceColor(), gameToUpdate.colorTurn, "turn didn't change")

				// rook moves always clear enpassant square
				assert.Equal(t, "", gameToUpdate.enPassantSquare, "enpassant square didn't clear")
				// check castling rights
				assert.Equal(t, tc.expectedWhiteCastlingRights, gameToUpdate.whiteCastlingRights, "expected white castling right to be different")
				assert.Equal(t, tc.expectedBlackCastlingRights, gameToUpdate.blackCastlingRights, "expected black castling right to be different")

			}
		})
	}
}

func TestGetAvailableMovesQueen(t *testing.T) {
	type testCase struct {
		name           string
		startingFEN    string
		startingSquare string
		expectedMoves  []coords
	}

	testCases := []testCase{
		{
			name:           "white queen in the middle, surrounded by enemies",
			startingFEN:    "6k1/8/2npp3/2pQp3/2ppn3/8/8/K7 w - - 0 1",
			startingSquare: "d5",
			expectedMoves:  []coords{{row: 2, col: 2}, {row: 2, col: 3}, {row: 2, col: 4}, {row: 3, col: 2}, {row: 3, col: 4}, {row: 4, col: 2}, {row: 4, col: 3}, {row: 4, col: 4}},
		},
		{
			name:           "black queen in the middle, surrounded by enemies",
			startingFEN:    "6k1/8/2NPN3/2RqR3/2PBP3/8/8/K7 b - - 0 1",
			startingSquare: "d5",
			expectedMoves:  []coords{{row: 2, col: 2}, {row: 2, col: 3}, {row: 2, col: 4}, {row: 3, col: 2}, {row: 3, col: 4}, {row: 4, col: 2}, {row: 4, col: 3}, {row: 4, col: 4}},
		},
		{
			name:           "white queen at a1, all moves available",
			startingFEN:    "8/8/8/8/8/1k6/4K3/Q7 w - - 0 1",
			startingSquare: "a1",
			expectedMoves: []coords{
				{row: 7, col: 1},
				{row: 7, col: 2},
				{row: 7, col: 3},
				{row: 7, col: 4},
				{row: 7, col: 5},
				{row: 7, col: 6},
				{row: 7, col: 7},
				{row: 0, col: 0},
				{row: 1, col: 0},
				{row: 2, col: 0},
				{row: 3, col: 0},
				{row: 4, col: 0},
				{row: 5, col: 0},
				{row: 6, col: 0},
				{row: 6, col: 1},
				{row: 5, col: 2},
				{row: 4, col: 3},
				{row: 3, col: 4},
				{row: 2, col: 5},
				{row: 1, col: 6},
				{row: 0, col: 7},
			},
		},
		{
			name:           "black queen at a8, all moves available",
			startingFEN:    "q7/8/8/8/8/1k6/4K3/8 b - - 0 1",
			startingSquare: "a8",
			expectedMoves: []coords{
				{row: 0, col: 1},
				{row: 0, col: 2},
				{row: 0, col: 3},
				{row: 0, col: 4},
				{row: 0, col: 5},
				{row: 0, col: 6},
				{row: 0, col: 7},
				{row: 1, col: 0},
				{row: 2, col: 0},
				{row: 3, col: 0},
				{row: 4, col: 0},
				{row: 5, col: 0},
				{row: 6, col: 0},
				{row: 7, col: 0},
				{row: 1, col: 1},
				{row: 2, col: 2},
				{row: 3, col: 3},
				{row: 4, col: 4},
				{row: 5, col: 5},
				{row: 6, col: 6},
				{row: 7, col: 7},
			},
		},
		{
			name:           "white queen at h8, all moves available",
			startingFEN:    "7Q/8/8/8/8/1k6/4K3/8 w - - 0 1",
			startingSquare: "h8",
			expectedMoves: []coords{
				{row: 0, col: 0},
				{row: 0, col: 1},
				{row: 0, col: 2},
				{row: 0, col: 3},
				{row: 0, col: 4},
				{row: 0, col: 5},
				{row: 0, col: 6},
				{row: 1, col: 7},
				{row: 2, col: 7},
				{row: 3, col: 7},
				{row: 4, col: 7},
				{row: 5, col: 7},
				{row: 6, col: 7},
				{row: 7, col: 7},
				{row: 1, col: 6},
				{row: 2, col: 5},
				{row: 3, col: 4},
				{row: 4, col: 3},
				{row: 5, col: 2},
				{row: 6, col: 1},
				{row: 7, col: 0},
			},
		},
		{
			name:           "black queen at h1, all moves available",
			startingFEN:    "8/8/8/8/8/1k6/4K3/7q b - - 0 1",
			startingSquare: "h1",
			expectedMoves: []coords{
				{row: 7, col: 0},
				{row: 7, col: 1},
				{row: 7, col: 2},
				{row: 7, col: 3},
				{row: 7, col: 4},
				{row: 7, col: 5},
				{row: 7, col: 6},
				{row: 0, col: 7},
				{row: 1, col: 7},
				{row: 2, col: 7},
				{row: 3, col: 7},
				{row: 4, col: 7},
				{row: 5, col: 7},
				{row: 6, col: 7},
				{row: 0, col: 0},
				{row: 1, col: 1},
				{row: 2, col: 2},
				{row: 3, col: 3},
				{row: 4, col: 4},
				{row: 5, col: 5},
				{row: 6, col: 6},
			},
		},
		{
			name:           "white queen in the middle, with allies one square in each direction",
			startingFEN:    "8/7k/1N1B1N2/8/1R1Q1R2/8/1N1B1N2/7K w - - 0 1",
			startingSquare: "d4",
			expectedMoves: []coords{
				{row: 3, col: 2},
				{row: 3, col: 3},
				{row: 3, col: 4},
				{row: 4, col: 2},
				{row: 4, col: 4},
				{row: 5, col: 2},
				{row: 5, col: 3},
				{row: 5, col: 4},
			},
		},
		{
			name:           "black queen in the middle, with allies one square in each direction",
			startingFEN:    "8/7K/1n1b1n2/8/1r1q1r2/8/1n1b1n2/7k b - - 0 1",
			startingSquare: "d4",
			expectedMoves: []coords{
				{row: 3, col: 2},
				{row: 3, col: 3},
				{row: 3, col: 4},
				{row: 4, col: 2},
				{row: 4, col: 4},
				{row: 5, col: 2},
				{row: 5, col: 3},
				{row: 5, col: 4},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			game, err := New(tc.startingFEN)
			if err != nil {
				t.Fatalf("error starting game: %s", err.Error())
			}

			startingCoords := game.board.coords[tc.startingSquare]
			q := game.getPieceAtCoords(startingCoords)

			assert.IsType(t, &queen{}, q, "expected a queen as piece in starting square")
			gotMoves := q.getAvailableMoves(startingCoords, game)

			assert.ElementsMatch(t, tc.expectedMoves, gotMoves, "expected different moves")
		})
	}
}

func callMoves(t *testing.T, game *Game, moves ...string) error {
	t.Helper()

	for _, move := range moves {
		err := game.Update(move)
		if err != nil {
			return fmt.Errorf("failed in move: %s. error: %w", move, err)
		}
	}

	return nil
}
