package main

import (
	"fmt"
	"github.com/EsGeh/tictactoe-golang/ui"
	"math/rand"
	"time"
)

const (
	title  = "Tic Tac Toe"
	legend = `** Keys **

^C: Quit
K, ↑: Up
J, ↓: Down
H, ←: Left
L, →: Right
<Enter>: Play here`
)

func main() {
	rand.Seed(time.Now().UnixNano())
	game := NewGame()
	aiPlay(game)
	var uiHandle ui.UI
	uiHandle = ui.NewUI(
		title,
		legend,
		func(row int, col int){ selectedFunc(uiHandle, game, row, col) },
		func(row, col int){ selectionChangedFunc(uiHandle, game, row, col) },
	)
	updateTable( uiHandle, game )
	uiHandle.Run()
}

func selectedFunc(
	ui ui.UI,
	game game,
	row int,
	col int,
) {

	gameState := calcGameState(game)
	if gameState != continues {
		ui.SetStatus("The Game is over")
		return
	}
	if game[row][col] != 0 {
		ui.SetStatus("Forbidden move! Field already taken!")
		return
	}
	// user cannot select anything (this is released after ai is finished):
	ui.LockTable()
	game[row][col] = 2
	updateTable(ui, game)
	ui.SetStatus("good move!")
	go aiThinkAndPlay(ui, game)
}

// give user hints based on cursor position:
func selectionChangedFunc(
	ui ui.UI,
	game game,
	row, col int,
) {
	var status = ""
	switch game[row][col] {
	case player1, player2:
		status = "Field already taken!"
	default:
		status = "Press <Enter> to play here..."
	}
	if status != "" {
		ui.SetStatus(status)
	}
}

func updateTable(ui ui.UI, game game) {
	content := make([][]string, 0, 3)
	for _, row := range game {
		contentRow := make([]string, 0, 3)
		for _, cell := range row {
			contentRow = append(
				contentRow,
				fmt.Sprintf(" %v ", cell ),
			)
		}
		content = append(
			content,
			contentRow,
		)
	}
	ui.SetTable( content )
}

func aiThinkAndPlay(
	ui ui.UI,
	game game,
) {
	ui.ConcurrentUpdate(
		func() {
			time.Sleep(1000 * time.Millisecond)
			statusText := aiPlay(game)
			updateTable(ui, game)
			ui.SetStatus(statusText)
			ui.UnlockTable()
		},
	)
}
