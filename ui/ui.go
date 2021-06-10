package ui

import (
	"github.com/rivo/tview"
)

type UI struct {
	app *tview.Application
	tableWidget *tview.Table
	statusWidget *tview.TextView
}

func NewUI(
	title string,
	legend string,
	selectedFunc func(row int, col int),
	selectionChangedFunc func(row int, col int),
) UI {
	ui := UI{}
	ui.app = tview.NewApplication()
	//table := newFieldTable()
	ui.tableWidget = tview.NewTable().SetBorders(true).SetSelectable(true, true).
		SetSelectedFunc(selectedFunc).
		SetSelectionChangedFunc(selectionChangedFunc)
	// updateTable(table, gameData)
	ui.statusWidget = tview.NewTextView().SetText("Welcome")
	legendWidget := tview.NewTextView().SetText(legend)
	root := tview.NewGrid().
		SetRows(1, 0, 1).
		SetColumns(20, 0).
		SetBorders(true).
		AddItem(tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(title), 0, 0, 1, 2, 0, 0, false).
		AddItem(tview.NewBox(), 1, 0, 1, 2, 0, 0, false).
		AddItem(legendWidget, 1, 0, 1, 1, 0, 0, false).
		AddItem(ui.statusWidget, 2, 0, 1, 2, 0, 0, false).
		AddItem(ui.tableWidget, 1, 1, 1, 1, 0, 0, false)
	ui.app.SetRoot(root, true).SetFocus(ui.tableWidget)
	return ui
}

func (ui UI) Run() {
	if err := ui.app.Run(); err != nil {
		panic(err)
	}
}

func (ui UI) SetStatus( status string ) {
	ui.statusWidget.SetText( status )
}

func (ui UI) SetTable(
	content [][]string,
) {
	for rowI, row := range content {
		for colI, cell := range row {
			ui.tableWidget.SetCell(
				rowI, colI,
				tview.NewTableCell(cell).
					SetAlign(tview.AlignCenter),
			)
		}
	}
}


func (ui UI) ConcurrentUpdate( f func() ) {
	ui.app.QueueUpdateDraw( f )
}

func (ui UI) LockTable() {
	ui.tableWidget.SetSelectable(false, false)
}

func (ui UI) UnlockTable() {
	ui.tableWidget.SetSelectable(true, true)
}

