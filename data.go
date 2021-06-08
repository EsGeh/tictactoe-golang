package main

import "fmt"

type CellValue int

const (
	Empty = iota
	Player1
	Player2
)

type game [][]CellValue

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

func (g game) String() string {
	ret := ""
	for row_i, row := range g {
		ret += "|"
		for _, cell := range row {
			ret += fmt.Sprint(cell)
			ret += "|"
		}
		if row_i != len(g)-1 {
			ret += "\n"
		}
	}
	return ret
}

func NewGame() game {
	game := make([][]CellValue, 3)
	for row_i, _ := range game {
		game[row_i] = make([]CellValue, 3)
	}
	return game
}
