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
