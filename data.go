package main

import "fmt"

type CellValue int

const (
	Empty = iota
	Player1
	Player2
)

type game [][]CellValue

type cellInfo struct {
	row, col int
	value    CellValue
}

type gameState int

const (
	Continue = iota
	Player1Wins
	Player2Wins
	GameOver
)

func (cell CellValue) String() string {
	ret := ""
	switch cell {
	case Empty:
		ret = fmt.Sprint("_")
	case Player1:
		ret = fmt.Sprint("X")
	case Player2:
		ret = fmt.Sprint("O")
	}
	return ret
}

func NewGame() game {
	game := make([][]CellValue, 3)
	for row, _ := range game {
		game[row] = make([]CellValue, 3)
	}
	return game
}

func calcGameState(gameData game) gameState {
	// 1. check if someone has won:
	{
		state := gameState(Continue)
		scanAllLines(
			gameData,
			func(line []cellInfo) bool {
				player1, player2 := 0, 0
				for _, cellInfo := range line {
					switch cellInfo.value {
					case Player1:
						player1++
					case Player2:
						player2++
					}
				}
				if player1 == 3 {
					state = Player1Wins
					return true
				}
				if player2 == 3 {
					state = Player2Wins
					return true
				}
				return false
			},
		)
		if state == Player1Wins || state == Player2Wins {
			return state
		}
	}
	// 2. check if the game is over?
	for _, row := range gameData {
		for _, cell := range row {
			if cell == Empty {
				return Continue
			}
		}
	}
	return GameOver
}

func scanAllLines(
	gameData game,
	cond func([]cellInfo) (found bool),
) {
	// scan rows
	for row, line := range gameData {
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
			cell := gameData[row][col]
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
			cell := gameData[row][col]
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
			cell := gameData[row][col]
			lineInfo = append(lineInfo, cellInfo{row, col, cell})
		}
		if cond(lineInfo) {
			return
		}
	}
}
