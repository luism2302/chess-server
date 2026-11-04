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
		{
			name:          "random invalid position",
			inputFENBoard: "rp3pk1/8/npq3br/1n1pb/K1R2PPR/2N2Q2/P4B1N/B2P4",
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
		turnAfterEnPassant bool // true white false black
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
			turnAfterEnPassant: false,
		},
		{
			name:               "enPassant in a6, to the left, white",
			initialFEN:         "rnbqkbnr/1p1ppppp/8/pPp5/8/8/P1PPPPPP/RNBQKBNR w KQkq a6 0 1",
			captureMove:        "b5-a6",
			takeSquare:         "a6",
			defenderPawnSquare: "a5",
			attackerPawnSquare: "b5",
			validEnPassant:     true,
			turnAfterEnPassant: false,
		},
		{
			name:               "enPassant in e3, to the right, black",
			initialFEN:         "rnbqkbnr/ppp1pppp/8/2P5/3pP3/8/PP1P1PPP/RNBQKBNR b KQkq e3 0 1",
			captureMove:        "d4-e3",
			takeSquare:         "e3",
			defenderPawnSquare: "e4",
			attackerPawnSquare: "d4",
			validEnPassant:     true,
			turnAfterEnPassant: true,
		},
		{
			name:               "enPassant in a3, to the left, black",
			initialFEN:         "rnbqkbnr/pp1ppppp/8/8/PpP5/8/3PPPPP/RNBQKBNR b KQkq a3 0 1",
			captureMove:        "b4-a3",
			takeSquare:         "a3",
			defenderPawnSquare: "a4",
			attackerPawnSquare: "b4",
			validEnPassant:     true,
			turnAfterEnPassant: true,
		},
		{
			name:               "enPassant in a3, to the left, black, but not performed",
			initialFEN:         "rnbqkbnr/pp1ppppp/8/8/PpP5/8/3PPPPP/RNBQKBNR b KQkq a3 0 1",
			captureMove:        "b4-b3",
			takeSquare:         "a3",
			defenderPawnSquare: "a4",
			attackerPawnSquare: "b4",
			validEnPassant:     false,
			turnAfterEnPassant: true,
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
				assert.Equal(t, newPawn(!tc.turnAfterEnPassant), takeToSquare.pieceAt)
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
			pieceMoved:              newPawn(true),
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
			pieceMoved:              newPawn(true),
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
			pieceMoved:              newPawn(false),
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
			pieceMoved:              newPawn(true),
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
			pieceMoved:              newPawn(false),
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
			pieceMoved:             newPawn(true),
			expectedPass:           false,
		},
		{
			name:                   "invalid two square move from white, piece at the end",
			startingFEN:            "rnbqkbnr/pppp1ppp/8/3P4/4p3/8/PPP1PPPP/RNBQKBNR w KQkq - 0 1",
			moveToTest:             "e2-e4",
			expectedEmptySquareErr: false,
			expectedWrongTurnErr:   false,
			expectedInvalidMoveErr: true,
			pieceMoved:             newPawn(true),
			expectedPass:           false,
		},
		{
			name:                   "invalid two square move from black, already moved",
			startingFEN:            "rnbqkbnr/ppppp2p/6p1/4Pp2/3P4/8/PPP2PPP/RNBQKBNR b KQkq d3 0 1",
			moveToTest:             "g6-g4",
			expectedEmptySquareErr: false,
			expectedWrongTurnErr:   false,
			expectedInvalidMoveErr: true,
			pieceMoved:             newPawn(false),
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
				assert.NotEqual(t, tc.pieceMoved.getPieceColor(), gameToUpdate.isWhiteTurn, "turn didnt change")
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
