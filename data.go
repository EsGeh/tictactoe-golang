package main

import "fmt"

type cellValue int

const (
	empty cellValue = iota
	player1
	player2
)

type game [][]cellValue

type cellInfo struct {
	row, col int
	value    cellValue
}

type gameState int

const (
	continues gameState = iota
	player1Wins
	player2Wins
	gameOver
)

func (cell cellValue) String() string {
	ret := ""
	switch cell {
	case empty:
		ret = fmt.Sprint("_")
	case player1:
		ret = fmt.Sprint("X")
	case player2:
		ret = fmt.Sprint("O")
	}
	return ret
}

func (gameState gameState) String() string {
	switch gameState {
	case player1Wins:
		return "player1Wins"
	case player2Wins:
		return "player2Wins"
	case gameOver:
		return "gameOver"
	}
	return "continue"
}

func NewGame() game {
	game := make([][]cellValue, 3)
	for row, _ := range game {
		game[row] = make([]cellValue, 3)
	}
	return game
}

func calcGameState(game game) gameState {
	// 1. check if someone has won:
	{
		gameState := continues
		scanAllLines(
			game,
			func(line []cellInfo) bool {
				player1Counter, player2Counter := 0, 0
				for _, cellInfo := range line {
					switch cellInfo.value {
					case player1:
						player1Counter++
					case player2:
						player2Counter++
					}
				}
				if player1Counter == 3 {
					gameState = player1Wins
					return true
				}
				if player2Counter == 3 {
					gameState = player2Wins
					return true
				}
				return false
			},
		)
		if gameState == player1Wins || gameState == player2Wins {
			return gameState
		}
	}
	// 2. check if the game is over?
	for _, row := range game {
		for _, cell := range row {
			if cell == empty {
				return continues
			}
		}
	}
	return gameOver
}

func scanAllLines(
	game game,
	cond func([]cellInfo) (found bool),
) {
	// scan rows
	for row, line := range game {
		lineInfo := make([]cellInfo, 0, 3)
		for col, cell := range line {
			lineInfo = append(lineInfo, cellInfo{row, col, cell})
		}
		if cond(lineInfo) {
			return
		}
	}

	// scan columns
	for col := 0; col < 3; col++ {
		lineInfo := make([]cellInfo, 0, 3)
		for row := 0; row < 3; row++ {
			cell := game[row][col]
			lineInfo = append(lineInfo, cellInfo{row, col, cell})
		}
		if cond(lineInfo) {
			return
		}
	}

	// scan (full) diagonal \
	{
		lineInfo := make([]cellInfo, 0, 3)
		for col := 0; col < 3; col++ {
			row := col
			cell := game[row][col]
			lineInfo = append(lineInfo, cellInfo{row, col, cell})
		}
		if cond(lineInfo) {
			return
		}
	}

	// scan (full) diagonal /
	{
		lineInfo := make([]cellInfo, 0, 3)
		for row := 0; row < 3; row++ {
			col := 2 - row
			cell := game[row][col]
			lineInfo = append(lineInfo, cellInfo{row, col, cell})
		}
		if cond(lineInfo) {
			return
		}
	}
}
