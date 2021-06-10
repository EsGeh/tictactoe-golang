package main

import (
	"fmt"
	"github.com/rivo/tview"
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

var (
	app          *tview.Application
	statusWidget *tview.TextView
	gameData     game
)

func updateTable(table *tview.Table, gameData game) {
	for row_i := 0; row_i < 3; row_i++ {
		for col_i := 0; col_i < 3; col_i++ {
			c := fmt.Sprint(gameData[row_i][col_i])
			table.SetCell(
				row_i, col_i,
				tview.NewTableCell(" "+c+" ").
					SetAlign(tview.AlignCenter),
			)
		}
	}
}

func pcThinkAndPlay(table *tview.Table) {
	time.Sleep(1000 * time.Millisecond)
	status := aiPlay(gameData)
	app.QueueUpdateDraw(
		func() {
			updateTable(table, gameData)
			statusWidget.SetText(status)
			table.SetSelectable(true, true)
		},
	)
}

func newFieldTable() (table *tview.Table) {
	table = tview.NewTable().SetBorders(true)
	table.SetSelectable(true, true).SetSelectedFunc(
		func(row int, col int) {
			status := calcGameState(gameData)
			if status != continues {
				// game is over!
				statusWidget.SetText("The Game is over")
				return
			}
			if gameData[row][col] != 0 {
				statusWidget.SetText("Forbidden move! Field already taken!")
				return
			}
			table.SetSelectable(false, false)
			gameData[row][col] = 2
			updateTable(table, gameData)
			statusWidget.SetText("good move!")
			go pcThinkAndPlay(table)
		},
	)
	table.SetSelectionChangedFunc(
		func(row int, col int) {
			var status = ""
			switch gameData[row][col] {
			case player1, player2:
				status = "Field already taken!"
			default:
				status = "Press <Enter> to play here..."
			}
			if status != "" {
				statusWidget.SetText(status)
			}
		},
	)
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	gameData = NewGame()
	aiPlay(gameData)
	app = tview.NewApplication()
	table := newFieldTable()
	updateTable(table, gameData)

	statusWidget = tview.NewTextView().SetText("Welcome")
	legendWidget := tview.NewTextView().SetText(legend)
	root := tview.NewGrid().
		SetRows(1, 0, 1).
		SetColumns(20, 0).
		SetBorders(true).
		AddItem(tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(title), 0, 0, 1, 2, 0, 0, false).
		AddItem(tview.NewBox(), 1, 0, 1, 2, 0, 0, false).
		AddItem(legendWidget, 1, 0, 1, 1, 0, 0, false).
		AddItem(statusWidget, 2, 0, 1, 2, 0, 0, false).
		AddItem(table, 1, 1, 1, 1, 0, 0, false)
	if err := app.SetRoot(root, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}
