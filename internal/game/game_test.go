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
