package main

import (
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	table := tview.NewTable().
		SetBorders(true).
		SetSelectable(true, false).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEscape {
				app.Stop()
			}
		})

	for r, t := range tasks {
		table.SetCell(r, 0, tview.NewTableCell(t.UUID).
			SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
		table.SetCell(r, 1, tview.NewTableCell(t.Title).
			SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
		table.SetCell(r, 2, tview.NewTableCell(t.Created.Format(time.UnixDate)).
			SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
		table.SetCell(r, 3, tview.NewTableCell(t.Description[:10]+"...").
			SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	}

	if err := app.SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
}
