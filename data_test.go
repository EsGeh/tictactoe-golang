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
	game := NewGame()
	str := fmt.Sprint(game)
	if str != `|_|_|_|
|_|_|_|
|_|_|_|` {
		t.Fatalf("fmt.Sprint( game ) returned \"%v\"", str)
	}
}
