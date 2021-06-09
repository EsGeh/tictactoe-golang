package main

import (
	"fmt"
	"testing"
)

func TestNewGame(t *testing.T) {
	game := NewGame()
	if len(game) != 3 {
		t.Fatalf("NewGame() returned wrong size \"%v\"", len(game))
	}
	for row_i, row := range game {
		if len(row) != 3 {
			t.Fatalf("len(NewGame()[%d]) == \"%d\", but %d expected", row_i, len(row), 3)
		}
	}
}

func TestGameToString(t *testing.T) {
	testMap := map[CellValue]string{
		Empty:   "_",
		Player1: "X",
		Player2: "O",
	}
	for cellValue, expectedStr := range testMap {
		resultStr := fmt.Sprint(cellValue)
		if resultStr != expectedStr {
			t.Fatalf("fmt.Sprint( %d ) returned \"%v\", but %v \"expected\"", cellValue, resultStr, expectedStr)
		}
	}
}

func TestScanAllLines(t *testing.T) {
	game := [][]CellValue{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	expectedLines := [][]cellInfo{
		{{0, 0, 0}, {0, 1, 1}, {0, 2, 2}}, // row 0
		{{1, 0, 3}, {1, 1, 4}, {1, 2, 5}}, // row 1
		{{2, 0, 6}, {2, 1, 7}, {2, 2, 8}}, // row 2
		{{0, 0, 0}, {1, 0, 3}, {2, 0, 6}}, // col 0
		{{0, 1, 1}, {1, 1, 4}, {2, 1, 7}}, // col 1
		{{0, 2, 2}, {1, 2, 5}, {2, 2, 8}}, // col 2
		{{0, 0, 0}, {1, 1, 4}, {2, 2, 8}}, // diag \
		{{0, 2, 2}, {1, 1, 4}, {2, 0, 6}}, // diag /
	}
	var linesScanned [][]cellInfo
	scanAllLines(
		game,
		func(cell []cellInfo) (found bool) {
			linesScanned = append(linesScanned, cell)
			return false
		},
	)
	for _, line := range linesScanned {
		if len(line) != 3 {
			t.Fatalf("scanAllLines scanned lines with len(line) != 3, but len(line) == %v, for %v", len(line), line)
		}
	}

	for _, expectedLine := range expectedLines {
		lineHasBeenScanned := false
		for _, line := range linesScanned {
			allEntriesSame := true
			for i := range expectedLine {
				if expectedLine[i] != line[i] {
					allEntriesSame = false
				}
			}
			if allEntriesSame {
				lineHasBeenScanned = true
			}
		}
		if !lineHasBeenScanned {
			t.Fatalf("scanAllLines didn't scan this line: %v", expectedLine)
		}
	}
}

func TestCalcGameState(t *testing.T) {
	testMap := []struct {
		gameData      game
		expectedState gameState
	}{
		{
			game{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			Continue,
		},
		// player 1 wins:
		{
			game{
				{1, 1, 1},
				{0, 0, 0},
				{0, 0, 0},
			},
			Player1Wins,
		},
		{
			game{
				{0, 0, 0},
				{1, 1, 1},
				{0, 0, 0},
			},
			Player1Wins,
		},
		{
			game{
				{0, 0, 0},
				{0, 0, 0},
				{1, 1, 1},
			},
			Player1Wins,
		},
		{
			game{
				{1, 0, 0},
				{1, 0, 0},
				{1, 0, 0},
			},
			Player1Wins,
		},
		{
			game{
				{0, 1, 0},
				{0, 1, 0},
				{0, 1, 0},
			},
			Player1Wins,
		},
		{
			game{
				{0, 0, 1},
				{0, 0, 1},
				{0, 0, 1},
			},
			Player1Wins,
		},
		{
			game{
				{1, 1, 1},
				{0, 0, 0},
				{0, 0, 0},
			},
			Player1Wins,
		},
		{
			game{
				{0, 0, 0},
				{1, 1, 1},
				{0, 0, 0},
			},
			Player1Wins,
		},
		{
			game{
				{0, 0, 0},
				{0, 0, 0},
				{1, 1, 1},
			},
			Player1Wins,
		},
		{
			game{
				{1, 0, 0},
				{1, 0, 0},
				{1, 0, 0},
			},
			Player1Wins,
		},
		{
			game{
				{0, 1, 0},
				{0, 1, 0},
				{0, 1, 0},
			},
			Player1Wins,
		},
		{
			game{
				{0, 0, 1},
				{0, 0, 1},
				{0, 0, 1},
			},
			Player1Wins,
		},
		{
			game{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			Player1Wins,
		},
		{
			game{
				{0, 0, 1},
				{0, 1, 0},
				{1, 0, 0},
			},
			Player1Wins,
		},
		// player 2 wins:
		{
			game{
				{2, 2, 2},
				{0, 0, 0},
				{0, 0, 0},
			},
			Player2Wins,
		},
		{
			game{
				{0, 0, 0},
				{2, 2, 2},
				{0, 0, 0},
			},
			Player2Wins,
		},
		{
			game{
				{0, 0, 0},
				{0, 0, 0},
				{2, 2, 2},
			},
			Player2Wins,
		},
		{
			game{
				{2, 0, 0},
				{2, 0, 0},
				{2, 0, 0},
			},
			Player2Wins,
		},
		{
			game{
				{0, 2, 0},
				{0, 2, 0},
				{0, 2, 0},
			},
			Player2Wins,
		},
		{
			game{
				{0, 0, 2},
				{0, 0, 2},
				{0, 0, 2},
			},
			Player2Wins,
		},
		{
			game{
				{2, 0, 0},
				{0, 2, 0},
				{0, 0, 2},
			},
			Player2Wins,
		},
		{
			game{
				{0, 0, 2},
				{0, 2, 0},
				{2, 0, 0},
			},
			Player2Wins,
		},
		// game over:
		{
			game{
				{1, 2, 1},
				{1, 2, 2},
				{2, 1, 1},
			},
			GameOver,
		},
		// continue:
		{
			game{
				{0, 2, 1},
				{1, 2, 2},
				{2, 1, 1},
			},
			Continue,
		},
		{
			game{
				{1, 0, 1},
				{1, 2, 2},
				{2, 1, 1},
			},
			Continue,
		},
	}
	for _, testEntry := range testMap {
		gameData, expectedState := testEntry.gameData, testEntry.expectedState
		state := calcGameState(gameData)
		if state != expectedState {
			t.Fatalf("calcGameState( %v ) returned %v, expected: %v", gameData, state, expectedState)
		}
	}
}
